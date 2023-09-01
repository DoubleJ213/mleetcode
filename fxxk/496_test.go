package fxxk

import (
	"fmt"
	"testing"
)

/*
496. 下一个更大元素 I
nums1中数字x的 下一个更大元素 是指x在nums2 中对应位置 右侧 的 第一个 比x大的元素。
给你两个 没有重复元素 的数组nums1 和nums2 ，下标从 0 开始计数，其中nums1是nums2的子集。
对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，
并且在 nums2 确定 nums2[j] 的 下一个更大元素 。
如果不存在下一个更大元素，那么本次查询的答案是 -1 。
返回一个长度为nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。

示例 1：
输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
输出：[-1,3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
- 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
- 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
示例 2：
输入：nums1 = [2,4], nums2 = [1,2,3,4].
输出：[3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
- 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。

提示：
1 <= nums1.length <= nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 104
nums1和nums2中所有整数 互不相同
nums1 中的所有整数同样出现在 nums2 中

进阶：你可以设计一个时间复杂度为 O(nums1.length + nums2.length) 的解决方案吗？
*/

/*
暴力算法这里就不写了
单调队列 可以看这里 https://labuladong.github.io/algo/di-yi-zhan-da78c/shou-ba-sh-daeca/dan-diao-z-1bebe/
从数组尾部开始遍历，删除比自己小的元素

滑动窗口有点神似
为了体现出stack特性，这里实现下golang版本的stack
*/

type MyStack496 struct {
	nums []int
}

func (m *MyStack496) Pop() int {
	if m.Empty() {
		return -1
	}
	tmp := m.nums[len(m.nums)-1]
	m.nums = m.nums[:len(m.nums)-1]
	return tmp
}

func (m *MyStack496) Push(x int) {
	m.nums = append(m.nums, x)
}

func (m *MyStack496) Peek() int {
	if m.Empty() {
		return -1
	}
	return m.nums[len(m.nums)-1]
}

func (m *MyStack496) Empty() bool {
	if len(m.nums) == 0 {
		return true
	}
	return false
}

func newStack() *MyStack496 {
	return &MyStack496{
		nums: make([]int, 0),
	}
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	resMap := make(map[int]int)
	stack := newStack()
	for i := len(nums2) - 1; i >= 0; i-- {
		//nums1和nums2中所有整数 互不相同
		this := nums2[i]
		for !stack.Empty() && stack.Peek() < this {
			stack.Pop()
		}
		if stack.Empty() {
			resMap[this] = -1
		} else {
			resMap[this] = stack.Peek()
		}
		stack.Push(this)
	}
	result := make([]int, 0)
	for _, value := range nums1 {
		result = append(result, resMap[value])
	}

	return result
}

func TestAl496(*testing.T) {
	fmt.Printf("%v", nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
