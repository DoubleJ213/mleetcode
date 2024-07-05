package fxxk

import (
	"fmt"
	"testing"
)

/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
*/
// leetcode submit region begin(Prohibit modification and deletion)
// 满足其要求的 正整数，从1开始，然后2  [1,2,3,4,5,7] 数字1  放到index0 处  数字2放到index1处
// 如果这个数字大于数组长度，说明缺了可能不止一个，这个数也不用排。如果小于0 不用交换处理
func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != i+1 {
			j := nums[i] - 1
			if nums[j] == j+1 { // 处理已经就位的
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFirstMissingPositive(t *testing.T) {
	fmt.Println(firstMissingPositive([]int{3, -1, -2, 5, -1, 1}))
	//fmt.Println(firstMissingPositive([]int{1, 2, 3}))
}

var v41 map[int]bool // 哈希表的大小与数组的长度是线性相关的，因此空间复杂度不符合题目要求。

func firstMissingPositive2(nums []int) int {
	v41 = make(map[int]bool)
	for _, num := range nums {
		if _, ok := v41[num]; ok {
			continue
		} else {
			v41[num] = true
		}
	}
	var ans int
	for ans = 1; ans < len(nums)+1; ans++ {
		if _, ok := v41[ans]; !ok {
			return ans
		}
	}
	return ans
}
