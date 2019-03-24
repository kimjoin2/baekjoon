package main

import (
	"fmt"
)

// max number for input
const maxNum int32 = 1000000
const divide2 int32 = 2
const divide3 int32 = 3
const minus1 int32 = 1

// main function for BAEKJOON question number 1463
func run1463() {

	// var for input integer
	var input int32
	var count int32
	// get input(integer)
	fmt.Scan(&input)

	cache := make(map[int32]int32)
	cache[1] = 0
	cache[2] = 1
	cache[3] = 1
	for i:=4; i<=int(input); i++ {
		var int32i = int32(i)
		var localPreCount = int32(maxNum)
		if int32i % divide2 == 0 {
			localPreCount = cache[int32i/divide2]
		}
		if int32i % divide3 == 0 {
			res := int32i / divide3
			if localPreCount > cache[res] {
				localPreCount = cache[res]
			}
		}
		{
			res := int32i-minus1
			if localPreCount > cache[res] {
				localPreCount = cache[res]
			}
		}
		cache[int32i] = localPreCount+1
	}
	count = cache[input]
	fmt.Println(count)
}
