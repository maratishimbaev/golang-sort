package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"
	"github.com/stretchr/testify/require"
)

var testSortInput = []string{"JS", "Go", "C++", "Go", "Python"}

var testSortResult = []string{"C++", "Go", "Go", "JS", "Python"}

var testSortWithOnlyUniqueResult = []string{"C++", "Go", "JS", "Python"}

var testReverseSortResult = []string{"Python", "JS", "Go", "Go", "C++"}

var testNumSortInput = []string{"3", "2", "5", "1", "4"}

var testNumSortResult = []string{"1", "2", "3", "4", "5"}

var testSortByColumnInput = []string{"zzz ddd", "yyy bbb", "www aaa", "qqq ccc"}

var testSortByColumnResult = []string{"www aaa", "yyy bbb", "qqq ccc", "zzz ddd"}

var testSortToFileInput = `zzz ddd
yyy bbb
www aaa
qqq ccc`

var testSortToFileResult = `www aaa
yyy bbb
qqq ccc
zzz ddd
`

func TestSort(t *testing.T) {
	sortOptions := SortOptions{
		IgnoreCase:  false,
		OnlyUnique:  false,
		IsReverse:   false,
		OutFileName: "",
		IsNum:       false,
		ColNum:      1,
	}

	require.Equal(t, sortStrings(testSortInput, sortOptions), testSortResult, "test sort failed")
}

func TestSortWithOnlyUnique(t *testing.T) {
	sortOptions := SortOptions{
		IgnoreCase:  false,
		OnlyUnique:  true,
		IsReverse:   false,
		OutFileName: "",
		IsNum:       false,
		ColNum:      1,
	}

	require.Equal(t, sortStrings(testSortInput, sortOptions), testSortWithOnlyUniqueResult, "test sort with only unique failed")
}

func TestReverseSort(t *testing.T) {
	sortOptions := SortOptions{
		IgnoreCase:  false,
		OnlyUnique:  false,
		IsReverse:   true,
		OutFileName: "",
		IsNum:       false,
		ColNum:      1,
	}

	require.Equal(t, sortStrings(testSortInput, sortOptions), testReverseSortResult, "test reverse sort failed")
}

func TestNumSort(t *testing.T) {
	sortOptions := SortOptions{
		IgnoreCase:  false,
		OnlyUnique:  false,
		IsReverse:   false,
		OutFileName: "",
		IsNum:       true,
		ColNum:      1,
	}

	require.Equal(t, sortStrings(testNumSortInput, sortOptions), testNumSortResult, "test num sort failed")
}

func TestSortByColumn(t *testing.T) {
	sortOptions := SortOptions{
		IgnoreCase:  false,
		OnlyUnique:  false,
		IsReverse:   false,
		OutFileName: "",
		IsNum:       false,
		ColNum:      2,
	}
	require.Equal(t, sortStrings(testSortByColumnInput, sortOptions), testSortByColumnResult, "test sort by column failed")
}

func TestSortToFile(t *testing.T) {
	in := bytes.NewBufferString(testSortToFileInput)

	sortOptions := SortOptions{
		IgnoreCase:  false,
		OnlyUnique:  false,
		IsReverse:   false,
		OutFileName: "log.txt",
		IsNum:       false,
		ColNum:      2,
	}

	unixSort(in, sortOptions)

	outFile, err := os.Open("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	var result string
	scanner := bufio.NewScanner(outFile)
	for scanner.Scan() {
		result = result + scanner.Text() + "\n"
	}

	require.Equal(t, result, testSortToFileResult, "test sort to file failed")
}
