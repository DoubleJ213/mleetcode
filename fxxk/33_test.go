package fxxk

import (
	"fmt"
	"testing"
)

/*
时间复杂度为 O(log n) 的算法
153. 寻找旋转排序数组中的最小值
*/
// leetcode submit region begin(Prohibit modification and deletion)
func search33(nums []int, target int) int {
	// 先二分法找旋转的索引找到连接点，看target在左边还是右边，继续二分法找
	l := 0
	r := len(nums) - 1
	for l < r {
		m := (r-l)>>1 + l
		//nums[m] nums[l] nums[r] 其实是找最小的那个数，并返回最小数的索引  4, 5, 6, 7, 0, 1, 2, 3
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	//l 对应最小的数
	var left int
	var right int
	if target == nums[len(nums)-1] {
		return len(nums) - 1
	} else if target > nums[len(nums)-1] { // [0:l] 递增
		left = 0
		right = l - 1
	} else { // [l:] 递增
		left = l
		right = len(nums) - 1
	}
	for left <= right {
		mid := (right-left)>>1 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSearchInRotatedSortedArray(t *testing.T) {
	fmt.Println(search33([]int{4, 5, 6, 0, 1, 2, 3}, 0))
	fmt.Println(search33([]int{4, 5, 6, 7, 0, 1, 2, 3}, 2))
	fmt.Println(search33([]int{4, 5, 6, 7, 8, 1, 2, 3}, 0))
	fmt.Println(search33([]int{4, 5, 6, 7, 8, 9, 1, 2, 3}, 0))
}
