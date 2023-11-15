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
var max673 int

func findNumberOfLIS(nums []int) int {
	//dp[i] 对应nums[i] 这个数结尾的最长递增的序列的长度
	//dp[i] = dp[i]
	dp673 = make([]int, len(nums))
	max673 = 1
	dp673[0] = 1
	for i := 1; i < len(nums); i++ {
		//	当前的数不大于前面的数，不能在当前的基础上加1
		// 类似dp[0] 起码是等于1的。
		dp673[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp673[i] = getMax(dp673[i], dp673[j]+1)
			}
			if dp673[i] > max673 {
				max673 = dp673[i]
			}
		}
	}

	count673 := 0
	maxCount := 0
	for k := 0; k < len(dp673); k++ {
		if dp673[k] == max673-1 {
			count673++
		}
		if dp673[k] == max673 {
			maxCount++
		}
	}
	if count673 == 0 {
		return maxCount
	}
	return count673
}

func TestAl673(t *testing.T) {
	//fmt.Println(findNumberOfLIS([]int{1, 2, 5}))
	//fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findNumberOfLIS([]int{1, 2, 2, 2, 2, 3, 47}))
	//[10,55,45,25,25]
}
