package fxxk

import (
	"fmt"
	"testing"
)

/*
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：
输入：grid = [
  [49,49,49,49,48],
  [49,49,48,49,48],
  [49,49,48,48,48],
  [48,48,48,48,48]
]
输出：1
示例 2：
输入：grid = [
  [49,49,48,48,48],
  [49,49,48,48,48],
  [48,48,49,48,48],
  [48,48,48,49,49]
]
输出：3

提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'
*/

//var dp200 [][]int

/*
初始一个二维数组，其值存对应grid[i][j]位置的岛屿的数量总和
怎么判断一个数是否是一个岛屿，同时是岛屿的一部分？
每个数和上下左右四个数相关。
如果本数是1 是陆地，那岛屿数量+1吗？可以这样认为那之后的陆地就得和前面这个陆地有依赖关系，就不能继续加1了
比如每个数都和前一个和上一个去判断。
如果自己是0，dp 前面和上面是0
让dp表示，遍历到grid 当前值已经有多少个岛屿。当等0的时候直接加上前面的值
*/
var dp200 int

func numIslands(grid [][]byte) int {
	dp200 = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			//row := 0    //行
			//column := 0 //列
			if grid[i][j] == byte('1') {
				dfs200(grid, i, j)
				dp200++
			}
		}
	}
	return dp200
}

func dfs200(grid [][]byte, row int, column int) {
	if !inNums(grid, row, column) {
		return
	}

	if grid[row][column] != byte('1') {
		return
	}

	if grid[row][column] == byte('2') {
		return
	}

	grid[row][column] = byte('2')

	//遍历当前岛屿上下左右四个岛屿
	// 前面各个if 就是处理四个岛屿的逻辑。如果对应的格子为2，表示已经遍历过，就是他和之前某个格子是连在一起的，不能再单独算一个岛屿
	//如果不等于1，那直接return，因为只有1的情况才需要继续判断他的相邻的节点
	dfs200(grid, row-1, column)
	dfs200(grid, row, column-1)
	dfs200(grid, row+1, column)
	dfs200(grid, row, column+1)

}

func inNums(grid [][]byte, r, c int) bool {
	if r >= 0 && c >= 0 && r < len(grid) && c < len(grid[0]) {
		return true
	}
	return false
}

func TestAl200(t *testing.T) {
	fmt.Printf("%v",
		numIslands([][]byte{
			{49, 49, 49, 49, 48},
			{49, 49, 48, 49, 48},
			{49, 49, 48, 48, 48},
			{48, 48, 48, 48, 48}}))
}
