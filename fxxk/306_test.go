package fxxk

import (
	"fmt"
	"strconv"
	"testing"
)

/*306. 累加数
累加数 是一个字符串，组成它的数字可以形成累加序列。
一个有效的 累加序列 必须 至少 包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。
给你一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是 累加数 。如果是，返回 true ；否则，返回 false 。
说明：累加序列里的数 不会 以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。

示例 1：
输入："112358"
输出：true
解释：累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
示例 2：
输入："199100199"
输出：true
解释：累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199

提示：1 <= num.length <= 35 num 仅由数字（0 - 9）组成
进阶：你计划如何处理由过大的整数输入导致的溢出?*/

// 使用 DFS 来求解，每次取一段，把它转换成一个数，
// 如果前面已经有两个数了，我们比较它与前两数之和的结果，
// 如果前面没有两个数或者比较结果相等，我们继续向下一层探索。
func isAdditiveNumber(num string) bool {
	return isAdditiveDfs(num, 0, 0, 0, 0)
}

func isAdditiveDfs(num string, index, count int, prepre, pre int64) bool {
	//index 最左边的索引
	//count 表示统计的次数
	//prepre 前前一个数
	//pre 前一个数
	if index >= len(num) {
		//终止条件，index递归每次加1，等于num长度及num全部递归完成
		return count > 2
	}
	var current int64
	//for循环横向，dfs纵向
	for i := index; i < len(num); i++ {
		levelNum, err := strconv.ParseInt(string(num[i]), 10, 0)
		if err != nil {
			return false
		}
		if string(num[index]) == "0" && i > index {
			//i == index 的时候 第一个数字为0，当成一个数处理是可以的
			//所以没必要继续处理01 或者02 这样i大于index的情况
			return false
		}
		current = current*10 + levelNum
		if count >= 2 {
			//只有统计过2次以上，才有必要比较当前数和前面2个数的和是否相同
			tmpSum := prepre + pre
			if current > tmpSum {
				//	当前的值大于前两项和 false
				return false
			}
			if current < tmpSum {
				// 当前值小于前两项的和，加位数，继续比较
				continue
			}
		} else {
			fmt.Print("count 小于2")
		}

		if isAdditiveDfs(num, i+1, count+1, pre, current) {
			//只有true的情况才return true，其余都是false
			return true
		}
	}
	return false
}

func TestAl305(*testing.T) {
	fmt.Print(isAdditiveNumber("12"))
	//fmt.Print(isAdditiveNumber("101123"))
}
