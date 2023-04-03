package grep

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func highlight(text, searchWord string) string {
	return strings.ReplaceAll(text, searchWord, "\033[1;31m"+searchWord+"\033[0m")
}

func Grep(searchWord, directory string) error {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			lineNumber := 1

			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, searchWord) {
					relPath, _ := filepath.Rel(directory, path)
					fmt.Printf("%s\n%d:%s\n\n", relPath, lineNumber, highlight(line, searchWord))
				}
				lineNumber++
			}

			if err := scanner.Err(); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
