package main

import "fmt"

func run2156() {
	inputCount := int32(0)
	fmt.Scan(&inputCount)

	inputNumbers := make([]int64, inputCount+1)
	for i:=int32(1); i<=inputCount; i++ {
		input := int64(0)
		fmt.Scan(&input)
		inputNumbers[i] = input
	}

	switch inputCount {
	case 1:
		fmt.Println(inputNumbers[1])
		return
	case 2:
		fmt.Println(inputNumbers[2] + inputNumbers[1])
		return
	default:
	}

	cacheI := make([]int64, inputCount+1) // include n-1
	cacheE := make([]int64, inputCount+1) // exclude n-1
	cache0 := make([]int64, inputCount+1) // exclude n-1

	{
		one := inputNumbers[1]
		two := inputNumbers[2]

		cacheI[1] = one
		cacheE[1] = one
		cache0[1] = 0

		cacheI[2] = one + two
		cacheE[2] = two
		cache0[2] = 0
	}

	for i:=int32(3); i<=inputCount; i++ {
		number := inputNumbers[i]
		{
			index := i-1
			cacheI[i] = max(cacheE[index], cache0[index]) + number

		}
		{
			index := i-2
			cacheE[i] = max(max(cacheI[index], cacheE[index]), cache0[index]) + number
		}
		{
			index := i-3
			cache0[i] = max(max(cacheI[index], cacheE[index]), cache0[index]) + number
		}
	}
	fmt.Println(max(
		max(max(cacheE[inputCount], cacheI[inputCount]), cache0[inputCount]),
		max(max(cacheE[inputCount-1], cacheI[inputCount-1]), cache0[inputCount-1])))
}

func max(a,b int64) int64 {
	if a>b {
		return a
	} else {
		return b
	}
}
