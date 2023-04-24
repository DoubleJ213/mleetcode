package fxxk

import (
	"fmt"
	"testing"
)

//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//你可以按 任何顺序 返回答案。

//示例 1：
//输入：n = 4, k = 2
//输出：[[2,4], [3,4], [2,3], [1,2], [1,3], [1,4]]

//示例 2：
//输入：n = 1, k = 1  输出：[[1]]

//提示：
//1 <= n <= 20  1 <= k <= n

func combine(n int, k int) [][]int {
	combineRes := make([][]int, 0)
	combineDfs(1, n, k, make([]int, 0), &combineRes)
	return combineRes
}

func combineDfs(start, target, k int, levelRes []int, res *[][]int) {
	if k == 0 {
		*res = append(*res, append([]int{}, levelRes...))
		return
	}
	for i := start; i <= target; i++ {
		combineDfs(i+1, target, k-1, append(levelRes, i), res)
		//需要深刻理解递归回溯 和 循环时，变量变化规律
		//combineDfs(start+1, target, k-1, append(levelRes, i), res)
		fmt.Print(i)
	}
}

func TestAl77(t *testing.T) {
	a := combine(4, 2)
	fmt.Print("abc")
	_ = a
}
