# go-grep (gg)

`go-grep` (also known as `gg`) is a simple command-line tool written in Go that mimics the functionality of the Unix `grep` command. It searches for a specified text pattern in one or more files within a given directory and its subdirectories.

## Features

- Displays line numbers containing the search pattern.
- Displays the relative file path of the file containing the search pattern.
- Starts the search from the current directory if no directory is specified.
- Highlights the search pattern in the output.
- Shows the total number of matches found.

## Installation

To install `gg`, first clone the repository:

```
git clone https://github.com/atsutama2/go-grep.git
```

Then, navigate to the project directory and build the `gg` binary:

```
cd go-grep
go build ./cmd/gg
```

## Usage

To use `gg`, execute the binary followed by the search pattern and the directory to start the search (optional). If no directory is specified, the search will start from the current directory.

```
./gg <search_word> [<directory>]
```

### Examples

1. Search for the word "Apple" in the current directory:

```
./gg Apple
```

2. Search for the word "Apple" in the `testdata` directory:

```
./gg Apple testdata/
```

The output will display the relative file path, line number, and the line containing the search pattern, with the search pattern highlighted. At the end of the output, the total number of matches found will be displayed.
