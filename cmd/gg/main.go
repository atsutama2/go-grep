package main

import (
	"fmt"
	"os"

	"github.com/atsutama2/go-grep/pkg/grep"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gg <search_word> [<directory>] | gg -class <search_word> [<directory>]")
		os.Exit(1)
	}

	var searchWord, directory string
	var classMode bool

	if os.Args[1] == "-class" {
		classMode = true
		searchWord = os.Args[2]
		directory = "."
		if len(os.Args) == 4 {
			directory = os.Args[3]
		}
	} else {
		searchWord = os.Args[1]
		directory = "."
		if len(os.Args) == 3 {
			directory = os.Args[2]
		}
	}

	err := grep.Grep(searchWord, directory, classMode, fmt.Printf)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
