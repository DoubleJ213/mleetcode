package fxxk

import (
	"fmt"
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)

var ans89 []int

// n 位格雷码序列 是一个由 2^n 个整数组成的序列
// 给你一个整数 n ，返回任一有效的 n 位格雷码序列
// 方法一 1 <= n <= 16 直接穷举
func grayCode(n int) []int {
	pre := []int{0, 1}
	if n == 1 {
		return pre
	}
	for i := 2; i <= n; i++ {
		current := make([]int, len(pre))
		copy(current, pre)
		for j := len(pre) - 1; j >= 0; j-- {
			current = append(current, pre[j]+1<<(i-1))
		}
		pre = current
	}
	return pre
}

//leetcode submit region end(Prohibit modification and deletion)

func TestGrayCode(t *testing.T) {
	//fmt.Println(grayCode(1))
	//fmt.Println(grayCode(2))
	fmt.Println(grayCode(3))
}
