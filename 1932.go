package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func run1932() {
	rows := 0
	fmt.Scan(&rows)
	//rows = 3

	inputData := make([][]int64, rows)
	reader := bufio.NewReader(os.Stdin)

	for i:=0; i<rows; i++ {
		inputRowData := make([]int64, i+1)
		inputData[i] = inputRowData
		stringInts, _ := reader.ReadString('\n')
		stringInts = strings.Replace(stringInts, "\n", "", -1)
		stringIntsArray := strings.Split(stringInts, " ")
		for j:=0; j<=i; j++ {
			integer, _ := strconv.Atoi(stringIntsArray[j])
			inputRowData[j] = int64(integer)
		}
	}

	for i:=rows-1; i>0; i-- {
		currRow := inputData[i]
		upperRow := inputData[i-1]

		for j:=0; j<len(upperRow); j++ {
			upperVal := upperRow[j]
			optionLeft := currRow[j]
			optionRight := currRow[j+1]

			if optionLeft > optionRight {
				upperVal += optionLeft
			} else {
				upperVal += optionRight
			}

			upperRow[j] = upperVal
		}
	}

	fmt.Println(inputData[0][0])
}
