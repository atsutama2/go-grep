package grep

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func SearchFiles(pattern string, root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if err := grep(pattern, path); err != nil {
				return err
			}
		}
		return nil
	})
}

func grep(pattern string, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++
		if re.MatchString(line) {
			absPath, _ := filepath.Abs(path)
			fmt.Printf("%s:%d:%s\n", absPath, lineNumber, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
