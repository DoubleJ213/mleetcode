package review

import (
	"fmt"
	"testing"
)

/*
300. 最长递增子序列
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

示例 1：
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
示例 2：
输入：nums = [0,1,0,3,2,3]
输出：4
示例 3：
输入：nums = [7,7,7,7,7,7,7]
输出：1

提示：
1 <= nums.length <= 2500
-10^4 <= nums[i] <= 10^4

进阶：
你能将算法的时间复杂度降低到 O(n log(n)) 吗?
*/

//最长递增子序列 dp300[i] 代表 nums[i] 最长递增
//怎么推导出 dp300[i] 呢
// dp300[i] 遍历 nums[0~i-1] 所有元素，如 nums[i] 比当前index 大，
// 那dp300[i] 就 = dp300[index] +1 然后 dp300[i] 取所有的 dp300[i] 的最大值
var dp300 []int

func lengthOfLIS(nums []int) int {
	dp300 = make([]int, len(nums))
	dp300[0] = 1
	for i := 1; i < len(nums); i++ {
		dp300[i] = 1
		current := nums[i]
		for j := 1; j < i; j++ {
			if current > nums[j] {
				dp300[i] = getMax(dp300[i], dp300[j]+1)
			}
		}
	}
	return dp300[len(nums)-1]
}

func TestAl300(t *testing.T) {
	//fmt.Printf("3 --- %v", lengthOfLIS([]int{1, 2, 3}))
	fmt.Printf("4 --- %v", lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
