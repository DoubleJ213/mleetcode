package fxxk

import (
	"fmt"
	"testing"
)

/*
26. 删除有序数组中的重复项

给你一个升序排列的数组nums，请你原地删除重复出现的元素，
使每个元素只出现一次，返回删除后数组的新长度。
元素的相对顺序应该保持一致。

由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。
更规范地说，如果在删除重复项之后有k个元素，那么nums的前k个元素应该保存最终结果。
将最终结果插入nums的前k个位置后返回k。
不要使用额外的空间，你必须在原地修改输入数组并在使用O(1)额外空间的条件下完成。
*/

//LeetCode 中文屎一样的描述. 上面这个有问题。
//返回新的数组长度就行.但是要保证原来的nums值是去重后的新的数组的值

func TestAl26(t *testing.T) {
	fmt.Println(removeDuplicates([]int{1, 1, 1}))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var slow, fast, step int
	for fast < len(nums) {
		if nums[fast] == nums[slow] {
			fast++
		} else {
			step++
			nums[step] = nums[fast]
			slow = fast
			fast++
		}
	}
	fmt.Println(nums[:step+1])
	return step + 1
}
