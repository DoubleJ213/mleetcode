package review

import (
	"fmt"
	"testing"
)

/*
239. 滑动窗口最大值
给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。

示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
示例 2：
输入：nums = [1], k = 1
输出：[1]

提示：
1 <= nums.length <= 105
-104<= nums[i] <= 104
1 <= k <= nums.length
*/

func TestMaxSlidingWindow(t *testing.T) {
	//fmt.Printf("%v", maxSlidingWindow([]int{1, -1, 3, -3, 5, 3, 6, 7}, 3))
	//fmt.Printf("%v", maxSlidingWindow([]int{7, 2, 4}, 2))
	fmt.Printf("%v", maxSlidingWindow2([]int{9, 10, 9, -7, -4, -8, 2, -6}, 5))
}

//
//var window []int
//var result []int
//
//func maxSlidingWindow(nums []int, k int) []int {
//	if nums == nil || len(nums) == 0 || len(nums) < k {
//		return nil
//	}
//	//window 存的是索引
//	window = make([]int, 0)
//
//	for i := 0; i < k; i++ {
//		//nums[i] >= nums[window[len(window)-1]]
//		push(nums, i)
//	}
//
//	size := len(nums)
//	result = make([]int, 0)
//	result = append(result, nums[window[0]])
//	//1, 3, -1, -3, 5, 3, 6, 7
//	for i := k; i < size; i++ {
//		push(nums, i)
//
//		for window[0] <= i-k {
//			//由于窗口往后挪，最大的数过期了
//			window = window[1:]
//		}
//		result = append(result, nums[window[0]])
//	}
//	return result
//}
//
//func push(nums []int, i int) { //1, 3, -1, -3, 5, 3, 6, 7
//	//nums[i] >= nums[window[len(window)-1]]
//	// 当新加入的数大于窗口中记录的最后一个数。
//	//不断删除比自己小的，已存在的数，从后往前删的，自己有可能会在最左边的数过期后成为大哥
//	for len(window) > 0 && nums[i] >= nums[window[len(window)-1]] {
//		window = window[:len(window)-1]
//		//删除窗口中已存在的所有数
//	}
//	//新的索引 加入到窗口中
//	window = append(window, i)
//}

func maxSlidingWindow2(nums []int, k int) []int {
	if nums == nil || len(nums) == 0 || len(nums) < k {
		return nil
	}
	window := make([]int, 0)
	result := make([]int, 0)
	pushIntoWindow := func(i int) {
		for len(window) > 0 && nums[i] >= nums[window[len(window)-1]] {
			//窗口中已经有元素。且窗口最后一个没有自己大，把他干掉就行
			window = window[:len(window)-1]
		}
		//把自己加进去
		window = append(window, i)
	}
	for i := 0; i < k; i++ {
		//开始构造滑动窗口。前面K个元素，直接往里面丢就行。不需要考虑最前面元素过期的问题
		pushIntoWindow(i)
	}
	result = append(result, nums[window[0]])
	for i := k; i < len(nums); i++ {
		pushIntoWindow(i)
		if window[0] <= i-k {
			window = window[1:]
		}
		result = append(result, nums[window[0]])
	}
	return result
}
