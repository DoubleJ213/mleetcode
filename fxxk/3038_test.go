package fxxk

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func maxOperations(nums []int) int {
	sum := nums[0] + nums[1]
	var res int
	for i := 1; i < len(nums); i += 2 {
		if nums[i-1]+nums[i] == sum {
			res++
		} else {
			break
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumNumberOfOperationsWithTheSameScoreI(t *testing.T) {

}
