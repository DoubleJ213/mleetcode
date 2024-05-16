package fxxk

import (
	"fmt"
	"testing"
	"time"
)

// fib(0) =0 fib(1) =1 fib(2) =1
// fib(n) = fib(N - 1) + fib(N - 2);
var res509 []int

func fib(x int) int {
	//num509 = make([]int, x+1)
	//num509[0] = 0
	//num509[1] = 1
	//num509[2] = 1
	if x == 0 {
		return 0
	}
	if x == 1 || x == 2 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}

func fib2(x int) int {
	res509 = make([]int, x+1)
	if x < 0 {
		return 0
	}
	return dp509(x)
}

func dp509(x int) int {
	if x == 0 {
		res509[0] = 0
		return res509[0]
	}
	if x == 1 || x == 2 {
		res509[1] = 1
		res509[2] = 1
		return 1
	}
	if res509[x] != 0 {
		//如果没有这个，res509[x]的值将被重复计算
		return res509[x]
	}
	res509[x] = dp509(x-1) + dp509(x-2)
	return res509[x]
}

func fib3(x int) int {
	num509 := make([]int, x+1)
	num509[0] = 0
	num509[1] = 1
	num509[2] = 1
	for i := 3; i <= x; i++ {
		num509[i] = num509[i-1] + num509[i-2]
	}
	return res509[x]
}

func TestAl509(t *testing.T) {
	a := time.Now()
	fmt.Printf("%d\n", fib2(30))
	fmt.Printf("%v\n", time.Since(a))

	b := time.Now()
	fmt.Printf("%d\n", fib(30))
	fmt.Printf("%v\n", time.Since(b))

	c := time.Now()
	fmt.Printf("%d\n", fib3(30))
	fmt.Printf("%v\n", time.Since(c))
}
