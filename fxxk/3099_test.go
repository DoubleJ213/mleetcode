package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func sumOfTheDigitsOfHarshadNumber(x int) int {
	sum := 0
	cur := x
	for cur > 0 {
		sum += cur % 10
		cur = cur / 10
	}
	if x%sum != 0 {
		return -1
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)

func TestHarshadNumber(t *testing.T) {
	fmt.Println(sumOfTheDigitsOfHarshadNumber(18))
	//fmt.Println(sumOfTheDigitsOfHarshadNumber(23))
}
