package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"
	"github.com/stretchr/testify/require"
)

var testSortInput = `JS
Go
C++
Go
Python`

var testSortResult = `C++
Go
Go
JS
Python
`

var testSortWithOnlyUniqueResult = `C++
Go
JS
Python
`

var testReverseSortResult = `Python
JS
Go
Go
C++
`

var testNumSortInput = `3
2
5
1
4`

var testNumSortResult = `1
2
3
4
5
`

var testSortByColumnInput = `zzz ddd
yyy bbb
www aaa
qqq ccc`

var testSortByColumnResult = `www aaa
yyy bbb
qqq ccc
zzz ddd
`

func TestSort(t *testing.T) {
	in := bytes.NewBufferString(testSortInput)
	out := bytes.NewBuffer(nil)
	sortStrings(in, out,
				false, false, false,
				"", false, 1)

	result := out.String()
	require.Equal(t, result, testSortResult, "test sort failed")
}

func TestSortWithOnlyUnique(t *testing.T) {
	in := bytes.NewBufferString(testSortInput)
	out := bytes.NewBuffer(nil)
	sortStrings(in, out,
		false, true, false,
		"", false, 1)

	result := out.String()
	require.Equal(t, result, testSortWithOnlyUniqueResult, "test sort with only unique failed")
}

func TestReverseSort(t *testing.T) {
	in := bytes.NewBufferString(testSortInput)
	out := bytes.NewBuffer(nil)
	sortStrings(in, out,
		false, false, true,
		"", false, 1)

	result := out.String()
	require.Equal(t, result, testReverseSortResult, "test reverse sort failed")
}

func TestNumSort(t *testing.T) {
	in := bytes.NewBufferString(testNumSortInput)
	out := bytes.NewBuffer(nil)
	sortStrings(in, out,
		false, false, false,
		"", true, 1)

	result := out.String()
	require.Equal(t, result, testNumSortResult, "test num sort failed")
}

func TestSortByColumn(t *testing.T) {
	in := bytes.NewBufferString(testSortByColumnInput)
	out := bytes.NewBuffer(nil)
	sortStrings(in, out,
		false, false, false,
		"", false, 2)

	result := out.String()
	require.Equal(t, result, testSortByColumnResult, "test sort by column failed")
}

func TestSortToFile(t *testing.T) {
	in := bytes.NewBufferString(testSortByColumnInput)
	out := bytes.NewBuffer(nil)
	sortStrings(in, out,
		false, false, false,
		"log.txt", false, 2)

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

	require.Equal(t, result, testSortByColumnResult, "test sort to file failed")
}
