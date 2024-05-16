package fxxk

import (
	"fmt"
	"testing"
)

/*
在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
*/

//leetcode submit region begin(Prohibit modification and deletion)

/*
乍一看，这是岛屿题目呀，但是和岛屿题不一样的地方是，岛屿的题目只需要遍历一次，找到面积啊，满足要求的岛屿个数啊就完事
这题需要找到最小分钟时。其实就是找到你使用不方式递归的不同值，然后选最小值
如果每种访问可能都遍历一遍，复杂度太高了。
先找到坏橘子，然后以每个坏橘子为起点，按照分钟传染给上下左右的规律，刷新第二分钟的数组，然后一直刷新，直到没有好橘子。
考虑几个情况，
a.终止条件，就是每个2附近没有1了.旁边可以是0也可以是2
b.当递归终止后。记录下当前的分钟数，同时需要继续判断是否还存在1，如果存在返回-1，没有1返回刚刚记录的数

注意-1的情况
*/

var res994 int
var countOne int
var errorList [][]int
var toDoList [][]int

func orangesRotting(grid [][]int) int {
	res994 = 0
	countOne = 0
	errorList = make([][]int, 0)
	for m := 0; m < len(grid); m++ {
		for n := 0; n < len(grid[0]); n++ {
			if grid[m][n] == 2 {
				current := []int{m, n}
				errorList = append(errorList, current)
			}
			if grid[m][n] == 1 {
				countOne++
			}
		}
	}
	for {
		toDoList = make([][]int, 0)
		for i := 0; i < len(errorList); i++ {
			//以每个坏橘子为起点.
			//每次刷新起点是变化的？这个怎么办
			x, y := errorList[i][0], errorList[i][1]
			setNum(grid, x-1, y)
			setNum(grid, x+1, y)
			setNum(grid, x, y-1)
			setNum(grid, x, y+1)
		}

		if len(toDoList) > 0 {
			res994++
		} else {
			break
		}
		errorList = toDoList
		//每个初始为2的坏橘子都把整个棋盘遍历一遍的方法
	}
	if countOne > 0 {
		return -1
	}
	return res994
}

func setNum(grid [][]int, m, n int) {
	if m < 0 || n < 0 || len(grid) <= m || len(grid[0]) <= n {
		return
	}
	if grid[m][n] == 1 {
		countOne--
		grid[m][n] = 2
		toDoList = append(toDoList, []int{m, n})
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRottingOranges(t *testing.T) {
	grid1 := [][]int{
		{2, 1, 1, 2},
		{1, 1, 0, 0},
		{0, 1, 1, 1},
		{1, 0, 1, 1},
	}
	fmt.Println(orangesRotting(grid1))
}
