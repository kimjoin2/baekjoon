package main

import (
	"fmt"
)

/*
RGB 세가지의 색칠 비용
n번째 집까지의 비용은 n-1번째 집의 상태에 따라 다름(R? G? B?)
따라서 3가지 상태를 동시에 저장해둬야함

n에서 r, g, b인 경우를 각각
cacheR, cacheG, cacheB 에 저장

cacheR[n] = max(cacheG[n-1], cacheB[n-1]) + R(n)
 */

func main() {
	inputCount := int32(0)
	fmt.Scan(&inputCount)

	inputRGBs := make([]rgb, inputCount)
	for i:=int32(0); i<inputCount; i++ {
		var input rgb
		fmt.Scan(&(input.r), &(input.g), &(input.b))
		inputRGBs[i] = input
	}

	cache := make([]rgb, inputCount)

	cache[0] = inputRGBs[0]

	for i:=int32(1); i<inputCount; i++ {
		preData := cache[i-1]
		data := inputRGBs[i]
		newData := cache[i]
		newData.r = min2(preData.g, preData.b) + data.r
		newData.g = min2(preData.r, preData.b) + data.g
		newData.b = min2(preData.g, preData.r) + data.b
		cache[i] = newData
	}
	resData := cache[inputCount-1]
	fmt.Println(min3(resData.r, resData.g, resData.b))
}

type rgb struct {
	r int32
	g int32
	b int32
}

func min3(a,b,c int32) int32 {
	res := min2(min2(a, b), c)
	return res
}

func min2(a,b int32) int32 {
	res := a
	if res > b {
		res = b
	}
	return res
}
/*
3
26 40 83
49 60 57
13 89 99
 */