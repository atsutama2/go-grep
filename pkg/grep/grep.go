package grep

import (
	"bufio"
	"fmt"
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

func processFile(searchWord, path, directory string, classMode bool, wg *sync.WaitGroup, mtx *sync.Mutex, matchCount *int32) {
	defer wg.Done()

	file, err := os.Open(path)
	if err != nil {
		mtx.Lock()
		fmt.Printf("Error: %v\n", err)
		mtx.Unlock()
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
		mtx.Lock()
		fmt.Printf("Error: %v\n", err)
		mtx.Unlock()
	}

	if len(results) > 0 {
		atomic.AddInt32(matchCount, int32(len(results))) // カウントを更新
		relPath, _ := filepath.Rel(directory, path)

		mtx.Lock()
		fmt.Printf("%s\n", colorPath(relPath))
		for _, result := range results {
			fmt.Printf("%d:%s\n", result.LineNumber, highlight(result.Line, searchWord))
		}
		fmt.Println()
		mtx.Unlock()
	}
}

func Grep(searchWord, directory string, classMode bool) error {
	var wg sync.WaitGroup
	var mtx sync.Mutex

	excludeList := []string{
		".git",
		"vendor",
		".vscode",
		"node_modules",
		"_build"} // 検索対象から除外するファイル/ディレクトリのリスト

	var matchCount int32 // 検索結果のカウント用の変数を追加

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, exclude := range excludeList {
			if info.Name() == exclude {
				if info.IsDir() {
					return filepath.SkipDir // ディレクトリをスキップ
				} else {
					return nil // ファイルをスキップ
				}
			}
		}

		if !info.IsDir() {
			wg.Add(1)
			go processFile(searchWord, path, directory, classMode, &wg, &mtx, &matchCount)
		}
		return nil
	})

	if err != nil {
		return err
	}

	wg.Wait()

	fmt.Printf("Total matches: %d\n", matchCount) // 検索結果のカウントを表示

	return nil
}
