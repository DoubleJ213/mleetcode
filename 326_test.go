package mleetcode

import (
	"testing"
)

//给定一个整数，写一个函数来判断它是否是 3的幂次方。如果是，返回 true ；否则，返回 false 。
//整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x
//示例 1：输入：n = 27 输出：true
//示例 2：输入：n = 0 输出：false
//示例 3：输入：n = 9 输出：true
//示例 4：输入：n = 45 输出：false
//
//提示：
//-231 <= n <= 231 - 1

func isPowerOfThree(n int) bool {
	//一种方法傻算，不管是循环还是递归都是傻算，不写了
	//写一下位运算的解法
	//3，9，27,81对应2进制为11、1001、11011、1010001
	//不像2的阶乘那样，位运算没有什么规律。
	//突然想到一个很傻逼的思路，把所有小于2^31-1的值且为3的幂的都列出来，然后直接暴力判断不就可以了
	//计算器算了下，也就19种个数嘛，还能接受
	//threePower := []int{1, 3, 9, 27, 81, 243, 729, 2187, 6561, 19683, 59049, 177147, 531441, 1594323, 4782969, 14348907, 43046721, 129140163, 387420489, 1162261467}
	//if threePower {
	//}
	//想了想好像没有python 那种a in []的判断方法
	//那我给他初始化成map，然后判断key是否存在，不然是switch，那19行写起来也是够长的
	// 比傻算还傻逼的方法，加了map，再判断。但是性能好像也不咋滴，可以改成switch再试一试，但是想到一个更好的方法
	abc := map[int]bool{
		1:          true,
		3:          true,
		9:          true,
		27:         true,
		81:         true,
		243:        true,
		729:        true,
		2187:       true,
		6561:       true,
		19683:      true,
		59049:      true,
		177147:     true,
		531441:     true,
		1594323:    true,
		4782969:    true,
		14348907:   true,
		43046721:   true,
		129140163:  true,
		387420489:  true,
		1162261467: true,
	}
	if _, ok := abc[n]; ok {
		return true
	} else {
		return false
	}
}

func isPowerOfThree1(n int) bool {
	// 既然知道范围内最大的数是1162261467，那么他除以3，不能有余数
	return n > 0 && 1162261467%n == 0
}

func TestAl326(t *testing.T) {
	isPowerOfThree1(3 * 3 * 4)
}
