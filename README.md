# go-grep (gg)
[![GitHub license](https://img.shields.io/github/license/onozaty/json2csv)](https://github.com/onozaty/json2csv/blob/main/LICENSE)
[![Test](https://github.com/atsutama2/go-grep/actions/workflows/ci.yml/badge.svg)](https://github.com/atsutama2/go-grep/actions/workflows/ci.yml)

高速で汎用性のあるテキスト検索ツール`go-grep`は、Goで書かれたコマンドラインツールで、ファイルやディレクトリ内のテキストを効率的に検索できます。
---
`go-grep` is a fast and versatile text search tool written in Go that allows for efficient searching of text within files and directories from the command line.

## Features

- 検索対象のファイルにおいて、検索パターンが含まれる行番号や相対パスを表示します。
- ディレクトリが指定されていない場合は、現在のディレクトリから検索を開始します。
- 出力には検索パターンがハイライト表示されます。
- 関数名の検索。(-func)
- 構造体名の検索。(-struct)
- 検索中に、検索対象外とするファイル名やディレクトリ名を指定することができます。
- 除外ファイルの格納場所は `$HOME/go-grep/exclude_list.txt` であり、ユーザーのホームディレクトリ内に `go-grep` ディレクトリを作成し、その中に `exclude_list.txt` ファイルを配置する必要があります。
---
- Displays line numbers containing the search pattern.
- Displays the relative file path of the file containing the search pattern.
- Starts the search from the current directory if no directory is specified.
- Highlights the search pattern in the output.
- Searches for method names. (go olny)
- Recursively searches for files and subdirectories in the specified directory.
- Allows exclusion of files and directories during the search. For example, .git or vendor.
- Skips files and directories that are excluded during the search.
- The location to store the exclude list file is $HOME/go-grep/exclude_list.txt.
- This means that you need to have a go-grep directory in your home directory and place the exclude_list.txt file in it. If the directory does not exist, you need to create it yourself.

## Installation

To install `gg`, first clone the repository:

```
git clone https://github.com/atsutama2/go-grep.git
```

Then, navigate to the project directory and build the `gg` binary:

```
cd go-grep
make build
(sudo mv gg /usr/local/bin)
input pass
```

Then, navigate to the project directory and build the `gg` binary:

```
gg <search_word> [<directory>]
```

The location for excluding files is as follows:
Please create exclude_list.txt.
```
$HOME/go-grep/exclude_list.txt
```

For example, it could look like this:
```
.git/
vendor/
.vscode/
node_modules/
_build/
lib/
pkg/
shardkey/
data/
bin/
gopls/
golangci-lint
coverage.html
```

## Usage

To use `gg`, execute the binary followed by the search pattern and the directory to start the search (optional). If no directory is specified, the search will start from the current directory.

```
gg <search_word> [<directory>]
```

```
atsutama2: ~/go/tmp/go-grep (feature/base-grep %=)$ gg Apple
README.md
56:1. Search for the word "Apple" in the current directory:
59:gg Apple
62:2. Search for the word "Apple" in the `testdata` directory:
65:gg Apple testdata/

Total matches: 4
```

You can search for method names.
```
gg -func <method name>
```

Search for struct names.
```
gg -struct <struct name>
```

```
atsutama2: ~/go/tmp/go-grep (feature/base-grep %=)$ gg -func processFile
pkg/grep/grep.go
40:func processFile(searchWord, path, directory string, classMode bool, wg *sync.WaitGroup, mtx *sync.Mutex, matchCount *int32, printfFunc func(string, ...interface{}) (int, error)) {

Total matches: 1
```

### Examples

1. Search for the word "Apple" in the current directory:

```
gg Apple
```

2. Search for the word "Apple" in the `testdata` directory:

```
gg Apple testdata/
```

The output will display the relative file path, line number, and the line containing the search pattern, with the search pattern highlighted. At the end of the output, the total number of matches found will be displayed.

The output will display the relative file path, line number, and the line containing the search pattern, with the search pattern highlighted.

## License
MIT
