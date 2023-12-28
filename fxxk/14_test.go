package fxxk

import (
	"fmt"
	"strings"
	"testing"
)

/*
 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

提示：
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成
*/

var ans string

func longestCommonPrefix(strs []string) string {
	ans = ""
	n := len(strs)
	for x := 0; x < n; x++ {
		pre := strs[0][x]
		for i := 1; i < len(strs); i++ {
			if x >= len(strs[i]) {
				return ans
			}
			cur := strs[i][x]
			if pre == cur {
				continue
			} else {
				return ans
			}
		}
		ans = strings.Join([]string{ans, string(pre)}, "")
	}
	return ans
}

func TestAl14(t *testing.T) {
	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"race", "racecar", "r"}))
}
