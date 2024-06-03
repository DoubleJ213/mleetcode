package fxxk

import (
	"fmt"
	"testing"
)

/*
验证已经填入的数字是否有效即可。
数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
*/
/*
	验证已经填入的数字是否有效即可。不是检查棋盘是否能填满
	不等于.的时候对这个数进行判断如果在横 竖 3*3 范围，
	存在另外一个和当前数一样的，就是不有效，其余均有效
	每个数都横竖3*3判断复杂度比较高。
	9*9棋盘，直接判断9个行，9个列 9个3*3都ok好像能省不少时间
*/
// leetcode submit region begin(Prohibit modification and deletion)
func isValidSudoku(board [][]byte) bool {
	//判断9个行 row  列 column 3*3 子board
	for i := 0; i < 9; i++ {
		row := make(map[byte]int)
		column := make(map[byte]int)
		subBoard := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if _, ok := row[board[i][j]]; ok {
					//存在了
					return false
				} else {
					//之前不存在，加入，等之后判断
					row[board[i][j]] = 1
				}
			}
			if board[j][i] != '.' {
				if _, ok := column[board[j][i]]; ok {
					//存在了
					return false
				} else {
					//之前不存在，加入，等之后判断
					column[board[j][i]] = 1
				}
			}
			//subBoard 9宫格 用 0-9 的ij 一起遍历? 不怎么好写
			//fmt.Println(string(board[i/3*3+j/3][i%3*3+j%3]))
			if board[i/3*3+j/3][i%3*3+j%3] != '.' {
				if _, ok := subBoard[board[i/3*3+j/3][i%3*3+j%3]]; ok {
					//存在了
					return false
				} else {
					//之前不存在，加入，等之后判断
					subBoard[board[i/3*3+j/3][i%3*3+j%3]] = 1
				}
			}
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidSudoku(t *testing.T) {
	//这题首先测试用例怎么构造,因为有字符串.的存在所以类型byte.
	//所以不能无脑双引号了，单引号表示byte
	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'.', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}
