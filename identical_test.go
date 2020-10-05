package mleetcode

import (
	"testing"
)

//leetcode 1512
func numIdenticalPairs(nums []int) int {
	//双指针 index1外层，index2内层
	var index1, index2 int
	var num int
	for index1 < len(nums) && index2 < len(nums) {
		if index1 < index2 {
			// 原始条件 i<j
			if nums[index1] == nums[index2] {
				num++
			}
		}
		index2++
		if index2 == len(nums) {
			// index2遍历一遍，index1+1，同时i<j,index2直接从index1+1开始，少遍历几次
			index1++
			index2 = index1 + 1
		}
	}

	return num
}

func TestIdenticalPairs(t *testing.T) {
	num := []int{1, 1, 1, 1}
	numIdenticalPairs(num)

}
