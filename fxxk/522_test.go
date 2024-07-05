package fxxk

import (
	"fmt"
	"testing"
)

/*
特殊序列 定义如下：该序列为某字符串 独有的子序列（即不能是其他字符串的子序列）。
s 的 子序列可以通过删去字符串 s 中的某些字符实现
例如，"abc" 是 "aebdc" 的子序列，
因为您可以删除"aebdc"中的下划线字符来得到 "abc" 。
"aebdc"的子序列还包括"aebdc"、 "aeb" 和 "" (空字符串)
*/
/*
只要2个字符串不是一摸一样的，其各自长度都能成为最长的特殊序列的长度
如果1个字符串a 是另一个字符串 ab 的子序列，结果2 也是没问题的
*/
//leetcode submit region begin(Prohibit modification and deletion)

func findLUSlength(strs []string) int {
	ans := -1
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) <= ans { // 不会让 ans 变大, 既在嵌套循环处不用重复计算，又不需要针对ans取最大值 排序一下好像也不错
			continue
		}
		var isSub bool // 是否是其他字符串的子串
		for j := 0; j < len(strs); j++ {
			if i != j {
				isSub = isSubStr522(strs[i], strs[j])
				if isSub {
					break
				}
			}
		}
		if !isSub {
			ans = len(strs[i])
		}
	}
	return ans
}

func isSubStr522(a, b string) bool {
	if len(a) > len(b) {
		return false
	}
	x, y := 0, 0
	for x < len(a) {
		for y < len(b) {
			if x < len(a) && a[x] == b[y] {
				x++
			}
			y++
		}
		if y >= len(b) {
			break
		}
	}
	if x == len(a) {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestUncommonSubsequenceIi(t *testing.T) {
	//fmt.Println(findLUSlength([]string{"aaaa", "aaa", "aaaaa", "a"}))
	fmt.Println(findLUSlength([]string{"aaa", "aaa", "aa", "aba", "cdc", "eae", "abcbd"}))
	//fmt.Println(isSubStr("bbababbbd", "aa"))
	//fmt.Println(isSubStr("abc", "ahbgdc"))
	//fmt.Println(isSubStr("b", "abc"))
}
