package mleetcode

import (
	"fmt"
	"testing"
)

//
//1614. 括号的最大嵌套深度
//本题由于题目描述看不懂，直接看示例吧
//给你一个 有效括号字符串 s，返回该字符串的 s 嵌套深度 。
//示例 1：
//输入：s = "(1+(2*3)+((8)/4))+1"
//输出：3
//解释：数字 8 在嵌套的 3 层括号中。
//示例 2：
//输入：s = "(1)+((2))+(((3)))"
//输出：3
//示例 3：
//输入：s = "1+(2*3)/(2-1)"
//输出：1
//示例 4：
//输入：s = "1"
//输出：0
//
//提示：
//1 <= s.length <= 100
//s 由数字 0-9 和字符 '+'、'-'、'*'、'/'、'('、')' 组成
//题目数据保证括号表达式 s 是 有效的括号表达式

func maxDepth1(s string) int {
	var depTmp int
	var max int
	for _, value := range s {
		if string(value) == "(" {
			depTmp++
			if max < depTmp {
				max = depTmp
			}
		} else if string(value) == ")" {
			depTmp--
		}

	}
	return max
}

func TestAl1614(t *testing.T) {
	fmt.Print(maxDepth1("()(())"))
}
