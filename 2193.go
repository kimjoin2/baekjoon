package main

import "fmt"

func run2193() {

	input := int8(0)
	fmt.Scan(&input)

	// init setting : 1
	prev0 := uint64(0)
	prev1 := uint64(1)
	prevCount := uint64(1)

	for i:=int8(2); i<=input; i++ {
		prevCount = prev1 + prev0*2
		tempPrev1 := prev1
		prev1 = prev0
		prev0 = tempPrev1 + prev0
	}

	fmt.Println(prevCount)
}
