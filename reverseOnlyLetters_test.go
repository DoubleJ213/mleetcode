package mleetcode

import (
	"fmt"
	"testing"
)

//917. 仅仅反转字母
func reverseOnlyLetters(S string) string {
	// TODO: 双指针
	i := 0
	j := len(S) - 1
	res := []byte(S)
	for i < j {
		for i < len(S)-1 && !isLetter(res[i]) && i < j {
			//i<len(S)-1 此处得加上防止越界的判断
			//i < j i遍历过多次导致i不小于j，是否可以提前结束本次循环
			i++
		}
		for i < len(S)-1 && !isLetter(res[j]) && i < j {
			j--
		}
		if i < j {
			res[i], res[j] = res[j], res[i]
			i++
			j--
		}
	}
	return string(res)
}

func isLetter(b byte) bool {
	if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') {
		return true
	}
	return false
}

func TestReverseOnlyLetters(t *testing.T) {
	var str = "ad-s-dragon-f"
	fmt.Printf("%v", reverseOnlyLetters(str))
}
