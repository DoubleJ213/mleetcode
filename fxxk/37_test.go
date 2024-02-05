package fxxk

import (
	"testing"
)

/*
37. 解数独

编写一个程序，通过填充空格来解决数独问题。
数独的解法需 遵循如下规则：
数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用 '.' 表示。

示例 1：
输入：board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
输出：[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
解释：输入的数独如上图所示，唯一有效的解决方案如下所示：

提示：
board.length == 9
board[i].length == 9
board[i][j] 是一位数字或者 '.'
题目数据 保证 输入数独仅有一个解
*/

//保证输入数独仅有一个解
//整一个一模一样大小的数组，记录数独对应位置的数是否可以修改。原来的值全部为false。
// 正好默认值就是false。当遍历board时如果发现是.对应的位置就改为true
var ans53 [][]byte

func solveSudoku(board [][]byte) {
	if len(board) != 9 {
		return
	}
	if len(board[0]) != 9 {
		return
	}
	ans53 = make([][]byte, 9)
	for x := 0; x < len(ans53); x++ {
		ans53[x] = make([]byte, 9)
	}
	//有必要的话，先赋值一遍,如果可以在下面递归回溯过程中再赋值
	//找到第一个可以修改的数
	//deep, x := getNext37(0, 0, false)
	//找到最后一个可以修改的数
	dfs37(board, 0, 0)
	return
}

func dfs37(board [][]byte, deep int, x int) bool {
	//返回值表示是否已经找到一个正确的解
	if deep == 9 {
		//	深度为9 deep从0开始，每次往下一层都加1。当deep为9的时候，证明前面9行都填充完毕
		// 光x等于9说明那一行是ok的
		copy(ans53, board)
		return true
	}

	if x == 9 {
		return dfs37(board, deep+1, 0)
	}

	if board[deep][x] != aa["."] {
		return dfs37(board, deep, x+1)
	}

	for _, num := range "123456789" {
		//能进到这里，确保是需要填入的数，且在范围之内

		//	横 纵 9宫格 三种检测 都可以就暂时把这个数填进去，然后继续下一个数
		//纵坐标判断，就是横坐标都是x的 不重复
		xs := isXTrue37(board, deep, x, byte(num))
		if !xs {
			continue
		}
		//横坐标判断 当前deep 9个数 数据不重复
		ys := isYTrue37(board, deep, x, byte(num))
		if !ys {
			continue
		}
		// 九宫格判断
		ss := isSubTrue37(board, deep, x, byte(num))
		if !ss {
			continue
		}
		board[deep][x] = byte(num)
		//注意进位的规则,deep不是一直加的
		// x从0开始，8之后一共9个数 。x++后 x到了9，就得换行继续递归了，同时x=0
		if dfs37(board, deep, x+1) {
			return true
		}
		//注意回溯，退位的规则
		//deep 不是一直减得。先x往前回退，当x回退小于0之后，deep就-1 x=8
		//当前已经是前面一个元素了，如果前面一个是给定的数，那就继续修改索引
		board[deep][x] = aa["."]
	}
	return false
}

func isSubTrue37(board [][]byte, deep int, x int, num byte) bool {
	status := true
	//9*9 3*3的格子一共9个，落在哪个，判断下那3*3其余8个数，没有和自己类似的就行
	for i := (deep / 3) * 3; i < (deep/3)*3+3; i++ {
		for j := (x / 3) * 3; j < (x/3)*3+3; j++ {
			if i != x && board[i][j] == num {
				status = false
				break
			}
		}
	}
	return status
}

func isYTrue37(board [][]byte, deep int, x int, num byte) bool {
	status := true
	for i := 0; i < 9; i++ {
		if i != x && board[deep][i] == num {
			status = false
			break
		}
	}
	return status
}

func isXTrue37(board [][]byte, deep int, x int, num byte) bool {
	status := true
	for i := 0; i < len(board); i++ {
		if i != deep && board[i][x] == num {
			status = false
			break
		}
	}
	return status
}

var aa = map[string]byte{"0": 48, "1": 49, "2": 50, "3": 51, "4": 52,
	"5": 53, "6": 54, "7": 55, "8": 56, "9": 57, ".": 46}

func TestAl37(t *testing.T) {
	/*nums := "."
	for _, num := range nums {
		fmt.Println(num)
	}*/
	solveSudoku([][]byte{{aa["5"], aa["3"], aa["."], aa["."], aa["7"], aa["."], aa["."], aa["."], aa["."]},
		{aa["6"], aa["."], aa["."], aa["1"], aa["9"], aa["5"], aa["."], aa["."], aa["."]},
		{aa["."], aa["9"], aa["8"], aa["."], aa["."], aa["."], aa["."], aa["6"], aa["."]},
		{aa["8"], aa["."], aa["."], aa["."], aa["6"], aa["."], aa["."], aa["."], aa["3"]},
		{aa["4"], aa["."], aa["."], aa["8"], aa["."], aa["3"], aa["."], aa["."], aa["1"]},
		{aa["7"], aa["."], aa["."], aa["."], aa["2"], aa["."], aa["."], aa["."], aa["6"]},
		{aa["."], aa["6"], aa["."], aa["."], aa["."], aa["."], aa["2"], aa["8"], aa["."]},
		{aa["."], aa["."], aa["."], aa["4"], aa["1"], aa["9"], aa["."], aa["."], aa["5"]},
		{aa["."], aa["."], aa["."], aa["."], aa["8"], aa["."], aa["."], aa["7"], aa["9"]}})
}
