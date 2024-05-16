package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
/*
3 <= nums.length <= 3000
-10^5 <= nums[i] <= 10^5
*/
var res15 [][]int

func threeSum(nums []int) [][]int {
	res15 = make([][]int, 0)
	sort.Ints(nums)
	if len(nums) < 3 || nums[0] > 0 {
		return nil
	}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else if nums[i]+nums[l]+nums[r] == 0 {
				res15 = append(res15, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] && nums[r] == nums[r+1] {
					l++
					r--
				}
			}
		}

	}
	return res15
}

//leetcode submit region end(Prohibit modification and deletion)

func TestThreeSum(t *testing.T) {
	//fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
	fmt.Println(threeSum([]int{0, 0, 0}))

}
