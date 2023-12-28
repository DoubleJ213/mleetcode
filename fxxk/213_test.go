package fxxk

import (
	"fmt"
	"testing"
)

/*
213. 打家劫舍 II

你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。
这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

示例 1：
输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2：
输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 3：
输入：nums = [1,2,3]
输出：3

提示：
1 <= nums.length <= 100
0 <= nums[i] <= 1000
*/

var dp213First []int
var dp213Last []int
var max213 int

func rob213(nums []int) int {
	/*
		2,3,2 围成一圈 dp213[0] = 2  dp213[1] = 3 dp213[2] =2
		其余逻辑和198一样。判断到最后一个数的时候 不能取第一个，那岂不是dp213[0]这个数虽然算出来，但是到后面不能用了，那怎么递推呢
		整一个长点的数组看一下 2,3,2,2,2 dp0=2 dp1=3 dp2=dp0+2 dp3=dp1+2 dp4 (原本dp2+2，现在不对了。
		不是因为dp2不对，dp2 用到dp0进行递推了，因为自己是最后一个元素，不能用dp0推到的结果)
		所以要算递推到最后一个的情况，就不能加上第一个元素。
		除了要递推最后一个数，其余照旧。然后这两种情况再取最大值
	*/
	n := len(nums)
	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return getMax(nums[0], nums[1])
	} else if n == 3 {
		return getMax(getMax(nums[0], nums[1]), nums[2])
	} else if n == 4 {
		return getMax(nums[0]+nums[2], nums[1]+nums[3])
	}
	//先递推,不算最后一个元素 nums[:n-1], 再递推不偷第一个的 nums[1:]
	max213 = 0
	dp213First = make([]int, n-1)
	dp213First[0] = nums[0]
	dp213First[1] = nums[1]
	dp213First[2] = nums[0] + nums[2]
	for i := 3; i < n-1; i++ {
		dp213First[i] = getMax(dp213First[i-2], dp213First[i-3]) + nums[i]
		max213 = getMax(max213, dp213First[i])
	}
	dp213Last = make([]int, n-1)
	last := nums[1:]
	dp213Last[0] = last[0]
	dp213Last[1] = last[1]
	dp213Last[2] = last[0] + last[2]
	for j := 3; j < n-1; j++ {
		dp213Last[j] = getMax(dp213Last[j-2], dp213Last[j-3]) + nums[j]
		max213 = getMax(max213, dp213Last[j])
	}
	return max213
}

func TestAl213(t *testing.T) {
	//fmt.Println(rob213([]int{10, 1, 1, 2, 10}))
	//fmt.Println(rob213([]int{1, 2, 3, 1}))
	//fmt.Println(rob213([]int{10, 1, 1}))
	fmt.Println(rob213([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}
