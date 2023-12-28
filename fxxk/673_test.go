package fxxk

import (
	"fmt"
	"testing"
)

/*
673. 最长递增子序列的个数
给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。
注意 这个数列必须是 严格 递增的。

示例 1:
输入: [1,3,5,4,7]
输出: 2
解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
示例 2:
输入: [2,2,2,2,2]
输出: 5
解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。

提示:
1 <= nums.length <= 2000
-10^6 <= nums[i] <= 10^6
*/

//dp[i] 表示以 nums[i] 这个数结尾的最长递增子序列的长度
//返回最长递增子序列的个数 数组最后再遍历一遍，看看最大值有几个
var dp673 []int
var max673 []int

//max673[i] 为考虑以 nums[i] 结尾的最长上升子序列的个数。
var maxNum673 int

func findNumberOfLIS(nums []int) int {
	//dp[i] 对应nums[i] 这个数结尾的最长递增的序列的长度
	//额外定义 max673 为考虑以 nums[i] 结尾的最长上升子序列的个数。
	//dp[i] = dp[i]
	n := len(nums)
	dp673 = make([]int, n)
	max673 = make([]int, n)
	maxNum673 = 1
	dp673[0] = 1
	max673[0] = 1
	for i := 1; i < len(nums); i++ {
		//	当前的数不大于前面的数，不能在当前的基础上加1
		// 类似dp[0] 起码是等于1的。
		dp673[i] = 1
		max673[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp673[i] < dp673[j]+1 {
					dp673[i] = dp673[j] + 1
					max673[i] = max673[j]
				} else if dp673[i] == dp673[j]+1 {
					max673[i] += max673[j]
				}
			}
		}
		maxNum673 = getMax(maxNum673, dp673[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		if dp673[i] == maxNum673 {
			ans += max673[i]
		}
	}
	return ans

	/*
		dp673[i] 对应nums[i] 这个数结尾的最长递增的序列的长度
		max673[k]表示 对应nums[i] 最长递增子序列的个数，及能达到这个长度的所有的路径的个数
		max673[k] 怎么由 前面的项推算出来呢
		1.当前数在dp中比前一个数大，不需要处理，证明是一条路径算出来的 max673[k] = max673[k-1]
		2. dp当前的数不大于前面一个比如 12 3 34 那是两个3组成的两条路径能到3，那就是 max673[k] = max673[k-1] +1
		3. 突发发现 规律1 有问题。比如 4前面有2个3， 4得判断比他小的3 是不是 两条路都能走到4
			有可能有一个长度为3的尾部是比长度为4这个数大的，所以4得判断比他小的所有的3个数，然后是不是每个3都能走到4，有几个算几个
		所以逻辑应该是 max673[k] =  dp[k-1]值的个数,且是这些值对应的nums数小于nums[k]
	*/

	//https://leetcode.cn/problems/number-of-longest-increasing-subsequence/solutions/1007341/gong-shui-san-xie-lis-de-fang-an-shu-wen-obuz/

}

func TestAl673(t *testing.T) {
	//fmt.Println(findNumberOfLIS([]int{1, 2, 5}))
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	//fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 5}))
	//fmt.Println(findNumberOfLIS([]int{1, 2, 2, 2, 2, 3, 4}))
	//[10,55,45,25,25]
}
