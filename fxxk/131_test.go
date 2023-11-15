package fxxk

import (
	"fmt"
	"testing"
)

/*
131. 分割回文串
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
回文串 是正着读和反着读都一样的字符串。

示例 1：
输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
示例 2：
输入：s = "a"
输出：[["a"]]

提示：
1 <= s.length <= 16
s 仅由小写英文字母组成
*/
var res131 [][]string
var path131 []string

func partition131(s string) [][]string {
	res131 = make([][]string, 0)
	path131 = make([]string, 0)
	traverse131(s, path131)
	return res131
}

func traverse131(s string, path []string) {
	if s == "" {
		tmp := make([]string, len(path))
		copy(tmp, path)
		res131 = append(res131, tmp)
		return
	}

	for i := 0; i <= len(s); i++ {
		level := s[:i]
		if level == "" {
			continue
		}
		if isOk131(level) {
			path = append(path, level)
			traverse131(s[i:], path)
			path = path[:len(path)-1]
		}
	}
}

func isOk131(s string) bool {
	left, right := 0, len(s)-1
	mid := (right-left)>>1 + left
	for i := 0; i <= mid; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func TestAl131(t *testing.T) {
	//fmt.Println(isOk131("abaa"))
	fmt.Println(partition131("abaa"))
}
