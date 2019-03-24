package main

import "fmt"

func run11727(){
	input := 0
	fmt.Scan(&input)
	fmt.Println(run11727Func(input))
}

func run11727Func(input int) uint64 {

	switch input {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 3
	case 3:
		return 5
	}

	cache := make([]uint64, input)
	cache[0] = 1
	cache[1] = 3
	cache[2] = 5
	for i:=3; i<input; i++ {
		cache[i] = (cache[i-1]%10007 + cache[i-2]*2%10007)%10007
	}
	return cache[input-1]
}