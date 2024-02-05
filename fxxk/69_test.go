package fxxk

import (
	"fmt"
	"testing"
)

/*
69. x 的平方根
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。

示例 1：
输入：x = 4
输出：2
示例 2：
输入：x = 8
输出：2
解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。

提示：
0 <= x <= 2^31 - 1
*/

/*
找本数是哪个平方之间的数？比如给个10。平方根肯定在1~10之间吧。
然后先找中间(10-1)/2+1,比如5 25 大了，继续找小一点的
这个题目就变成，找一个最大的整数，其平方小于等于x。
*/
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	//比如x=6
	//0,1,2,3,4,5,6 之间？范围还能小，先这样试一试了。
	//确保index 和 对应数字一样大。省事
	left := 0
	right := x
	max := 1
	for left <= right {
		mid := (right-left)/2 + left
		if mid*mid == x {
			max = mid
			break
		} else if mid*mid < x {
			max = mid
			left = mid + 1
		} else if mid*mid > x {
			right = mid - 1
		}
	}
	return max
}

func TestAl69(t *testing.T) {
	fmt.Println(mySqrt(48))
}
