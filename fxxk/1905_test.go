package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
var used1905 [][]bool
var isLd1905 bool //证明当前已经是个岛屿了

/*
grid2 的一个岛屿，被 grid1 的一个岛屿 完全 包含，
也就是说 grid2 中该岛屿的每一个格子都被 grid1 中同一个岛屿完全包含，那么我们称 grid2 中的这个岛屿为 子岛屿 。
请你返回 grid2 中 子岛屿 的 数目 。
0 （表示水域）和 1 （表示陆地）任何矩阵以外的区域都视为水域。
题目已经限制入参了，数组1*1以上
*/
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	x := len(grid1)
	y := len(grid1[0])
	used1905 = make([][]bool, 0)
	for i := 0; i < x; i++ {
		used1905 = append(used1905, make([]bool, y))
	}
	var ans1905 int
	for m := 0; m < x; m++ {
		for n := 0; n < y; n++ {
			isLd1905 = false
			if traverse1905(grid1, grid2, m, n) {
				ans1905++
			}
		}
	}
	return ans1905
}

func traverse1905(grid1, grid2 [][]int, m, n int) bool {
	// 是一个岛屿的条件。首先自己是1，然后别人是1。然后上下左右继续遍历如果上下左右
	// grid2 的一个岛屿，被 grid1 的一个岛屿 完全包含 所以一旦grid2 是岛屿但是grid1 不是
	// 所以grid2 当前这个岛屿直接不用遍历了
	if m < 0 || n < 0 || len(grid2) <= m || len(grid2[0]) <= n ||
		grid2[m][n] == 0 || used1905[m][n] {
		//如果出界了，返回，但是不影响之前的结果
		//2 的边界，返回，但是不影响之前的结果
		//之前判断过，也返回，也不影响之前的结果
		return isLd1905
	}

	used1905[m][n] = true
	// grid2[m][n] == 0 已经处理这里只能是1 所以 图2 为岛屿
	isLd1905 = true

	up := traverse1905(grid1, grid2, m-1, n)
	down := traverse1905(grid1, grid2, m+1, n)
	left := traverse1905(grid1, grid2, m, n-1)
	right := traverse1905(grid1, grid2, m, n+1)

	if grid1[m][n] == 0 {
		//2的其余相连的1，全都返回false，全部标志位遍历过就行
		return false
	} else {
		//当前为1岛屿  其他也为1岛屿 才继续递归其余直接true		//上 、下 、左、右
		return up && down && left && right
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCountSubIslands(t *testing.T) {
	grid1 := [][]int{
		{1, 1, 1, 0, 0},
		{0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 1, 1}}
	grid2 := [][]int{
		{1, 1, 1, 0, 0},
		{0, 0, 1, 1, 1},
		{0, 1, 0, 0, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 1, 0}}
	/*
		grid2 := [][]int{
			{0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1},
			{0, 1, 0, 1, 0},
			{0, 1, 0, 1, 0},
			{1, 0, 0, 0, 1}}

		grid1 := [][]int{
			{1, 0, 1, 0, 1},
			{1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1},
			{1, 0, 1, 0, 1}}*/

	fmt.Println(countSubIslands(grid1, grid2))
}
