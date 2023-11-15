package fxxk

import (
	"fmt"
	"testing"
)

/*
1143. 最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
一个字符串的 子序列 是指这样一个新的字符串：
它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

示例 1：
输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。
示例 2：
输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc" ，它的长度为 3 。
示例 3：
输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0 。

提示：
1 <= text1.length, text2.length <= 1000
text1 和 text2 仅由小写英文字符组成。
*/

/*
直接暴力算法
遍历text1的所有子串，然后遍历text2的所有子串
嵌套循环找最长，或者整个map判断key是否存在找最长
估计会超时， 后来仔细一看 子序列不一定连续 没有暴力法

dp问题就看fib问题，先暴力算法，然后添加备忘录的暴力算法
再之后自底向上求dp[]数组所有的值
重叠子问题、最优子结构、状态转移方程就是动态规划三要素
明确 base case -> 明确「状态」-> 明确「选择」 -> 定义 dp 数组/函数的含义。
base case 零钱0的时候最小需要0个硬币
状态是 零钱数
选择是 每次增加了某个面值的硬币
明确 dp 函数/数组的定义 子问题的处理办法
得理解dp数组的含义

https://blog.csdn.net/weixin_43771947/article/details/123242337
*/
var dpRes1143 [][]int

func longestCommonSubsequence(text1 string, text2 string) int {
	//dp[][]
	//text1 = "abcde", text2 = "ace"
	lenX := len(text1)
	lenY := len(text2)
	dpRes1143 = make([][]int, lenX+1)
	for k := 0; k < lenX+1; k++ {
		dpRes1143[k] = make([]int, lenY+1)
	}
	for i := 1; i < lenX+1; i++ {
		for j := 1; j < lenY+1; j++ {
			if text1[i-1] == text2[j-1] {
				//最后一位相同则1+两个长度都-1的dp值
				dpRes1143[i][j] = 1 + dpRes1143[i-1][j-1]
			} else {
				dpRes1143[i][j] = getMax(dpRes1143[i-1][j], dpRes1143[i][j-1])
			}
		}
	}
	return dpRes1143[lenX][lenY]
}

func TestAl1143(t *testing.T) {
	fmt.Printf("%v", longestCommonSubsequence("abcde", "ace"))
}
