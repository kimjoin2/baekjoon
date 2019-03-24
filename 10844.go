package main

import (
	"fmt"
)

// std in 을 포함한 함수
func run10844() {
	input := 0
	fmt.Scan(&input)
	fmt.Println(run10844Func(input))
}

func run10844Func(input int) uint64 {
	const div uint64 = 1000000000

	numberCount := [10]uint64{0,1,1,1,1,1,1,1,1,1}

	for i:=2; i<=input; i++ {
		var tempNumberCount [10]uint64

		for j:=1; j<=8; j++ {
			tempNumberCount[j-1] += numberCount[j] % div
			tempNumberCount[j+1] += numberCount[j] % div
		}

		{
			tempNumberCount[1] += numberCount[0] % div
			tempNumberCount[8] += numberCount[9] % div
		}
		numberCount = tempNumberCount
	}

	count := uint64(0)
	for i:=0; i<10; i++ {
		count += numberCount[i] % div
	}

	return count % div
}

