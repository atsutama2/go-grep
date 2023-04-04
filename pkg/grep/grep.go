package grep

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
)

type SearchResult struct {
	LineNumber int
	Line       string
}

var lineBufferPool = &sync.Pool{
	New: func() interface{} {
		return &strings.Builder{}
	},
}

func highlight(text, searchWord string) string {
	re, err := regexp.Compile("(?i)" + regexp.QuoteMeta(searchWord))
	if err != nil {
		return text // エラーが発生した場合、元のテキストを返します
	}
	return re.ReplaceAllStringFunc(text, func(match string) string {
		return "\033[1;31m" + match + "\033[0m"
	})
}

func colorPath(path string) string {
	return "\033[1;34m\033[1m" + path + "\033[0m"
}

func processFile(searchWord, path, directory string, classMode bool, wg *sync.WaitGroup, mtx *sync.Mutex, matchCount *int32, printfFunc func(string, ...interface{}) (int, error)) {
	defer wg.Done()

	fileInfo, err := os.Stat(path)
	if err != nil {
		mtx.Lock()
		printfFunc("Error: %v\n", err)
		mtx.Unlock()
		return
	}

	if fileInfo.IsDir() {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		mtx.Lock()
		printfFunc("Error: %v\n", err)
		mtx.Unlock()
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	buf := make([]byte, 0, bufio.MaxScanTokenSize)
	scanner.Buffer(buf, 10*bufio.MaxScanTokenSize)

	lineNumber := 1

	var results []SearchResult

	for scanner.Scan() {
		line := scanner.Text()

		if classMode {
			match, err := regexp.MatchString(`func\s+`+searchWord+`\s*\(.*\)`, line)
			if err != nil {
				continue
			}
			if match {
				results = append(results, SearchResult{LineNumber: lineNumber, Line: line})
			}
		} else {
			lineBuffer := lineBufferPool.Get().(*strings.Builder)
			lineBuffer.Reset()
			lineBuffer.WriteString(strings.ToLower(line))
			lowerCaseLine := lineBuffer.String()

			if strings.Contains(lowerCaseLine, strings.ToLower(searchWord)) {
				results = append(results, SearchResult{LineNumber: lineNumber, Line: line})
			}

			lineBufferPool.Put(lineBuffer)
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		if err == bufio.ErrTooLong {
			mtx.Lock()
			fmt.Printf("Error: Skipping file with too long line: %s\n", path)
			mtx.Unlock()
			return
		} else {
			mtx.Lock()
			fmt.Printf("Error: %v\n", err)
			mtx.Unlock()
		}
	}

	if len(results) > 0 {
		atomic.AddInt32(matchCount, int32(len(results))) // カウントを更新
		relPath, _ := filepath.Rel(directory, path)

		mtx.Lock()
		printfFunc("%s\n", colorPath(relPath))
		for _, result := range results {
			printfFunc("%d:%s\n", result.LineNumber, highlight(result.Line, searchWord))
		}
		printfFunc("\n")
		mtx.Unlock()
	}
}

func Grep(searchWord, directory string, classMode bool, printfFunc func(string, ...interface{}) (int, error)) error {
	var wg sync.WaitGroup
	var mtx sync.Mutex

	// $HOME/go-grep/exclude_list.txt ファイルのパスを取得
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error: failed to get user home directory: %v", err)
	}
	excludeListPath := filepath.Join(homeDir, "go-grep", "exclude_list.txt")

	// excludeListPath から excludeList を読み込む
	excludeListBytes, err := ioutil.ReadFile(excludeListPath)
	if err != nil {
		return fmt.Errorf("Error: failed to read exclude list file: %v", err)
	}
	excludeList := strings.Split(strings.TrimSpace(string(excludeListBytes)), "\n")

	var matchCount int32 // 検索結果のカウント用の変数を追加

	// Create a semaphore
	semaphore := make(chan struct{}, 10)

	// Add a new channel to signal the completion of filepath.Walk
	walkDone := make(chan error)

	go func() {
		err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("Error: failed to walk path %q: %v", path, err)
			}

			for _, exclude := range excludeList {
				if strings.Contains(path, exclude) {
					if info.IsDir() {
						return filepath.SkipDir // ディレクトリをスキップ
					} else {
						return nil // ファイルをスキップ
					}
				}
			}

			if info.IsDir() {
				return nil
			}

			if (info.Mode() & os.ModeSocket) != 0 { // ディレクトリとソケットファイルをスキップします
				return nil
			}

			// Acquire the semaphore
			semaphore <- struct{}{}

			wg.Add(1)
			go func() {
				// Release the semaphore when the function completes
				defer func() { <-semaphore }()
				processFile(searchWord, path, directory, classMode, &wg, &mtx, &matchCount, printfFunc)
			}()
			return nil
		})

		// Signal the completion of filepath.Walk
		walkDone <- err
		close(walkDone)
	}()

	// Wait for filepath.Walk to complete
	err = <-walkDone

	if err != nil {
		return err
	}

	wg.Wait()

	fmt.Printf("Total matches: %d\n", matchCount) // 検索結果のカウントを表示

	return nil
}
