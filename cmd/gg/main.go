package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/atsutama2/go-grep/pkg/grep"
)

func main() {
	versionFlag := flag.Bool("version", false, "Show the version of gg")
	classFlag := flag.Bool("class", false, "Search for method names")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("gg version: %s\n", grep.Version)
		return
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: gg <search_word> [<directory>] | gg -class <search_word> [<directory>]")
		os.Exit(1)
	}

	var searchWord, directory string
	classMode := *classFlag

	searchWord = args[0]
	directory = "."
	if len(args) == 2 {
		directory = args[1]
	}

	err := grep.Grep(searchWord, directory, classMode, fmt.Printf)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
