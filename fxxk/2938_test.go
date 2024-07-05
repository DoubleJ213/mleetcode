package fxxk

import (
	"fmt"
	"testing"
)

/*
1 和 0 分别代表黑色和白色的球
在每一步中，你可以选择两个相邻的球并交换它们。
返回「将所有黑色球都移到右侧，所有白色球都移到左侧所需的 最小步数」。
100 ->  010 -> 001
1100 ->  1010 -> 1001 -> 0101 -> 0011
*/
// leetcode submit region begin(Prohibit modification and deletion)
func minimumSteps(s string) int64 {
	var ans int64
	left := 0
	right := len(s) - 1
	for left < right {
		for s[left] == '0' && left < right {
			left++
		}
		for s[right] == '1' && left < right {
			right--
		}
		ans += int64(right - left)
		right--
		left++
		//string类型 这里不能真的交换 但是可以记录次数
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSeparateBlackAndWhiteBalls(t *testing.T) {
	fmt.Println(minimumSteps("1001"))
}
