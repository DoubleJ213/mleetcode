package fxxk

import (
	"fmt"
	"testing"
)

/*
402. 移掉 K 位数字
给你一个以字符串表示的非负整数num 和一个整数 k ，
移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。

示例 1 ：
输入：num = "1432219", k = 3
输出："1219"
解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219 。
示例 2 ：
输入：num = "10200", k = 1
输出："200"
解释：移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
示例 3 ：
输入：num = "10", k = 2
输出："0"
解释：从原数字移除所有的数字，剩余为空就是 0 。

提示：
1 <= k <= num.length <= 105
num 仅由若干位数字（0 - 9）组成
除了 0 本身之外，num 不含任何前导零
*/

type MyStack402 struct {
	nums []uint8
}

func (m *MyStack402) Pop() uint8 {
	if m.Empty() {
		return 47
	}
	tmp := m.nums[len(m.nums)-1]
	m.nums = m.nums[:len(m.nums)-1]
	return tmp
}

func (m *MyStack402) Push(x uint8) {
	m.nums = append(m.nums, x)
}

func (m *MyStack402) Peek() uint8 {
	if m.Empty() {
		return 47
	}
	return m.nums[len(m.nums)-1]
}

func (m *MyStack402) Empty() bool {
	if len(m.nums) == 0 {
		return true
	}
	return false
}

func newStack402() *MyStack402 {
	return &MyStack402{
		nums: make([]uint8, 0),
	}
}

func removeKdigits(num string, k int) string {
	if len(num)-k < 0 {
		return num
	}
	stack := []byte{}
	//给定一个长度为 n 的数字序列 [D0 D1 D2 D3 Dn-1]，
	//从左往右找到第一个位置 i（i>0）使得 Di<Di−1，
	//并删去Di−1;如果不存在，说明整个数字序列单调不降，删去最后一个数字即可。
	//上述过程操作一次，可以去除一个元素，k个元素删除，就操作k次。
	//这个算是暴力法了，如果超时怎么优化。
	//第一遍遍历，借助一个单调栈，并记录比本元素小的
	for i := range num {
		for len(stack) > 0 && stack[len(stack)-1] > num[i] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}

	stack = stack[:len(stack)-k]
	//输出不能有任何前导零
	index := 0
	for index < len(stack) && stack[index] == byte(48) {
		index++
	}
	//字符串为空时返回0
	result402 := string(stack[index:])
	if result402 == "" {
		return "0"
	}
	return result402
}

/*
字符串0 对应ascii 48
1-49 2-50 3-51 4-52
5-53 6-54 7-55 8-56 9-57
*/

func TestAl402(t *testing.T) {
	//a := removeKdigits("10200", 1)
	a := removeKdigits("10", 2)
	fmt.Printf("%v", a)
}
