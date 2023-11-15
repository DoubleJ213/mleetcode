package fxxk

import (
	"fmt"
	"testing"
)

/*
583. 两个字符串的删除操作
给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
每步 可以删除任意一个字符串中的一个字符。

示例 1：
输入: word1 = "sea", word2 = "eat"
输出: 2
解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"

示例  2:
输入：word1 = "leetcode", word2 = "etco"
输出：4

提示：
1 <= word1.length, word2.length <= 500
word1 和 word2 只包含小写英文字母
*/

/*
删除成最长的公共子序列（1143题）就可以了。
先看最长子序列长度，然后再看一共要删多少次
*/

var dp583 [][]int

func minDistance583(word1 string, word2 string) int {
	lenX := len(word1)
	lenY := len(word2)

	dp583 = make([][]int, lenX+1)
	for k := 0; k <= lenX; k++ {
		dp583[k] = make([]int, lenY+1)
	}

	for i := 1; i <= lenX; i++ {
		for j := 1; j <= lenY; j++ {
			if word1[i-1] == word2[j-1] {
				dp583[i][j] = 1 + dp583[i-1][j-1]
			} else {
				dp583[i][j] = getMax(dp583[i][j-1], dp583[i-1][j])
			}
		}
	}

	longest := dp583[lenX][lenY]
	return lenX + lenY - 2*longest
}

func TestAl583(t *testing.T) {
	//fmt.Printf("%v", minDistance583("sea", "eat"))
	fmt.Printf("%v", minDistance583("leetcode", "etco"))
}
