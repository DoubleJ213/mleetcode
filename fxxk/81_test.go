package fxxk

import (
	"fmt"
	"testing"
)

/*
之前33 题，升序 不重复 然后可以返回具体index 这题 可以重复，所以返回是否true or false
其余思路差不多，先找最小数，然后看是在左边还是右边，继续二分法找。找到了提前 true 找不到 false
*/
// leetcode submit region begin(Prohibit modification and deletion)
func search81(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1
	for l < r {
		for nums[l] == nums[l+1] && l+1 < r {
			l++
		}
		for nums[r] == nums[r-1] && r-1 > l {
			r--
		}
		m := (r-l)>>1 + l
		if nums[m] == target {
			return true
		}
		if nums[m] > nums[r] { // nums[m] > nums[r]  找最小数，不是找比target大还是小
			//找最小的数，有可能最小数有多个，找最小数，最左边的索引
			l = m + 1
		} else {
			r = m
		}
	}
	//nums[0:l] nums[l:]
	var left, right int
	if nums[len(nums)-1] == target {
		return true
	} else if nums[len(nums)-1] > target {
		left = l
		right = len(nums) - 1
	} else if nums[len(nums)-1] < target { // 最后一个数比目标值小，那目标值就在前面一段
		left = 0
		right = l - 1
	}
	for left <= right {
		mid := (right-left)>>1 + left
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSearchInRotatedSortedArrayIi(t *testing.T) {
	//fmt.Println(search81([]int{4, 5, 6, 6, 7, 0, 0, 1, 2, 3}, 3))
	//fmt.Println(search81([]int{2, 4, 4, 4, 4, 4, 4, 5, 6, 0, 0, 1, 2}, 1))
	//fmt.Println(search81([]int{1, 0, 0, 0, 0, 0, 1, 1, 1}, 0))
	fmt.Println(search81([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2))
	fmt.Println(search81([]int{1, 1}, 0))
}
