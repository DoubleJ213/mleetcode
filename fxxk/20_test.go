package fxxk

import (
	"fmt"
	"testing"
)

//20. 有效的括号
func isValid(s string) bool {
	//TODO：字符串遍历 rune -- int32 byte -- uint8 string 深究
	var stack = make([]rune, 0)
	strMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, c := range s {
		if _, ok := strMap[c]; !ok {
			//确定是左边半个，直接加到数组最后
			stack = append(stack, c)
		} else if len(stack) > 0 && stack[len(stack)-1] == strMap[c] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}

	}
	return len(stack) == 0
}
func TestIsValid(t *testing.T) {
	var str = "()(){}[]"
	status := isValid(str)
	fmt.Printf("%v", status)
}
