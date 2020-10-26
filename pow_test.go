package mleetcode

import (
	"fmt"
	"testing"
)

//50. Pow(x, n)
//n 为有符号整数
//递归、位移
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n%2 == 0 {
		// 偶数
		//if n&1 == 0 {
		//x = x&(x-1) 最低位清0
		//x&-x 得到最低位的1
		return myPow(x*x, n/2)
	}
	return x * myPow(x*x, n/2)
}

func myPow1(x float64, n int) float64 {
	return 0
}

func TestPow(t *testing.T) {
	fmt.Printf("%v", myPow(2, 10))
	fmt.Printf("%v", myPow1(2, 10))
}
