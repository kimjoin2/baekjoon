package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func run11053(){
	inputCount := 0
	fmt.Scan(&inputCount)
	reader := bufio.NewReader(os.Stdin)

	stdInString, _ := reader.ReadString('\n')
	fmt.Println(run11053Func(inputCount, stdInString))
}

func run11053Func(inputCount int, stdInString string) int {
	inputIntArray := make([]int, inputCount)
	stdInString = strings.Replace(stdInString, "\n", "", -1)
	stdInStringArray := strings.Split(stdInString, " ")

	for j:=0; j<len(stdInStringArray); j++ {
		integer, _ := strconv.Atoi(stdInStringArray[j])
		inputIntArray[j] = integer
	}

	result := 0
	cache := make([]int, inputCount)
	cache[0] = 1
	for i:=1; i<inputCount; i++ {
		number := inputIntArray[i]

		for j:=0; j<i; j++ {
			beforeNumber := inputIntArray[j]
			if beforeNumber < number && cache[i] < cache[j]+1 {
				cache[i] = cache[j]+1
			}
		}
		if cache[i] == 0 {
			cache[i] = 1
		}
	}

	for _, val := range cache {
		if result < val {
			result = val
		}
	}

	return result
}

func main() {
	run11053()
}