package main

import (
	"fmt"
)

/*
2×n 크기의 직사각형을 1×2, 2×1 타일로 채우는 방법의 수를 구하는 프로그램을 작성하시오.

2 : 2
3 : 3 = (n(2)-1)*2 + 1
4 : 5
5 : 8
6 : 13

7 : 21
8 : 34
9 : 55

1증가하여 홀수라면 기존의 2배
1증가하여 짝수라면
*/

func run11726() {
	var input int32
	fmt.Scan(&input)

	cache := make(map[int32]uint64)
	cache[1] = 1
	cache[2] = 2
	cache[3] = 3
	for i:=int32(4); i<=input; i++ {
		cache[i] = (cache[i-1]%10007 + cache[i-2]%10007)%10007
		delete(cache, i-2)
	}
	fmt.Println(cache[input])
}
