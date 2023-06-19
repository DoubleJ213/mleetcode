package fxxk

import (
	"fmt"
	"testing"
)

/*
704. 二分查找
给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。

示例 1:
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例2:
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1

提示：
你可以假设 nums中的所有元素是不重复的。
n将在[1, 10000]之间。
nums的每个元素都将在[-9999, 9999]之间。
*/

func search(nums []int, target int) int {
	if nums == nil || len(nums) <= 0 {
		return -1
	}
	return bs(nums, target, 0, len(nums)-1)
}

func bs(nums []int, target, left, right int) int {
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	if right-1 == left || right == left {
		return -1
	}
	//mid := (left + right) / 2
	//防止溢出，用下面的写法
	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return bs(nums, target, left, mid)
	} else if nums[mid] < target {
		return bs(nums, target, mid, right)
	}
	return -1
}

func TestAl704(t *testing.T) {
	fmt.Printf("%v", search2([]int{5}, -5))
	fmt.Printf("%v", search2([]int{5}, 5))
	fmt.Printf("%v", search2([]int{-1, 0, 5}, 5))
	fmt.Printf("%v", search2([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Printf("%v", search2([]int{-1, 0, 3, 5, 9, 12}, 2))
}

func search2(nums []int, target int) int {
	if nums == nil || len(nums) <= 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
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
