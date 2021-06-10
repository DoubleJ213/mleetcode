package mleetcode

import (
	"fmt"
	"testing"
)

func longestPalindrome2(s string) string {
	length := len(s)
	if len(s) < 2 {
		return s
	}
	maxLen := 0
	begin := 0
	for i := 0; i < length-1; i++ {
		//i为中心位置
		//奇数
		oddLen := getMaxLen(s, i, i)
		//偶数
		evenLen := getMaxLen(s, i, i+1)
		moreLen := 0
		if oddLen > evenLen {
			moreLen = oddLen
		} else {
			moreLen = evenLen
		}
		if moreLen > maxLen {
			maxLen = moreLen
			begin = i - (maxLen-1)/2
		}
	}

	return s[begin : maxLen+begin]
}

func getMaxLen(s string, i, j int) int {
	index := 0
	maxLen := 0
	isOdd := true
	if i != j {
		isOdd = false
	}
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			break
		}
		i--
		j++
		index++
	}
	if isOdd {
		//奇数 aba
		maxLen = 2*index - 1
	} else {
		//偶数
		maxLen = 2 * index
	}
	return maxLen
}

func TestB(t *testing.T) {
	test1 := "babad"
	fmt.Println(longestPalindrome2(test1))

	test2 := "asca"
	fmt.Println(longestPalindrome2(test2))
}
