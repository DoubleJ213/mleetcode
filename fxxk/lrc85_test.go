package fxxk

import (
	"fmt"
	"testing"
)

/*
LCR 085. 括号生成
正整数 n 代表生成括号的对数，请设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：
输入：n = 1
输出：["()"]

提示：
1 <= n <= 8
*/

/*
有关括号问题，你只要记住以下性质，思路就很容易想出来：
1、一个「合法」括号组合的左括号数量一定等于右括号数量，这个很好理解。
2、对于一个「合法」的括号字符串组合 p，必然对于任何 0 <= i < len(p) 都有：
子串 p[0..i] 中左括号的数量都大于或等于右括号的数量。
*/
var stsLrc85 map[string]int
var ansLrc85 []string

func generateParenthesis(n int) []string {
	ansLrc85 = make([]string, 0)
	stsLrc85 = make(map[string]int)
	stsLrc85["("] = 0
	stsLrc85[")"] = 0
	dfsLrc85(2*n, 0, "")
	return ansLrc85
}

func dfsLrc85(n int, deep int, path string) {
	if n == deep {
		//	遍历到最后一个了
		//copy path
		ansLrc85 = append(ansLrc85, path)
		return
	}
	//可以重复选
	for _, value := range "()" {
		stsLrc85[string(value)] ++
		//( > ) 并且 不能大于n的一半
		if isOkLrc85(value, n) {
			deep++
			path = path + string(value)
			dfsLrc85(n, deep, path)
			path = path[:len(path)-1]
			deep--
		}
		stsLrc85[string(value)] --
	}
}

func isOkLrc85(value int32, n int) bool {
	//stsLrc85[string(value)]
	if stsLrc85[string(value)] <= n/2 && stsLrc85["("] >= stsLrc85[")"] {
		return true
	}
	return false
}

func TestAllrc85(t *testing.T) {
	fmt.Printf("%v", generateParenthesis(3))
}
