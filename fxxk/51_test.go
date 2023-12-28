package fxxk

import (
	"fmt"
	"testing"
)

/*
51. N 皇后

按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

示例 1：
输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：
输入：n = 1
输出：[["Q"]]

提示：
1 <= n <= 9
*/

/*
这是个很经典的问题，不过很久之前看过因为是困难题当时就没写。现在也不记得怎么解了，重新思考试一试
先理解下题目
n = 1
1*1的棋盘

放一个Q就好了
那返回值 [["Q"]]

n=2 2*2的棋盘
第一行 q .  .q 两种情况

遍历第二行，继续放Q看看有几种情况
满足情况的记录一下
q .
. q  都不满足， 所以为nil

n=3 3*3的棋盘
第一行
q..   .q.  ..q
q..
..q
.q. 不满足


.q.
...
不满足


..q
q..
... 不满足


所以n=3 也等于nil

n=4 4*4的棋盘

q... .q.. ..q. ...q 四种

q...
..q.

q...
...q
有点像树的结构，我尝试写一下
*/

/*
纵坐标表示深度很容易搞定，就是同一个纵坐标，只放一个。就是当前层，能放下q就好了，后面不用放了
但是橫坐标就有讲究了。不能和前面所有的横坐标相同，而且和前面一行的 横坐标也有关系
斜线有两个（x,y） （x-1, y-1） (x-1,y+1)
所以得把这个树所有已经存在的横坐标都记录下来 preX 前面所有q存在过的x.那那种一层有2种情况咋弄
所以prex也分叉了，是个树结构
*/

/*
1 <= n <= 9 所以肯定有投机取巧的兄弟
*/
var path51 [][]string

// 不要投机取巧。直接return
func solveNQueens(n int) [][]string {
	path51 = make([][]string, 0)
	if n < 1 || n > 9 {
		return nil
	}
	//indexPath 记录遍历到当前，所有出现过Q的位置, 等于1表示出现过 0表示未出现过q
	qPath := make([][]int, n)
	for i := 0; i < n; i++ {
		qPath[i] = make([]int, n)
	}

	dfs51(n, 0, qPath)
	return path51
}

func isPostTrue(deep int, x int, qPath [][]int) bool {
	status := true
	i := 1
	j := 1
	for deep-i >= 0 && x+j < len(qPath[i]) {
		if qPath[deep-i][x+j] == 1 {
			status = false
			break
		}
		j++
		i++
	}
	return status
}

func isPreTrue(deep int, x int, qPath [][]int) bool {
	status := true
	i := 1
	j := 1
	for deep-i >= 0 && x-j >= 0 {
		//第一个（0,0） （1,1）
		//deep-1 x-1
		if qPath[deep-i][x-j] == 1 {
			status = false
			break
		}
		i++
		j++
	}
	return status
}

func isXTrue(deep int, x int, qPath [][]int) bool {
	status := true
	for i := 0; i < deep; i++ {
		if qPath[i][x] == 1 {
			status = false
			break
		}
	}
	return status
}

func dfs51(n, deep int, qPath [][]int) {
	//判断横坐标，纵坐标是否满足要求
	//判断两个斜线是否满足要求
	if n == deep {
		//	n 为 4 那就是4*4的棋盘
		//	deep 为 0123， index 也是0123
		// 合并结果集
		//将qPath 翻译成字符串，然后加到结果集中
		tmp := getResult(qPath)
		path51 = append(path51, tmp)
		return
	}
	for i := 0; i < n; i++ {
		//isXTrue
		//斜线是否满足,斜线有两种。
		//isPreTrue \
		//isPostTrue /
		x := isXTrue(deep, i, qPath)
		if !x {
			continue
		}
		r := isPreTrue(deep, i, qPath)
		if !r {
			continue
		}
		o := isPostTrue(deep, i, qPath)
		if !o {
			continue
		}
		qPath[deep][i] = 1
		deep++
		dfs51(n, deep, qPath)
		deep--
		qPath[deep][i] = 0

		//纵坐标每次只改一个qPath的值就行，先不判断
		//isYTrue(0, i, qPath)
	}
}

func getResult(qPath [][]int) []string {
	ans := make([]string, 0)
	for i := 0; i < len(qPath); i++ {
		level := ""
		cur := ""
		for j := 0; j < len(qPath[i]); j++ {
			if qPath[i][j] == 1 {
				cur = "Q"
			} else {
				cur = "."
			}
			level += cur
		}
		ans = append(ans, level)
	}
	return ans
}

func TestAl51(t *testing.T) {
	fmt.Println(solveNQueens(4))
}

func isYTrue(deep int, x int, qPath [][]int) {
	for i := 0; i <= deep; i++ {
		if qPath[i][x] == 1 {
			break
		}
	}
	return
}
