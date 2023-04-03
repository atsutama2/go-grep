package main

import (
	"fmt"
	"os"

	"github.com/atsutama2/go-grep/pkg/grep"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gg \"search_word\"")
		os.Exit(1)
	}

	searchWord := os.Args[1]
	if err := grep.SearchFiles(searchWord, "."); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
