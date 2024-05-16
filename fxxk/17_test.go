package fxxk

import (
	"fmt"
	"strings"
	"testing"
)

//17. 电话号码的字母组合
//给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

// var phoneMap map[string]string = map[string]string{
var phoneMap = map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}

var letterResult []string

func letterCombinations(digits string) []string {
	if strings.Contains(digits, "1") || len(digits) == 0 {
		return nil
	}
	letterResult = make([]string, 0)
	backtrack(digits, "")
	return letterResult
}

func backtrack(digits, res string) {
	//当题目中出现 “所有组合” 等类似字眼时，我们第一感觉就要想到用回溯。
	//定义函数 backtrack(res, digits)
	//当 digits 非空时，对于 digits[0] 中的每一个字母 letter
	//执行回溯 backtrack(res + letter,digits[1:]，直至 digits 为空。
	//最后将 res 加入到结果中。

	if len(digits) == 0 {
		letterResult = append(letterResult, res)
	} else {
		for _, value := range phoneMap[string(digits[0])] {
			backtrack(digits[1:], res+string(value))
		}
	}
	return
}

//贴一段python代码
/*
class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits: return []

        phone = {'2':['a','b','c'],
                 '3':['d','e','f'],
                 '4':['g','h','i'],
                 '5':['j','k','l'],
                 '6':['m','n','o'],
                 '7':['p','q','r','s'],
                 '8':['t','u','v'],
                 '9':['w','x','y','z']}

        def backtrack(conbination,nextdigit):
            if len(nextdigit) == 0:
                res.append(conbination)
            else:
                for letter in phone[nextdigit[0]]:
                    backtrack(conbination + letter,nextdigit[1:])

        res = []
        backtrack('',digits)
        return res
*/
func TestAl17(t *testing.T) {
	//fmt.Println(letterress1("245"))
	fmt.Println(letterCombinations("245"))
}
