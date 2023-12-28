package fxxk

import (
	"fmt"
	"testing"
)

/*
52. N 皇后 II
n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。

示例 1：
输入：n = 4
输出：2
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：
输入：n = 1
输出：1

提示：
1 <= n <= 9
*/

var ans52 int

//哈哈，这不是51改一改呀，有手就行啊，好了改它。
func totalNQueens(n int) int {
	ans52 = 0
	if n < 1 || n > 9 {
		return ans52
	}
	//indexPath 记录遍历到当前，所有出现过Q的位置, 等于1表示出现过 0表示未出现过q
	qPath := make([][]int, n)
	for i := 0; i < n; i++ {
		qPath[i] = make([]int, n)
	}

	dfs52(n, 0, qPath)
	return ans52
}

func dfs52(n, deep int, qPath [][]int) {
	//判断横坐标，纵坐标是否满足要求
	//判断两个斜线是否满足要求
	if n == deep {
		//	n 为 4 那就是4*4的棋盘
		//	deep 为 0123， index 也是0123
		// 合并结果集
		//将qPath 翻译成字符串，然后加到结果集中
		ans52++
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
		dfs52(n, deep, qPath)
		deep--
		qPath[deep][i] = 0

		//纵坐标每次只改一个qPath的值就行，先不判断
		//isYTrue(0, i, qPath)
	}
}

func TestAl52(t *testing.T) {
	fmt.Println(totalNQueens(4))
}
