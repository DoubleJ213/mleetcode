package fxxk

import (
	"fmt"
	"testing"
)

/*
1020. 飞地的数量
给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。
一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。
返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量
输入：grid = [
[0,0,0,0],
[1,0,1,0],
[0,1,1,0],
[0,0,0,0]]
输出：3
解释：有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。

示例 2：
输入：grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
输出：0
解释：所有 1 都在边界上或可以到达边界。

提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 500
grid[i][j] 的值为 0 或 1
*/

/*
 0 表示一个海洋单元格、1 表示一个陆地单元格
算的是飞地的数量
*/

var sts1020 bool
var total1020 int

// 数量面积一样，先这么定义了
var area1020 int

func numEnclaves(grid [][]int) int {
	total1020 = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				sts1020 = true
				area1020 = 0
				dfs1020(grid, i, j)
				if sts1020 {
					total1020 += area1020
				}
			}
		}
	}
	return total1020
}

func dfs1020(nums [][]int, r int, c int) {
	if !isInNums1020(nums, r, c) {
		return
	}
	if nums[r][c] == 0 {
		//水
		return
	} else if nums[r][c] == 2 {
		//陆地 已经判断过的
		if isEdge1020(nums, r, c) {
			//每当是陆地的时候，判断是否是边界，一旦是边界，false
			sts1020 = false
		}
		return
	} else if nums[r][c] == 1 {
		if isEdge1020(nums, r, c) {
			sts1020 = false
		}
		nums[r][c] = 2
		area1020++
	}

	dfs1020(nums, r-1, c)
	dfs1020(nums, r, c-1)
	dfs1020(nums, r+1, c)
	dfs1020(nums, r, c+1)
}

func isEdge1020(nums [][]int, r, c int) bool {
	if r == 0 || c == 0 || r == len(nums)-1 || c == len(nums[0])-1 {
		return true
	}
	return false
}

func isInNums1020(nums [][]int, r, c int) bool {
	if r >= 0 && r < len(nums) && c >= 0 && c < len(nums[0]) {
		return true
	}
	return false
}

func TestAl1020(t *testing.T) {
	fmt.Printf("%v", numEnclaves([][]int{
		{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
		{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 0}}))
}
