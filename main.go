package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SortOptions struct {
	IgnoreCase bool
	OnlyUnique bool
	IsReverse bool
	OutFileName string
	IsNum bool
	ColNum int
}

func isLessString(firstStr, secondStr string,
				  isNum, ignoreCase, isReverse bool,
				  colNum int) (bool, error) {
	var isLess bool

	if isNum {
		firstNum, firstErr := strconv.Atoi(firstStr)
		secondNum, secondErr := strconv.Atoi(secondStr)

		if firstErr != nil {
			return false, firstErr
		}
		if secondErr != nil {
			return false, secondErr
		}

		isLess = firstNum <= secondNum
	} else {
		firstStr := strings.Split(firstStr, " ")[colNum - 1]
		secondStr := strings.Split(secondStr, " ")[colNum - 1]

		if ignoreCase {
			firstStr = strings.ToLower(firstStr)
			secondStr = strings.ToLower(secondStr)
		}

		isLess = firstStr <= secondStr
	}

	if isReverse {
		return !isLess, nil
	}
	return isLess, nil
}

func readStrings(input io.Reader) ([]string, error) {
	// create string slice
	strList := make([]string, 0)

	// add string to slice
	strScanner := bufio.NewScanner(input)
	for strScanner.Scan() {
		strList = append(strList, strScanner.Text())
	}

	if err := strScanner.Err(); err != nil {
		return nil, err
	}

	return strList, nil
}

func sortStrings(strList []string, sortOptions SortOptions) []string {
	sort.Slice(strList, func(i, j int) bool {
		isLess, err := isLessString(strList[i], strList[j],
			sortOptions.IsNum, sortOptions.IgnoreCase, sortOptions.IsReverse,
			sortOptions.ColNum)
		if err != nil {
			log.Fatal(err)
		}

		return isLess
	})

	var prevStr string
	isDiffWithIgnoreCase := func(firstStr, secondStr string) bool {
		return strings.ToLower(firstStr) != strings.ToLower(secondStr)
	}

	var sortStrList []string

	for _, str := range strList {
		if !sortOptions.OnlyUnique || (prevStr != str && (!sortOptions.IgnoreCase || isDiffWithIgnoreCase(prevStr, str))) {
			sortStrList = append(sortStrList, str)
		}
		prevStr = str
	}

	return sortStrList
}

func writeStrings(strList []string, outFileName string) error {
	output := os.Stdout
	if outFileName != "" {
		fileOutput, err := os.Create(outFileName)
		output = fileOutput
		if err != nil {
			return err
		}
	}

	for _, str := range strList {
		_, _ = fmt.Fprintln(output, str)
	}

	return nil
}

func unixSort(input io.Reader, sortOptions SortOptions) error {
	strList, readErr := readStrings(input)
	if readErr != nil {
		return readErr
	}

	strList = sortStrings(strList, sortOptions)

	writeErr := writeStrings(strList, sortOptions.OutFileName)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

func main() {
	sortOptions := SortOptions{
		IgnoreCase:  *flag.Bool("f", false, "ignore case"),
		OnlyUnique:  *flag.Bool("u", false, "only unique"),
		IsReverse:   *flag.Bool("r", false, "reverse sort"),
		OutFileName: *flag.String("o", "", "output to inFile"),
		IsNum:       *flag.Bool("n", false, "numeric sort"),
		ColNum:      *flag.Int("k", 1, "sort by column"),
	}

	flag.Parse()

	// get file name
	fileName := flag.Arg(0)

	// open file
	inFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer inFile.Close()

	err = unixSort(bufio.NewReader(inFile), sortOptions)
	if err != nil {
		log.Fatal(err)
	}
}
