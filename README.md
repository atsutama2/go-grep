# go-grep (gg)
[![GitHub license](https://img.shields.io/github/license/onozaty/json2csv)](https://github.com/onozaty/json2csv/blob/main/LICENSE)
[![Test](https://github.com/atsutama2/go-grep/actions/workflows/go.yml/badge.svg)](https://github.com/atsutama2/go-grep/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/atsutama2/go-grep/branch/main/graph/badge.svg?token=DPGXS4UDAP)](https://codecov.io/gh/atsutama2/go-grep)

高速で汎用性のあるテキスト検索ツール`go-grep`は、Goで書かれたコマンドラインツールで、ファイルやディレクトリ内のテキストを効率的に検索できます。
大文字小文字を区別せずにマッチングし、特定のファイルやディレクトリを除外できます。
---
`go-grep` is a fast and versatile text search tool written in Go as a command-line tool. It efficiently searches text within files and directories, matching without regard to case, and allowing exclusion of specific files or directories.

## Features

- 検索パターンを含む行番号を表示します。
- 検索パターンを含むファイルの相対パスを表示します。
- ディレクトリが指定されていない場合は現在のディレクトリから検索を開始します。
- 出力に検索パターンをハイライト表示します。
- Method名検索
- 指定されたディレクトリ内のファイルやサブディレクトリを再帰的に検索します。
- 検索中に、検索対象外とするファイル名やディレクトリ名を指定することができます。例えば、.git や vendor など。
- 検索中に検索対象外とされたファイルやディレクトリはスキップされます。
- 除外ファイルの格納場所は $HOME/go-grep/exclude_list.txt になります。
- つまり、ユーザーのホームディレクトリ内に go-grep ディレクトリがある必要があり、その中に exclude_list.txt ファイルを配置する必要があります。もしディレクトリが存在しない場合は、自分で作成する必要があります。
---
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
make build
```

Then, navigate to the project directory and build the `gg` binary:

```
gg <search_word> [<directory>]
```

## Usage

To use `gg`, execute the binary followed by the search pattern and the directory to start the search (optional). If no directory is specified, the search will start from the current directory.

```
gg <search_word> [<directory>]
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
![スクリーンショット 2023-04-03 18 25 46](https://user-images.githubusercontent.com/77330599/229469245-e5e40f71-263f-4f3c-9cfd-2144d9b50675.png)

The output will display the relative file path, line number, and the line containing the search pattern, with the search pattern highlighted.

## License
MIT
