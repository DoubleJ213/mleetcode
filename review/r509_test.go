package review

import (
	"fmt"
	"testing"
)

func TestAl509(t *testing.T) {
	//fmt.Printf("5 --- %v\n", fib(5))
	fmt.Printf("5 --- %v\n", fib1(5))
	fmt.Printf("5 --- %v\n", fib2(5))
}

func fib2(i int) int {
	//带备忘录的递归
	return 1
}

var res509 []int

func fib1(input int) int {
	//自底向上 DP table
	if input == 1 {
		return 1
	}
	if input == 2 {
		return 1
	}
	res509 = make([]int, input+1)
	res509[1] = 1
	res509[2] = 1
	for i := 3; i <= input; i++ {
		res509[i] = res509[i-1] + res509[i-2]
	}
	return res509[input]
}

//fib(1)= 1
//fib(2)= 1
func fib(input int) int {
	//递归
	if input == 1 {
		return 1
	}
	if input == 2 {
		return 1
	}
	return fib(input-1) + fib(input-2)
}
