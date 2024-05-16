package fxxk

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func merge88(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if (m > 0 && nums2[n-1] > nums1[m-1]) || m == 0 {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeSortedArray(t *testing.T) {
	//	nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	//merge88([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	//merge88([]int{0}, 0, []int{1}, 1)
	merge88([]int{1}, 1, []int{0}, 0)
}
