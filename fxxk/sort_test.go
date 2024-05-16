package fxxk

import (
	"fmt"
	"testing"
)

func TestQuickSort1(t *testing.T) {
	var nums = []int{5, 4, 3, 2, 1}
	//fmt.Println(quickSort(nums))
	//fmt.Println(quickSort1(nums))
	fmt.Println(mergeSort(nums))
	//fmt.Println(mergeSort1(nums))
	//fmt.Println(BubbleSort(nums))
}

func quickSort1(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	qs1(nums, 0, len(nums)-1)
	return nums
}

func qs1(nums []int, left int, right int) {
	if left >= right {
		return
	}
	tmp := nums[right]
	i, j := left, right
	for {
		for i < right && nums[i] <= tmp {
			i++
		}
		for j > left && nums[j] >= tmp {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			break
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	qs1(nums, left, i-1)
	qs1(nums, i, right)

}

func mergeSort1(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort1(nums[:mid])
	right := mergeSort1(nums[mid:])
	return merge1(left, right)
}

func merge1(left, right []int) []int {
	res := make([]int, len(left)+len(right))
	i, j, index := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res[index] = left[i]
			i++
		} else {
			res[index] = right[j]
			j++
		}
		index++
	}
	copy(res[index:], left[i:])
	copy(res[index+len(left)-i:], right[j:])
	return res
}

func BubbleSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	index := len(nums)
	for index > 0 {
		for i := 1; i < index; i++ {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
			}
		}
		index--
	}
	return nums
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))
	l, r, i := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}
		i++
	}
	//把最后的数直接追加上
	copy(result[i:], left[l:])
	copy(result[i+len(left)-l:], right[r:])
	return result
}

func quickSort(nums []int) []int {
	right := len(nums) - 1
	left := 0

	qs2(nums, left, right)
	return nums
}

func qs2(nums []int, left int, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	tmp := nums[right]
	for {
		for nums[i] <= tmp && i < right {
			i++
		}
		for nums[j] >= tmp && j > left {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			break
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	/*
		因为i已经和基准数进行交换了，i在完成最后一次交换之前还往上找了一个大于基准数的数
		确认下这个i是否一定等于j+1
	*/
	qs2(nums, left, i-1)
	qs2(nums, i, right)
}
