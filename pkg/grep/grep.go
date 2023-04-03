package grep

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func highlight(text, searchWord string) string {
	return strings.ReplaceAll(text, searchWord, "\033[1;31m"+searchWord+"\033[0m")
}

func colorPath(path string) string {
	return "\033[1;34m\033[1m" + path + "\033[0m"
}

func processFile(searchWord, path, directory string, wg *sync.WaitGroup, mtx *sync.Mutex) {
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

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchWord) {
			relPath, _ := filepath.Rel(directory, path)

			mtx.Lock()
			fmt.Printf("%s\n%d:%s\n\n", colorPath(relPath), lineNumber, highlight(line, searchWord))
			mtx.Unlock()
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		mtx.Lock()
		fmt.Printf("Error: %v\n", err)
		mtx.Unlock()
	}
}

func Grep(searchWord, directory string) error {
	var wg sync.WaitGroup
	var mtx sync.Mutex

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			wg.Add(1)
			go processFile(searchWord, path, directory, &wg, &mtx)
		}
		return nil
	})

	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}
