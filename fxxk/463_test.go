package fxxk

import (
	"fmt"
	"testing"
)

/*

463. 岛屿的周长
给定一个 row x col 的二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地， grid[i][j] = 0 表示水域。
网格中的格子 水平和垂直 方向相连（对角线方向不相连）。
整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。
格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。

示例 1：
输入：grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
输出：16
解释：它的周长是上面图片中的 16 个黄色的边
示例 2：
输入：grid = [[1]]
输出：4
示例 3：
输入：grid = [[1,0]]
输出：4

提示：
row == grid.length
col == grid[i].length
1 <= row, col <= 100
grid[i][j] 为 0 或 1
*/

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	//递归每个岛屿，然后计算其有效的周长。岛屿周长本来是4的，当他边上每有一个岛屿的时候，周长就减去1 和他接壤的也减去1，相当于减去2
	//所以，总周长 = 4 * 土地个数 - 2 * 接壤边的条数。
	//遍历矩阵，遍历到土地，就 land++ 第200题,再求接壤的边数量也是一个方法
	//将这个“相邻关系”对应到 DFS 遍历中，就是：每当在 DFS 遍历中，
	// 从一个岛屿方格走向一个非岛屿方格，就将周长加 1。
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				perimeter += dfs463(grid, i, j)
			}
		}
	}

	return perimeter
}

func dfs463(nums [][]int, r int, c int) int {
	//能进递归的，说明之前那个一定是1
	if !inNums463(nums, r, c) {
		//不在数组中都是0
		return 1
	}
	if nums[r][c] == 2 {
		//如果是2，证明遍历过，不用继续计算周长
		return 0
	} else if nums[r][c] == 0 {
		return 1
	}

	//不是2 不是0 证明是1，1证明是一个新的没遍历过的陆地岛屿
	nums[r][c] = 2

	return dfs463(nums, r-1, c) + dfs463(nums, r, c-1) + dfs463(nums, r+1, c) + dfs463(nums, r, c+1)
}

func inNums463(grid [][]int, r, c int) bool {
	//判断是不是在数组内，如果不在数组内，都是0
	if r >= 0 && c >= 0 && r < len(grid) && c < len(grid[0]) {
		return true
	}
	return false
}

func TestAl463(t *testing.T) {
	fmt.Printf("%v", islandPerimeter([][]int{
		{1, 1, 1, 1, 0, 1},
		{1, 1, 0, 1, 0, 1},
		{1, 1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1}}))
}
