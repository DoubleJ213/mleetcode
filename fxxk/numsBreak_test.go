package fxxk

import (
	"fmt"
	"testing"
)

func numsBreak(nums []int) int {
	var ret int
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i+1] >= nums[i] {
			continue
		}
		divisor := (nums[i] - 1) / nums[i+1]
		ret += divisor
		nums[i] /= divisor + 1
	}
	return ret
}

func TestNumsBreak(t *testing.T) {
	nums1 := []int{7, 9, 8}
	fmt.Printf("break %v", numsBreak(nums1))
	nums2 := []int{9, 1}
	fmt.Printf("break %v", numsBreak(nums2))
	nums3 := []int{1, 7, 9, 8}
	fmt.Printf("break %v", numsBreak(nums3))
	nums4 := []int{5, 12, 14, 8}
	fmt.Printf("break %v", numsBreak(nums4))
}
