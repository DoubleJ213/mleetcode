package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)

var status55 []bool

func canJump(nums []int) bool {
	status55 = make([]bool, len(nums))
	status55[0] = true
	for i := 0; i < len(nums); i++ {
		if status55[i] {
			if i == len(nums)-1 {
				return true
			}
			for j := 1; j <= nums[i]; j++ {
				status55[i+j] = true
				if status55[len(nums)-1] {
					return true
				}
			}
		}
	}
	return status55[len(nums)-1]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestJumpGame(t *testing.T) {
	//fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{1}))
}
