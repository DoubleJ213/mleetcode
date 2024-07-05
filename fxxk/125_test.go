package fxxk

import (
	"fmt"
	"testing"
)

// 大写字符转换为小写字符、并移除所有非字母数字字符之后
// 字母和数字都属于字母数字字符。
// [97,122]小写字母
// [65 ,90]大写字母
// [48,57] 数字
// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome125(s string) bool {
	a := []byte(s)
	l := 0
	r := len(a) - 1
	for l <= r {
		for l < r && notIn125(a[l]) {
			l++
		}
		for l < r && notIn125(a[r]) {
			r--
		}
		if 'A' <= a[l] && a[l] <= 'Z' {
			a[l] += 32
		}
		if 'A' <= a[r] && a[r] <= 'Z' {
			a[r] += 32
		}
		if a[l] != a[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func notIn125(b byte) bool {
	if 'a' <= b && b <= 'z' {
		return false
	} else if 'A' <= b && b <= 'Z' {
		return false
	} else if '0' <= b && b <= '9' {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidPalindrome(t *testing.T) {
	//fmt.Println(isPalindrome(" A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome125(" "))
}
