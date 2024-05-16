package fxxk

import (
	"fmt"
	"testing"
)

/*
202. 快乐数
编写一个算法来判断一个数 n 是不是快乐数。
「快乐数」 定义为：
对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

示例 1：
输入：n = 19
输出：true
解释：
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
示例 2：
输入：n = 2
输出：false

提示：
1 <= n <= 2^31 - 1
*/

// 先想一下，如果这个数不是快乐数，这个数会不会越来越大。
// 可以通过是否陷入循环，是否是快乐数？
// 看答案了，判断循环，是否有环 就用快慢指针参考链表判断的那题
func isHappy(n int) bool {
	slow := getHappy(n)
	fast := getHappy(slow)
	for fast != slow {
		slow = getHappy(slow)
		fast = getHappy(getHappy(fast))
	}
	if fast != 1 {
		return false
	}
	return true
}

func getHappy(n int) int {
	sum := 0
	pre := n
	if pre < 10 {
		return pre * pre
	}
	for pre >= 10 {
		post := pre % 10
		sum += post * post
		pre = pre / 10
		if pre < 10 {
			sum += pre * pre
		}
	}
	return sum
}
func TestAl202(t *testing.T) {
	//fmt.Println(isHappy(10))
	//fmt.Println(isHappy(19))
	//fmt.Println(isHappy(2))
	fmt.Println(isHappy(88))
}
