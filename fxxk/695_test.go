package fxxk

import (
	"fmt"
	"testing"
)

/*
695.最大的岛屿面积

给你一个大小为 m x n 的二进制矩阵 grid 。
岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。
你可以假设 grid 的四个边缘都被 0（代表水）包围着。
岛屿的面积是岛上值为 1 的单元格的数目。
计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

输入：grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
输出：6
解释：答案不应该是 11 ，因为岛屿只能包含水平或垂直这四个方向上的 1 。
示例 2：

输入：grid = [[0,0,0,0,0,0,0,0]]
输出：0

*/

/*
最大的岛屿面积，不是所有的岛屿面积
*/
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			area := 0
			if grid[i][j] == 1 {
				area = dfs695(grid, i, j)
			}
			maxArea = getMax(maxArea, area)
		}
	}
	return maxArea
}

func dfs695(nums [][]int, r int, c int) int {
	if !inNums695(nums, r, c) {
		return 0
	}

	//if nums[r][c] == 2 {
	//	return 0
	//} else if nums[r][c] == 0 {
	//	return 0
	//}

	if nums[r][c] != 1 {
		return 0
	}
	//相邻的，且第一次遍历的1
	nums[r][c] = 2
	return 1 + dfs695(nums, r-1, c) + dfs695(nums, r, c-1) + dfs695(nums, r+1, c) + dfs695(nums, r, c+1)
}

func inNums695(grid [][]int, r, c int) bool {
	if r >= 0 && c >= 0 && r < len(grid) && c < len(grid[0]) {
		return true
	}
	return false
}

func TestAl695(t *testing.T) {
	fmt.Printf("%v", maxAreaOfIsland([][]int{
		{1, 1, 1, 1, 0, 1},
		{1, 1, 0, 1, 0, 1},
		{1, 1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1}}))
}
