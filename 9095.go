// this file is for Q.9095
package main

import "fmt"

func run9095() {
	inputCount := int32(0)
	fmt.Scan(&inputCount)

	inputNumbers := make([]int32, inputCount)
	maxNum := int32(0)
	for i:=int32(0); i<inputCount; i++ {
		input := int32(0)
		fmt.Scan(&input)
		inputNumbers[i] = input
		if input > maxNum {
			maxNum = input
		}
	}

	cache := make(map[int32]int32)
	cache[1] = 1
	cache[2] = 2
	cache[3] = 4
	for i:=int32(4); i<=maxNum; i++ {
		cache[i] = cache[i-1] + cache[i-2] + cache[i-3]
	}

	for i:=int32(0); i<inputCount; i++ {
		fmt.Println(cache[inputNumbers[i]])
	}
}
