package fxxk

import (
	"fmt"
	"testing"
)

/*
53. 最大子数组和
给你一个整数数组nums，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组是数组中的一个连续部分。

示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：
输入：nums = [1]
输出：1
示例 3：
输入：nums = [5,4,-1,7,8]
输出：23

提示：
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4

进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解
*/

/*
就是找dp table
明确dp[i]的含义
用dp[i-1]  推导 dp[i]
dp[i] 表示 以 nums[i] 为结尾的「最⼤⼦数组和」
这种数组虽然不能直接算出来整个数组的最大子数组和，但是通过遍历一次dp求最大值就行

dp[i] 表示 以 nums[i] 为结尾的「最⼤⼦数组和」
那个如何通过用dp[i-1]  推导 dp[i]
因为是连续的，当前结尾。要么当前数是一个数组，要么和前面一个数结尾的数组组成另一个新的数组
*/
func maxSubArray(nums []int) int {
	dp53 := make([]int, len(nums))
	dp53[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		//这里可以发现dp[i]只可能和dp[i-1]有关，可以考虑优化下
		dp53[i] = getMax(dp53[i-1]+nums[i], nums[i])
	}

	res := -10001
	for j := 0; j < len(dp53); j++ {
		res = getMax(res, dp53[j])
	}
	return res
}

func TestAl53(t *testing.T) {
	//fmt.Printf("3 --- %v\n", maxSubArray([]int{-2, -1, -1, -2, 1, 2}))
	fmt.Printf("-2 --- %v\n", maxSubArray([]int{-2, -2}))
	//fmt.Printf("6 --- %v\n", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	//fmt.Printf("23 --- %v\n", maxSubArray([]int{5, 4, -1, 7, 8}))
	//fmt.Printf("1 --- %v\n", maxSubArray([]int{1}))
}

//-3, 2, -1, 4, -1, 2, 1, -5, 4
