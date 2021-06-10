package mleetcode

import (
	"fmt"
	"testing"
)

func longestPalindrome(s string) string {
	length := len(s)
	if len(s) < 2 {
		return s
	}

	maxLen := 1
	begin := 0
	for i := 0; i < length-1; i++ {
		//左边界
		for j := i + 1; j < length; j++ {
			//右边界
			if j-i+1 > maxLen && isPalinDrome(s, i, j) {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	//字符串截取
	return s[begin : maxLen+begin]
}

func isPalinDrome(s string, i, j int) bool {
	//判断是否是回文字符串
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func TestA(t *testing.T) {
	test1 := "abcabacd"
	fmt.Println(longestPalindrome(test1))

	test2 := "asca"
	fmt.Println(longestPalindrome(test2))
}
