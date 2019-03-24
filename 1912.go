package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// std in 을 포함한 함수
func run1912() {
	inputCount := 0
	fmt.Scan(&inputCount)
	reader := bufio.NewReader(os.Stdin)

	stdInString, _ := reader.ReadString('\n')
	fmt.Println(run1912Func(inputCount, stdInString))
}

func run1912Func(inputCount int, stdInString string) int64 {
	// 입력 받은 정수들의 배열
	inputIntArray := make([]int, inputCount)
	stdInString = strings.Replace(stdInString, "\n", "", -1)
	stdInStringArray := strings.Split(stdInString, " ")

	maxInput := int64(math.MinInt64)

	for j:=0; j<len(stdInStringArray); j++ {
		integer, _ := strconv.Atoi(stdInStringArray[j])
		inputIntArray[j] = integer
		if maxInput < int64(integer) {
			maxInput = int64(integer)
		}
	}

	plusIntArray := make([]int64, 0)
	sum := int64(0)
	for _, num := range inputIntArray {
		if sum != 0 && ((num > 0) != (sum > 0)) {
			plusIntArray = append(plusIntArray, sum)
			sum = 0
		}
		sum += int64(num)
	}

	plusIntArray = append(plusIntArray, sum)

	maxVal := int64(math.MinInt64)
	currentSum := int64(0)
	for i:=0; i<len(plusIntArray); i++ {
		currentInt := plusIntArray[i]
	//for i:=0; i<len(inputIntArray); i++ {
	//	currentInt := inputIntArray[i]
		tempSum := currentSum + int64(currentInt)
		if tempSum > 0 {
			currentSum = tempSum
			maxVal = maxInt64(currentSum, maxVal)
		} else {
			currentSum = 0
		}
	}

	maxVal = maxInt64(maxVal, maxInput)

	return maxVal
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}