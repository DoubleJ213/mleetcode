package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func missingRolls(rolls []int, mean int, n int) []int {
	// 筛子，数值1~6
	// 如果存在多组符合要求的答案，只需要返回其中任意一组即可
	ans := make([]int, n)
	sum := 0
	for i := 0; i < len(rolls); i++ {
		sum += rolls[i]
	}
	need := len(rolls)*mean - sum
	if mean*n+need < n || mean*n+need > 6*n {
		return nil
	}
	multiple, remain := getSubNums(need, n)
	subStep := 1
	if remain < 0 {
		subStep = -1
		remain = 0 - remain
	}
	for j := 0; j < n; j++ {
		if j < remain {
			ans[j] = mean + multiple + subStep
		} else {
			ans[j] = mean + multiple
		}

	}
	return ans
}

/*
need 是 n 的 倍数，余数多少
比如need 5 n =4
*/
func getSubNums(need int, n int) (int, int) {
	if need < 0 {
		a, b := getSubNums(0-need, n)
		return -a, -b
	}

	return need / n, need % n
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindMissingObservations(t *testing.T) {
	// mean 平均值 n 缺几个数
	fmt.Println(missingRolls([]int{1, 5, 6}, 5, 2))
}
