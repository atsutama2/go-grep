# go-grep (gg)
[![GitHub license](https://img.shields.io/github/license/onozaty/json2csv)](https://github.com/onozaty/json2csv/blob/main/LICENSE)

高速で汎用性のあるテキスト検索ツール`go-grep`は、Goで書かれたコマンドラインツールで、ファイルやディレクトリ内のテキストを効率的に検索できます。大文字小文字を区別せずにマッチングし、特定のファイルやディレクトリを除外できます。

## Features

- Displays line numbers containing the search pattern.
- Displays the relative file path of the file containing the search pattern.
- Starts the search from the current directory if no directory is specified.
- Highlights the search pattern in the output.

## Installation

To install `gg`, first clone the repository:
```
git clone https://github.com/atsutama2/go-grep.git
```
Then, navigate to the project directory and build the `gg` binary:

```
cd go-grep
make build
sudo mv gg /usr/local/bin
```
Then, navigate to the project directory and build the `gg` binary:

```
./gg <search_word> [<directory>]
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

![スクリーンショット 2023-04-03 18 25 46](https://user-images.githubusercontent.com/77330599/229469245-e5e40f71-263f-4f3c-9cfd-2144d9b50675.png)


The output will display the relative file path, line number, and the line containing the search pattern, with the search pattern highlighted.

## License
MIT
