package fxxk

import (
	"fmt"
	"testing"
)

/*
258. 各位相加
给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。返回这个结果。
示例 1:
输入: num = 38
输出: 2
解释: 各位相加的过程为：
38 --> 3 + 8 --> 11
11 --> 1 + 1 --> 2
由于 2 是一位数，所以返回 2。
示例 2:
输入: num = 0
输出: 0

提示：
0 <= num <= 2^31 - 1
进阶：你可以不使用循环或者递归，在 O(1) 时间复杂度内解决这个问题吗？
*/

func addDigits(num int) int {
	sum := 0
	if num <= 0 {
		return 0
	}
	sum = getDigits(num)
	for sum >= 10 {
		sum = getDigits(sum)
	}
	return sum
}

func getDigits(num int) int {
	sum := 0
	pre := num
	if pre < 10 {
		return pre
	}
	for pre >= 10 {
		post := pre % 10
		sum += post
		pre = pre / 10
		if pre < 10 {
			sum += pre
		}
	}
	return sum
}

func TestAl258(t *testing.T) {
	fmt.Println(addDigits(38))
	//fmt.Println(addDigits(88))
}
