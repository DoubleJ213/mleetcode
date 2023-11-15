package fxxk

import (
	"fmt"
	"testing"
)

/*
22. 括号生成

数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：
输入：n = 1
输出：["()"]

提示：
1 <= n <= 8
*/

var res22 []string
var path22 []byte
var status22 map[int]int

/*
( 40
) 41
40的数量要大于41的数量
搞一个map,key 40 和41 然后比较 40 和41的value值
无重可复
*/
func generateParenthesis(n int) []string {
	res22 = make([]string, 0)
	if n == 0 {
		return res22
	}
	status22 = make(map[int]int)
	status22[40] = 0
	status22[41] = 0
	path22 = make([]byte, 0)
	input22 := []byte{40, 41}
	traverse22(n, input22, path22)
	return res22
}

func traverse22(n int, input, path []byte) {
	if 2*n == len(path) {
		tmp := make([]byte, len(path))
		copy(tmp, path)
		res22 = append(res22, string(tmp))
		return
	}
	for _, in := range input {
		status22[int(in)] = status22[int(in)] + 1
		if isOk22(n) {
			path = append(path, in)
			traverse22(n, input, path)
			//左扩号大于等于右括号
			path = path[:len(path)-1]
		}
		status22[int(in)] = status22[int(in)] - 1
	}
}

func isOk22(n int) bool {
	if status22[40] > n || status22[41] > n || status22[40] < status22[41] {
		return false
	}
	return true
}

func TestAl22(t *testing.T) {
	fmt.Printf("%v", generateParenthesis(3))
}
