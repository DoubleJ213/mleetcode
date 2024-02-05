package fxxk

import (
	"fmt"
	"testing"
)

//https://leetcode.cn/problems/ZL6zAn/description/
//LCR 105. 岛屿的最大面积
//本题与主站 695 题相同： https://leetcode-cn.com/problems/max-area-of-island/

/*
给定一个由 0 和 1 组成的非空二维数组 grid ，用来表示海洋岛屿地图。
一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。
你可以假设 grid 的四个边缘都被 0（代表水）包围着。
找到给定的二维数组中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

示例 1:
输入: grid = [
[0,0,1,0,0,0,0,1,0,0,0,0,0],
[0,0,0,0,0,0,0,1,1,1,0,0,0],
[0,1,1,0,1,0,0,0,0,0,0,0,0],
[0,1,0,0,1,1,0,0,1,0,1,0,0],
[0,1,0,0,1,1,0,0,1,1,1,0,0],
[0,0,0,0,0,0,0,0,0,0,1,0,0],
[0,0,0,0,0,0,0,1,1,1,0,0,0],
[0,0,0,0,0,0,0,1,1,0,0,0,0]]
输出: 6
解释: 对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。
示例 2:
输入: grid = [[0,0,0,0,0,0,0,0]]
输出: 0

提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 50
grid[i][j] is either 0 or 1
*/
var visitedL105 [][]int

func maxAreaOfIsland1(grid [][]int) int {
	visitedL105 = make([][]int, len(grid))
	for x := 0; x < len(grid); x++ {
		visitedL105[x] = make([]int, len(grid[x]))
	}
	areaL105 := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			area := 0
			if visitedL105[i][j] == 0 && grid[i][j] == 1 {
				area = dfsLrc105(grid, i, j)
			}
			areaL105 = getMax(areaL105, area)
		}
	}
	return areaL105
}

func dfsLrc105(nums [][]int, x int, y int) int {
	if !isInNums105(nums, x, y) {
		return 0
	}

	if visitedL105[x][y] == 1 {
		return 0
	}
	//能走到这里，那就是满足什么条件的，没有被访问过，在数组范围内
	visitedL105[x][y] = 1
	current := nums[x][y]
	if current == 1 {
		return 1 + dfsLrc105(nums, x-1, y) +
			dfsLrc105(nums, x+1, y) +
			dfsLrc105(nums, x, y-1) +
			dfsLrc105(nums, x, y+1)
	}
	return 0

}

func isInNums105(grid [][]int, x int, y int) bool {
	if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) {
		return true
	}
	return false
}

func TestLrc105(t *testing.T) {
	fmt.Printf("%v", maxAreaOfIsland1([][]int{
		{1, 1, 0, 1, 0, 1},
		{0, 0, 0, 1, 0, 1},
		{1, 1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1}}))
}
