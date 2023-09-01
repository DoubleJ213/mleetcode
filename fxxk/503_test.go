package fxxk

import (
	"fmt"
	"testing"
)

/*503. 下一个更大元素 II
给定一个循环数组nums（nums[nums.length - 1]的下一个元素是nums[0]），返回nums中每个元素的 下一个更大元素 。
数字 x的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，
这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

示例 1:
输入: nums = [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。

示例 2:
输入: nums = [1,2,3,4,3]
输出: [2,3,4,-1,4]

提示:
1 <= nums.length <= 104
-109<= nums[i] <= 109
*/

/*
这题和496有些区别，496不是环形的，最后一个数必然是-1，没有更大的数
这题可以考虑延长原数组，新的数组[nums1...,nums1...] -> numsNew
然后对numsNew套用模板
貌似可以通过修改索引位置来遍历新的数组，不用真的初始化一段内存来存这个新的数组
假如长度为n，那索引的范围  0～n-1 n~n+n-1
*/
func nextGreaterElements(nums []int) []int {
	greater := newStack()
	n := len(nums)
	res503 := make([]int, n)
	for i := 2*n - 1; i >= 0; i-- {
		// % 求模（余数）
		for !greater.Empty() && greater.Peek() <= nums[i%n] {
			greater.Pop()
		}
		if i <= n-1 {
			if greater.Empty() {
				res503[i] = -1
			} else {
				res503[i] = greater.Peek()
			}
		}
		greater.Push(nums[i%n])
	}
	return res503
}

func TestAl503(t *testing.T) {
	fmt.Printf("%v", nextGreaterElements([]int{1, 2, 3, 4, 3}))
}
