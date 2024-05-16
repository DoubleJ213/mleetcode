package fxxk

import (
	"fmt"
	"testing"
)

/*
264. 丑数 II
给你一个整数 n ，请你找出并返回第 n 个 丑数 。
丑数 就是质因子只包含 2、3 和 5 的正整数。

示例 1：
输入：n = 10
输出：12
解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
示例 2：
输入：n = 1
输出：1
解释：1 通常被视为丑数。

提示：
1 <= n <= 1690
*/

var dp264 []int

// dp264[i] 表示第i个丑数
func nthUglyNumber(n int) int {
	dp264 = make([]int, n)
	dp264[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		dp264[i] = getMin(getMin(dp264[i2]*2, dp264[i3]*3), dp264[i5]*5)
		if dp264[i] == dp264[i2]*2 {
			//如果当前的丑数是由*2得到的，i2的索引进一位
			i2++
		}

		if dp264[i] == dp264[i3]*3 {
			i3++
		}
		if dp264[i] == dp264[i5]*5 {
			i5++
		}

	}
	return dp264[n-1]
}

func TestAl264(t *testing.T) {
	fmt.Println(nthUglyNumber(10))
}
