package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("test1")
	fmt.Println("test2")
	fmt.Println("test3")
	fmt.Println("test4")

	versionFlag := flag.Bool("version", false, "Show the version of gg")
	funcFlag := flag.Bool("func", false, "Search for function names")
	structFlag := flag.Bool("struct", false, "Search for struct names")

	flag.Parse()

	if *versionFlag {
		fmt.Printf("gg version: %s\n", Version)
		return
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: gg <search_word> [<directory>] | gg -func <search_word> [<directory>]")
		os.Exit(1)
	}

	var searchWord, directory string
	funcMode := *funcFlag
	structMode := *structFlag

	searchWord = args[0]
	directory = "."
	if len(args) == 2 {
		directory = args[1]
	}

	err := Grep(searchWord, directory, funcMode, structMode, fmt.Printf)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
