package review

import (
	"fmt"
	"testing"
)

/*
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：
输入：digits = ""
输出：[]
示例 3：
输入：digits = "2"
输出：["a","b","c"]

提示：
0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。
*/

var res17 []string

var phoneMap17 = map[string]string{"2": "abc", "3": "def",
	"4": "ghi", "5": "jkl",
	"6": "mno", "7": "pqrs",
	"8": "tuv", "9": "wxyz"}

func letterCombinations17(digits string) []string {
	res17 = make([]string, 0)
	if len(digits) == 0 {
		return res17
	}
	traverse17(digits, 0, make([]byte, 0))
	return res17
}

func traverse17(digits string, index int, path []byte) {
	if len(path) == len(digits) {
		tmp := make([]byte, len(path))
		copy(tmp, path)
		res17 = append(res17, string(tmp))
		return
	}

	digit := phoneMap17[string(digits[index])]
	for j := 0; j < len(digit); j++ {
		traverse17(digits, index+1, append(path, digit[j]))
	}
}

func TestAl17(t *testing.T) {
	fmt.Println(letterCombinations17(""))
	fmt.Println(letterCombinations17("2"))
	fmt.Println(letterCombinations17("245"))
}
