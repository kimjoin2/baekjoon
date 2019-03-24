package main

import "fmt"

/*
n번째 최대값은 n-1, n-2 에서 n번째 값을 더한 값 중에서 최대값
여기서 n-1 값이 n-2 값을 포함할 경우에는 사용할 수 없음(3번 연속된 값을 더할 수 없음)

n번 째의 최대값은 두가지 패턴으로 계산해나가야할 것 같음
n-1 을 포함하는 최대값과 : i(n)
n-1 을 포함하지 않는 최대값 : e(n)
n번 째의 최대값 : max((n) + e(n-1), (n) + i(n-2), (n) + e(n-2))

i(n) = e(n-1) + n
e(n) = max(i(n-2), e(n-2)) + n
 */

func run2579() {
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

	switch inputCount {
	case 1:
		fmt.Println(inputNumbers[1])
		return
	case 2:
		fmt.Println(inputNumbers[2] + inputNumbers[1])
		return
	default:
	}

	cacheI := make(map[int32]int32) // include n-1
	cacheE := make(map[int32]int32) // exclude n-1

	{
		one := inputNumbers[1]
		two := inputNumbers[2]

		cacheI[1] = one
		cacheE[1] = one

		cacheI[2] = one + two
		cacheE[2] = two
	}

	/*
	i(n) = e(n-1) + n
	e(n) = max(i(n-2), e(n-2)) + n
	*/
	for i:=int32(3); i<=inputCount; i++ {
		number := inputNumbers[i]
		cacheI[i] = cacheE[i-1] + number

		{
			index := i-2
			in := cacheI[index]
			en := cacheE[index]

			if in > en {
				cacheE[i] = in + number
			} else {
				cacheE[i] = en + number
			}
		}
	}

	if cacheE[inputCount] > cacheI[inputCount] {
		fmt.Println(cacheE[inputCount])
	} else {
		fmt.Println(cacheI[inputCount])
	}
}
/*
6
10
20
15
25
10
20
*/