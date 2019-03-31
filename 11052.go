package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func run11052(){
	inputCount := 0
	fmt.Scan(&inputCount)
	reader := bufio.NewReader(os.Stdin)

	stdInString, _ := reader.ReadString('\n')
	fmt.Println(run11052Func(inputCount, stdInString))
}

func run11052Func(inputCount int, stdInString string) int {
	inputIntArray := make([]int, inputCount)
	stdInString = strings.Replace(stdInString, "\n", "", -1)
	stdInStringArray := strings.Split(stdInString, " ")

	for j:=0; j<inputCount; j++ {
		integer, _ := strconv.Atoi(stdInStringArray[j])
		inputIntArray[j] = integer
	}

	rateArray := make([]float64, inputCount)
	for i:=0; i<inputCount; i++ {
		rateArray[i] = float64(inputIntArray[i]) / float64(i+1)
	}

	money := 0

	cardCount := 0
	for cardCount < inputCount {
		remainCardCount := inputCount - cardCount
		maxRate := 0.0
		maxRateIndex := 0
		for index, rate := range rateArray {
			if maxRate < rate && remainCardCount >= index+1 {
				maxRate = rate
				maxRateIndex = index
			}
		}
		currentBuyCount := remainCardCount / (maxRateIndex+1)
		cardCount += currentBuyCount * (maxRateIndex+1)
		money += currentBuyCount * inputIntArray[maxRateIndex]
	}

	return money
}

func main() {
	run11052()
}
