package fxxk

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func maximumPrimeDifference(nums []int) int {
	l := 0
	r := len(nums) - 1
	for !isTrue3115(nums[l]) && l <= r {
		l++
	}
	for !isTrue3115(nums[r]) && l <= r {
		r--
	}
	return r - l
}

func isTrue3115(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumPrimeDifference(t *testing.T) {
	//fmt.Println(maximumPrimeDifference([]int{4, 2, 9, 5, 31}))
	//isTrue(31)
}
