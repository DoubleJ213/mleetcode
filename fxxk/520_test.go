package fxxk

import (
	"fmt"
	"testing"
)

/*
全部字母都是大写，比如 "USA" 。
单词中所有字母都不是大写，比如 "leetcode" 。
如果单词不只含有一个字母，只有首字母大写， 比如 "Google" 。
先看下小写大写的范围
*/
// TODO 统计大写字母个数 如果0个大写字母，true 如果全部大写字母 true 如果1个大写字母，且是index=0为大写字母 true，其余false
//leetcode submit region begin(Prohibit modification and deletion)
func detectCapitalUse(word string) bool {
	l := 0
	for _, c := range word {
		// unicode.IsUpper(c)
		if getUpOrDown(byte(c)) {
			l++
		}
	}
	if l == 0 || l == len(word) || (l == 1 && getUpOrDown(word[0])) {
		return true
	}
	return false
	//return l == 0 || l == len(word) || (l == 1 && getUpOrDown(word[0]))
}

// true 大写 false 小写
func getUpOrDown(b byte) bool {
	if b >= 'A' && b <= 'Z' { //	大写字母
		return true
	} else if b >= 'a' && b <= 'z' { //	小写字母
		return false
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDetectCapital(t *testing.T) {
	//fmt.Println(detectCapitalUse("a")) //97
	//fmt.Println(detectCapitalUse("z")) //122
	//fmt.Println(detectCapitalUse("A")) //65
	//fmt.Println(detectCapitalUse("Z")) //90
	fmt.Println(detectCapitalUse("ggg")) //90

}

func detectCapitalUse2(word string) bool {
	//w := []byte(word)
	if len(word) <= 0 {
		return false
	}
	if len(word) == 1 {
		return true
	}
	var first bool            // true 大写 false 小写
	if getUpOrDown(word[0]) { //	首字母是大写字母
		//已经限制 word 由小写和大写英文字母组成
		first = true
	}
	var second bool
	if getUpOrDown(word[1]) {
		second = true
	}
	if !first && second {
		return false
	}
	for i := 2; i < len(word); i++ {
		if first { //	首字母是大写字母
			//要么全是小写，要么全是大写，否则 false
			if second != getUpOrDown(word[i]) {
				return false
			}
		} else { // 首字母 小写
			if word[i] >= 'A' && word[i] <= 'Z' {
				return false
			}
		}
	}
	return true
}
