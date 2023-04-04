package grep_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/atsutama2/go-grep/pkg/grep"
)

func createTestFiles(basePath string) {
	os.MkdirAll(filepath.Join(basePath, "testdir"), os.ModePerm)

	testFiles := map[string]string{
		"testfile1.txt": `This is a test file.
func myFunction() {
}`,
		"testfile2.txt": `This is another test file.
func anotherFunction() {
}`,
		"testdir/testfile3.txt": `This is a test file in a subdirectory.
func myFunction() {
}`,
	}

	for path, content := range testFiles {
		fullPath := filepath.Join(basePath, path)
		err := ioutil.WriteFile(fullPath, []byte(content), os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("Failed to create test file: %v", err))
		}
	}
}

func TestGrep(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "go-grep-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	createTestFiles(tempDir)

	tests := []struct {
		name       string
		searchWord string
		classMode  bool
		want       []string
	}{
		{
			name:       "search myFunction",
			searchWord: "myFunction",
			classMode:  false,
			want: []string{
				"testfile1.txt",
				"testdir/testfile3.txt",
			},
		},
		{
			name:       "search anotherFunction",
			searchWord: "anotherFunction",
			classMode:  false,
			want: []string{
				"testfile2.txt",
			},
		},
		{
			name:       "search myFunction in class mode",
			searchWord: "myFunction",
			classMode:  true,
			want: []string{
				"testfile1.txt",
				"testdir/testfile3.txt",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := make([]string, 0)

			err := grep.Grep(test.searchWord, tempDir, test.classMode, func(format string, a ...interface{}) (n int, err error) {
				if strings.Contains(format, "%s") {
					results = append(results, a[0].(string))
				}
				return 0, nil
			})
			if err != nil {
				t.Fatalf("Error running Grep: %v", err)
			}

			if len(results) != len(test.want) {
				t.Errorf("Search results count mismatch. Got %d, want %d", len(results), len(test.want))
			}

			for i, res := range results {
				if res != test.want[i] {
					t.Errorf("Mismatch in search results. Got %s, want %s", res, test.want[i])
				}
			}
		})
	}
}
