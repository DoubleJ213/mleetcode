package fxxk

import (
	"fmt"
	"testing"
)

/*
712. 两个字符串的最小ASCII删除和
给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和 。

示例 1:
输入: s1 = "sea", s2 = "eat"
输出: 231
解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。

在 "eat" 中删除 "t" 并将 116 加入总和。
结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。

示例 2:
输入: s1 = "delete", s2 = "leet"
输出: 403
解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。

提示:
0 <= s1.length, s2.length <= 1000
s1 和 s2 由小写英文字母组成
*/

/*

 */

var dp712 [][]int

/*
不同于583 1143等公共子序列问题
这个dp数组里面可以存这个公共子序列的长度也可以累加对应的值，但是起整体逻辑是一致的
类似链表里面的元素不一定是int类型一样
*/
func minimumDeleteSum(s1 string, s2 string) int {
	lenX := len(s1)
	lenY := len(s2)
	dp712 = make([][]int, lenX+1)
	for k := 0; k < lenX+1; k++ {
		dp712[k] = make([]int, lenY+1)
	}

	for i := 1; i < lenX+1; i++ {
		for j := 1; j < lenY+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp712[i][j] = dp712[i-1][j-1] + int(s1[i-1])
			} else {
				dp712[i][j] = getMax(dp712[i][j-1], dp712[i-1][j])
			}
		}
	}
	var total int
	for _, a := range s1 {
		total = total + int(a)
	}
	for _, b := range s2 {
		total = total + int(b)
	}
	return total - 2*dp712[lenX][lenY]
}

func TestAl712(t *testing.T) {
	fmt.Printf("%v\n", minimumDeleteSum("abcde", "ace"))
	fmt.Printf("%v\n", minimumDeleteSum("sea", "eat"))
	fmt.Printf("%v\n", minimumDeleteSum("delete", "leet"))
}
