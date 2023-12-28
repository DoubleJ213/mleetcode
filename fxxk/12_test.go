package fxxk

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

/*
12. 整数转罗马数字

罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。
27 写做  XXVII, 即为 XX + V + II 。
通常情况下，罗马数字中小的数字在大的数字的右边。
但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。
同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给你一个整数，将其转为罗马数字。

示例 1:
输入: num = 3
输出: "III"
示例 2:
输入: num = 4
输出: "IV"
示例 3:
输入: num = 9
输出: "IX"
示例 4:
输入: num = 58
输出: "LVIII"
解释: L = 50, V = 5, III = 3.
示例 5:
输入: num = 1994
输出: "MCMXCIV"
解释: M = 1000, CM = 900, XC = 90, IV = 4.

提示：
1 <= num <= 3999
*/

/*
XV 15
IV 4
12 写做 XII ，即为 X + II
9 表示为 IX。
*/

var romanMap12 = map[int]string{1: "A", 2: "AA", 3: "AAA", 5: "B", 6: "BA", 7: "BAA", 8: "BAAA",
	4: "AB", 9: "IX", 90: "XC", 900: "CM"}

var aMap = map[int]string{1: "I", 2: "X", 3: "C", 4: "M"}
var bMap = map[int]string{1: "V", 2: "L", 3: "D"}

var ans12 string

func intToRoman(num int) string {
	toRoman(num, 1, "")
	return ans12
}

func toRoman(num, deep int, path string) string {
	//先拆num
	//deep12 表示当前current是哪一位的 个位 1 十位 2
	//取余
	if num == 0 {
		ans12 = path
		return ans12
	}
	current := 0
	if num < 10 {
		current = num
	} else {
		current = num % 10
	}

	if current == 9 {
		current = current * int(math.Pow(float64(10), float64(deep-1)))
	}

	var currentString string
	if _, ok := romanMap12[current]; ok {
		currentString = romanMap12[current]
	}
	currentString = strings.ReplaceAll(currentString, "A", aMap[deep])
	currentString = strings.ReplaceAll(currentString, "B", bMap[deep])

	//9特殊处理，直接找9、90、900,省得进位找值，少个if else
	toRoman(num/10, deep+1, fmt.Sprintf("%s%s", currentString, path))
	return path
}

func TestAl12(t *testing.T) {
	fmt.Printf("II --- %v\n", intToRoman(2))
	//fmt.Printf("CDXLIV --- %v\n", intToRoman(444))
	//fmt.Printf("%v\n", intToRoman(30))
	fmt.Printf("LVIII --- %v\n", intToRoman(58))
	fmt.Printf("MCMXCIV --- %v\n", intToRoman(1994))
	//DCCCIII
}
