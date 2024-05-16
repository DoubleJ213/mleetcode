package fxxk

import (
	"fmt"
	"testing"
)

/*
912. 排序数组
给你一个整数数组nums，请你将该数组升序排列。

示例 1：
输入：nums = [5,2,3,1]
输出：[1,2,3,5]
示例 2：
输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]

提示：
1 <= nums.length <= 5 * 104
-5 * 104 <= nums[i] <= 5 * 104
*/
func sortArray(nums []int) []int {
	//归并排序，有点类似二叉树后序遍历
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := sortArray(nums[0:mid])
	right := sortArray(nums[mid:])

	return merge912(left, right)
}

func merge912(left []int, right []int) []int {
	total := len(left) + len(right)
	merged := make([]int, total)
	l, r := 0, 0
	for l < len(left) || r < len(right) {
		if l == len(left) {
			copy(merged[l+r:], right[r:])
			break
		}
		if r == len(right) {
			copy(merged[l+r:], left[l:])
			break
		}
		leftNum := left[l]
		rightNum := right[r]
		if leftNum < rightNum {
			merged[l+r] = leftNum
			l++
		} else {
			merged[l+r] = rightNum
			r++
		}
	}
	return merged
}

func TestAl912(t *testing.T) {
	fmt.Println(sortArray([]int{2, 6, 7, 5, 1, 3, 4}))
}
