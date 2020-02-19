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

// get flags
var ignoreCase = flag.Bool("f", false, "ignore case")
var onlyUnique = flag.Bool("u", false, "only unique")
var isReverse = flag.Bool("r", false, "reverse sort")
var outFileName = flag.String("o", "", "output to inFile")
var isNum = flag.Bool("n", false, "numeric sort")
var colNum = flag.Int("k", 1, "sort by column")

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

func sortStrings(input io.Reader,
				 ignoreCase, onlyUnique, isReverse bool,
				 outFileName string, isNum bool, colNum int) error {
	// create string slice
	strList := make([]string, 0)

	// add string to slice
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		strList = append(strList, scanner.Text())
	}

	// sort
	sort.Slice(strList, func(i, j int) bool {
		isLess, err := isLessString(strList[i], strList[j],
							isNum, ignoreCase, isReverse,
							colNum)
		if err != nil {
			log.Fatal(err)
		}

		return isLess
	})

	var prevStr string
	isDiffWithIgnoreCase := func(firstStr, secondStr string) bool {
		return strings.ToLower(firstStr) != strings.ToLower(secondStr)
	}

	output := os.Stdout
	if outFileName != "" {
		fileOutput, err := os.Create(outFileName)
		output = fileOutput
		if err != nil {
			return err
		}
	}

	// print sort strings
	for _, str := range strList {
		if !onlyUnique || (prevStr != str && (!ignoreCase || isDiffWithIgnoreCase(prevStr, str))) {
			_, _ = fmt.Fprintln(output, str)
		}
		prevStr = str
	}

	// print errors
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()

	// get file name
	fileName := flag.Arg(0)

	// open file
	inFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer inFile.Close()

	err = sortStrings(bufio.NewReader(inFile),
				*ignoreCase, *onlyUnique, *isReverse,
				*outFileName, *isNum, *colNum)
	if err != nil {
		log.Fatal(err)
	}
}
