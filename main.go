package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortStrings(input io.Reader, output io.Writer,
				 ignoreCase, onlyUnique, isReverse bool,
				 outFileName string, isNum bool, colNum int) {
	// create string slice
	strList := make([]string, 0, 10)

	// add string to slice
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		strList = append(strList, scanner.Text())
	}

	// sort
	sort.Slice(strList, func(i, j int) bool {
		var isLess bool

		if isNum {
			firstNum, _ := strconv.Atoi(strList[i])
			secondNum, _ := strconv.Atoi(strList[j])

			isLess = firstNum <= secondNum
		} else {
			firstStr := strings.Split(strList[i], " ")[colNum - 1]
			secondStr := strings.Split(strList[j], " ")[colNum - 1]

			if ignoreCase {
				firstStr = strings.ToLower(firstStr)
				secondStr = strings.ToLower(secondStr)
			}

			isLess = firstStr <= secondStr
		}

		if isReverse {
			return !isLess
		}
		return isLess
	})

	var prevStr string
	isDiffWithIgnoreCase := func(firstStr, secondStr string) bool {
		return strings.ToLower(firstStr) != strings.ToLower(secondStr)
	}

	outFile, _ := os.Create(outFileName)

	// print sort strings
	for _, str := range strList {
		if !onlyUnique || (prevStr != str && (!ignoreCase || isDiffWithIgnoreCase(prevStr, str))) {
			if outFileName != "" {
				_, _ = outFile.WriteString(str + "\n")
			} else {
				_, _ = fmt.Fprintln(output, str)
			}
		}
		prevStr = str
	}

	// print errors
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	// get flags
	ignoreCase := flag.Bool("f", false, "ignore case")
	onlyUnique := flag.Bool("u", false, "only unique")
	isReverse := flag.Bool("r", false, "reverse sort")
	outFileName := flag.String("o", "", "output to inFile")
	isNum := flag.Bool("n", false, "numeric sort")
	colNum := flag.Int("k", 1, "sort by column")

	flag.Parse()

	// get file name
	fileName := flag.Arg(0)

	// open file
	inFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer inFile.Close()

	sortStrings(bufio.NewReader(inFile), os.Stdout,
				*ignoreCase, *onlyUnique, *isReverse,
				*outFileName, *isNum, *colNum)
}
