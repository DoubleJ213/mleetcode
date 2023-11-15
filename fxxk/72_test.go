package fxxk

import (
	"fmt"
	"testing"
)

/*
72. 编辑距离
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符

示例 1：
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')

示例 2：
输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

提示：
0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成
*/

/*
if s1[i] == s2[j]:
    啥都别做（skip）
    i, j 同时向前移动
else:
    三选一：
        插入（insert）
        删除（delete）
        替换（replace）
*/

var dp72 [][]int

func minDistance(word1 string, word2 string) int {
	lenX := len(word1)
	lenY := len(word2)
	dp72 = make([][]int, lenX+1)
	for k := 0; k < lenX+1; k++ {
		dp72[k] = make([]int, lenY+1)
	}
	for i := 1; i < lenX+1; i++ {
		dp72[i][0] = i
	}

	for j := 1; j < lenY+1; j++ {
		dp72[0][j] = j
	}
	for i := 1; i < lenX+1; i++ {
		for j := 1; j < lenY+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp72[i][j] = dp72[i-1][j-1]
			} else {
				//不相同的情况上有三种操作的可能
				//1）删除左边
				//2）删除右边
				//3）替换/
				//插入跟删除本质是同一个种操作
				dp72[i][j] = 1 + getMin72(dp72[i-1][j], dp72[i][j-1], dp72[i-1][j-1])
			}
		}
	}
	return dp72[lenX][lenY]
}

func getMin72(x, y, z int) int {
	var min int
	if x >= y {
		min = y
	} else {
		min = x
	}
	if min >= z {
		return z
	} else {
		return min
	}
}

func TestAl72(t *testing.T) {
	//fmt.Printf("%v\n", minDistance("horse", "ros"))
	fmt.Printf("%v\n", minDistance("", "a"))
}
