package fxxk

import (
	"fmt"
	"testing"
)

/*
1254. 统计封闭岛屿的数目
二维矩阵 grid 由 0 （土地）和 1 （水）组成。岛是由最大的4个方向连通的 0 组成的群，封闭岛是一个 完全 由1包围（左、上、右、下）的岛。
请返回 封闭岛屿 的数目。

示例 1：
输入：grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
输出：2
解释：
灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。
示例 2：
输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
输出：1
示例 3：
输入：grid = [[1,1,1,1,1,1,1],
             [1,0,0,0,0,0,1],
             [1,0,1,1,1,0,1],
             [1,0,1,0,1,0,1],
             [1,0,1,1,1,0,1],
             [1,0,0,0,0,0,1],
             [1,1,1,1,1,1,1]]
输出：2

提示：
1 <= grid.length, grid[0].length <= 100
0 <= grid[i][j] <=1
*/

/*
这题1是水，0是陆地
靠边的陆地不算封闭的陆地，就是一旦上下左右不在数组内，那这个岛屿就不计算到结果内
跟1020.飞地的数量就差几行字……改两下就过了
*/
var st1254 bool

func closedIsland(grid [][]int) int {
	ans1254 := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			//0 是陆地
			if grid[i][j] == 0 {
				st1254 = true
				//能进循环的条件限制一下。是陆地，且与其连接的陆地不能为边界。
				//与这个岛屿连接的岛屿不能为边界。如果与其相连的是水那就不继续递归
				dfs1254(grid, i, j)
				if st1254 {
					ans1254++
				}
			}
		}
	}
	return ans1254
}

func dfs1254(nums [][]int, r int, c int) {
	if !inNums1254(nums, r, c) {
		return
	}

	if nums[r][c] == 1 {
		// 1表示水
		return
	} else if nums[r][c] == 2 {
		// 2表示遍历过的陆地
		if isEdge(nums, r, c) {
			st1254 = false
		}
		return
	} else if nums[r][c] == 0 {
		//0表示未遍历过的陆地
		if isEdge(nums, r, c) {
			st1254 = false
		}
		nums[r][c] = 2
	}

	dfs1254(nums, r-1, c)
	dfs1254(nums, r, c-1)
	dfs1254(nums, r+1, c)
	dfs1254(nums, r, c+1)
}

func isEdge(grid [][]int, r, c int) bool {
	if r == 0 || c == 0 || r == len(grid)-1 || c == len(grid[0])-1 {
		return true
	}
	return false
}

func inNums1254(grid [][]int, r, c int) bool {
	if r >= 0 && c >= 0 && r < len(grid) && c < len(grid[0]) {
		return true
	}
	return false
}

func TestAl1254(t *testing.T) {
	fmt.Printf("%v", closedIsland([][]int{
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
