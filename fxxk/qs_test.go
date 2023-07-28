package fxxk

import "testing"

func quicksort(nums []int) []int {
	var left, right int
	left = 0
	right = len(nums) - 1
	qs(nums, left, right)
	return nums
}

func qs(nums []int, left int, right int) {
	if left >= right {
		return
	}
	temp := nums[right]
	i := left
	j := right
	for {
		for nums[i] <= temp && i < right {
			i++
		}
		for nums[j] >= temp && j > left {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			break
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	qs(nums, left, j)
	qs(nums, j+1, right)
}

func TestQuickSort(t *testing.T) {
	var nums = []int{2, 2}
	quicksort(nums)
}
