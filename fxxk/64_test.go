package fxxk

import (
	"fmt"
	"math"
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
/*
找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
每次只能向下或者向右移动一步。
首先能想到这是个dp题目，然后这个总和的最小值。其实就是选路比较。几种路线里面选最小值
假设有 最小值 dp[n][m], dp[n][m] = getMin(dp[n][m-1], dp[n-1][m])+ grid[n][m]
*/
var dp64 [][]int

/*
m == grid.length
n == grid[i].length
1 <= m, n <= 200
*/
func minPathSum(grid [][]int) int {
	x := len(grid)
	y := len(grid[0])
	dp64 = make([][]int, 0)
	for i := 0; i < x+1; i++ {
		dp64 = append(dp64, make([]int, y+1))
	}

	for m := 0; m < x+1; m++ {
		for n := 0; n < y+1; n++ {
			if m == 0 || n == 0 {
				dp64[m][n] = math.MaxInt
			} else if m == 1 && n == 1 {
				dp64[m][n] = grid[m-1][n-1]
			} else {
				dp64[m][n] = getMin(dp64[m-1][n], dp64[m][n-1]) + grid[m-1][n-1]
			}
			//traverse64(grid, m-1, n-1)
		}
	}
	return dp64[x][y]
}

//
//func traverse64(grid [][]int, m, n int) {
//	// 限制不出grid范围
//	if n < 0 || m < 0 || len(grid)-1 < m || len(grid[0])-1 < n {
//		return
//	}
//
//	fmt.Println(grid[m][n])
//	dp64[m][n] = getMin(dp64[m-1][n], dp64[m][n-1]) + grid[m][n]
//	traverse64(grid, m+1, n)
//	traverse64(grid, m, n+1)
//}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumPathSum(t *testing.T) {
	fmt.Println(minPathSum([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	//fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))

}
