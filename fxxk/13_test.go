package fxxk

import (
	"fmt"
	"testing"
)

/*
1 <= s.length <= 15
s 仅含字符 ('I', 'V', 'X', 'L', 'C', 'D', 'M')
题目数据保证 s 是一个有效的罗马数字，且表示整数在范围 [1, 3999] 内
IL 和 IM 这样的例子并不符合题目要求，
49 应该写作 XLIX，999 应该写作 CMXCIX 。

字符  	数值
I 	 	1
V 	 	5
X	  	10
L 		50
C 	 	100
D 	 	500
M 	 	1000
*/
// leetcode submit region begin(Prohibit modification and deletion)

var g13 = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "IX": 9,
	"X": 10, "XX": 20, "XXX": 30, "XL": 40, "L": 50, "XC": 90,
	"C": 100, "CC": 200, "CCC": 300, "CD": 400, "D": 500, "CM": 900,
	"M": 1000,
}

// g13 存了所有可能性的 先看三位的字符串能不能找到，然后再看两位的能不能找到，最后再找一位，结果相加就行
func romanToInt(s string) int {
	ans, size, i := 0, 3, 0
	for i < len(s) {
		cur := ""
		if size+i <= len(s) { // 剩余长度可能不不足3位
			cur = s[i : size+i]
		} else {
			cur = s[i:]
			size = len(s) - i
		}

		if num, ok := g13[cur]; ok {
			ans += num
			i += size
			size = 3
		} else {
			size--
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRomanToInteger(t *testing.T) {
	//fmt.Println(romanToInt("III"))
	//fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
	//fmt.Println(romanToInt("CMXCIX"))
}
