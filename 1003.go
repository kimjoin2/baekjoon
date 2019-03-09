package main

import "fmt"

func main() {
	inputCount := int32(0)
	fmt.Scan(&inputCount)

	inputNumbers := make([]int32, inputCount+1)
	maxNum := int32(0)
	for i:=int32(1); i<=inputCount; i++ {
		input := int32(0)
		fmt.Scan(&input)
		inputNumbers[i] = input
		if input > maxNum {
			maxNum = input
		}
	}

	cache0 := make(map[int32]int32)
	cache0[0] = 1
	cache0[1] = 0

	cache1 := make(map[int32]int32)
	cache1[0] = 0
	cache1[1] = 1

	for i:=int32(2); i<=maxNum; i++ {
		cache0[i] = cache0[i-1] + cache0[i-2]
		cache1[i] = cache1[i-1] + cache1[i-2]
	}

	for i:=int32(1); i<=inputCount; i++ {
		index := inputNumbers[i]
		fmt.Println(cache0[index], cache1[index])
	}
}
