package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func nextPermutation(nums []int) {
	// [1,2,5,1] [1,5,1,2]
	//从后往前，开始找，找当前数右边的，比当前数大，但是是最小的大于当前数的
	last := -1
	left := -1
	for r := len(nums) - 2; r >= 0; r-- {
		pivlot := nums[r]
		min := 101
		for i := len(nums) - 1; i > r; i-- {
			index := nums[i]
			if index <= pivlot {
				continue
			} else {
				if min > index {
					//	找到比当前数大的数，还得找最小的大于当前数的
					last = i
					min = index
				}
			}
		}
		if last != -1 {
			nums[r], nums[last] = min, nums[r]
			left = r
			break
		}
	}
	//	再排序 排从 left 到 len[nums]-1 /// left ~ right 排序
	right := len(nums) - 1
	if right <= left+1 {
		return
	}
	//quickSort 相当于二叉树前序遍历 把当前的数找到合适的位置，然后再排他的左\右子树
	quickSort31(nums, left+1, right)
}

func quickSort31(nums []int, left, right int) {
	if left >= right {
		return
	}
	r := right
	l := left
	for l < r {
		for nums[r] >= nums[right] && r > left {
			r--
		}
		for nums[l] <= nums[right] && right > l {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		} else {
			break
		}
	}
	nums[l], nums[right] = nums[right], nums[l]
	quickSort31(nums, left, l-1)
	quickSort31(nums, l+1, right)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNextPermutation(t *testing.T) {
	a := []int{1, 1, 2, 5, 6, 1, 3}
	b := []int{1, 2, 3}
	c := []int{3, 2, 1}
	d := []int{1, 1, 5}
	nextPermutation(a)
	nextPermutation(b)
	nextPermutation(c)
	nextPermutation(d)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	//quickSort31(a, 0, len(a)-1)
	//fmt.Println(a)
}
