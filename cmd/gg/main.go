package main

import (
	"fmt"
	"os"

	"github.com/atsutama2/go-grep/pkg/grep"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gg <search_word> [<directory>]")
		os.Exit(1)
	}

	searchWord := os.Args[1]

	directory := "."
	if len(os.Args) == 3 {
		directory = os.Args[2]
	}

	err := grep.Grep(searchWord, directory)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
