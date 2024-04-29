package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)

/*
singleNumber
限制条件 on 时间  o1 空间 所以排除排序后比较的傻算可能性。
第一反应 a ^ a = 0 能排除a 但是这题是三个 同一个数，如果是2两个，直接无脑异或。
如果数组翻倍变成偶数个a 那 要找的b 也是偶数个b了，这题防止无脑异或了
但是怎么才能达到时间空间复杂度要求。
按照复杂度要求，大致逻辑肯定是。一个变量记录要返回的值，然后一遍循环，然后出结果。

1 <= nums.length <= 3 * 10^4
-2^31 <= nums[i] <= 2^31 - 1
*/
func singleNumber137(nums []int) int {
	//题目限制数组长度大于等于1
	var ans int32
	//比如 数组 2, 3, 2, 2
	//-2^31 <= nums[i] <= 2^31 - 1 32位 有符号 首位是 正负号
	for i := 0; i < 32; i++ {
		//MaxInt32  = 1<<31 - 1           // 2147483647 		//MinInt32  = -1 << 31
		var sum int
		for j := 0; j < len(nums); j++ {
			cur := nums[j]
			if (cur>>i)&1 == 1 {
				sum++
			}
		}

		if sum%3 == 1 {
			ans += 1 << i
		}
	}

	return int(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSingleNumberIi(t *testing.T) {
	//fmt.Println(singleNumber137([]int{2, 2, 3, 2}))
	fmt.Println(singleNumber137([]int{1, 1, -2, 1}))
	fmt.Println(singleNumber137([]int{2147483645, -2147483648, 2147483645, 2147483645}))

}
