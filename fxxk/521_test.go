package fxxk

import (
	"testing"
)

/*
之前写过522 数组中有很多字符串，找最长的特殊序列
现在就2个字符串，然后求最长序列*/
// leetcode submit region begin(Prohibit modification and deletion)
func findLUSlength521(a string, b string) int {
	if len(a) > len(b) {
		return len(a)
	} else if len(a) < len(b) {
		return len(b)
	} else if a != b {
		//一样长 但是不相同
		return len(a)
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestUncommonSubsequenceI(t *testing.T) {

}
