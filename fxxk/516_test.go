package fxxk

import (
	"fmt"
	"testing"
)

/*
516. 最长回文子序列
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

示例 1：
输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。
示例 2：
输入：s = "cbbd"
输出：2
解释：一个可能的最长回文子序列为 "bb" 。

提示：
1 <= s.length <= 1000
s 仅由小写英文字母组成
*/

var dp516 []int

func longestPalindromeSubseq(s string) int {
	dp516 = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		for j := 0; j < i; j++ {
			var subProblem int
			if isOk131(s[:i]) {
				subProblem = dp516[i] + 1
			} else {
				subProblem = dp516[i]
			}
			dp516[i] = getMax(dp516[i], subProblem)
		}
	}
	return dp516[len(s)-1]
}

func isOk516() {

}

func TestAl516(t *testing.T) {
	fmt.Printf("4 -- %v\n", longestPalindromeSubseq("bbbab"))
	//fmt.Printf("3 -- %v\n", longestPalindromeSubseq("bbba"))
	//fmt.Printf("2 -- %v\n", longestPalindromeSubseq("cbbd"))
}
