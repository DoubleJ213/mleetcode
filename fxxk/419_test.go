package fxxk

import (
	"fmt"
	"testing"
)

/*
战舰 只能水平或者垂直放置在 board 上。
换句话说，战舰只能按 1 x k（1 行，k 列）或 k x 1（k 行，1 列）的形状建造，
其中 k 可以是任意大小，没有相邻的战舰
返回在甲板 board 上放置的 战舰 的数量
进阶：你可以实现一次扫描算法，并只使用 O(1) 额外空间，并且不修改 board 的值来解决这个问题吗？
*/
//leetcode submit region begin(Prohibit modification and deletion)
var visited419 [][]bool

func countBattleships(board [][]byte) int {
	var ans int
	visited419 = make([][]bool, len(board))
	for x := 0; x < len(board); x++ {
		visited419[x] = make([]bool, len(board[x]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if !visited419[i][j] {
				//没有访问过
				if board[i][j] == 'X' { //是个战舰
					ans++
					dfs419(board, i, j)
				}
			}
		}
	}
	return ans
}

func dfs419(board [][]byte, i int, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return //越界退出
	}
	if visited419[i][j] {
		return //已经访问过退出
	}
	visited419[i][j] = true
	if board[i][j] == 'X' {
		dfs419(board, i-1, j) //上
		dfs419(board, i+1, j) //下
		dfs419(board, i, j-1) //左
		dfs419(board, i, j+1) //右
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBattleshipsInABoard(t *testing.T) {
	//fmt.Println(countBattleships([][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}))
	fmt.Println(countBattleships([][]byte{{'.'}}))
}
