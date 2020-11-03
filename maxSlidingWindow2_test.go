package mleetcode

import (
	"fmt"
	"testing"
)

// window中记录index，window最左边为窗口中最大的数
// 判断窗口大小的数据量，如果超过k个，移除最左边那个，如果不超过，下面的方法
// 如果新加入的数小于窗口最左边的数，则加入窗口
// 如果新加入的数大于窗口最左边的数，则窗口中的数全部移除，把当前的index放到window中
func maxSlidingWindow2(nums []int, k int) []int {
	if nums == nil {
		return []int{}
	}
	window := make([]int, 0)
	res := make([]int, 0)
	for i, x := range nums {
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}
		//len(window)是否一定大于0
		for len(window) > 0 && nums[window[len(window)-1]] <= x {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		fmt.Println(i, x, window)
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}

	}
	return res
}

func TestMaxSlidingWindow2(t *testing.T) {
	nums2 := []int{1, 4, -1, 3, 2, 3, 6, 7}
	nums := []int{7, 2, 4}
	nums3 := []int{9, 10, 9, -7, -4, -8, 2, -6}

	aaa := maxSlidingWindow2(nums2, 3)
	bbb := maxSlidingWindow2(nums, 2)
	ccc := maxSlidingWindow2(nums3, 5)

	fmt.Printf("%v", aaa)
	fmt.Printf("%v", bbb)
	fmt.Printf("%v", ccc)
}
