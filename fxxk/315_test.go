package fxxk

import (
	"fmt"
	"testing"
)

/*
315. 计算右侧小于当前元素的个数
给你一个整数数组 nums ，按要求返回一个新数组counts 。
数组 counts 有该性质： counts[i] 的值是 nums[i] 右侧小于nums[i] 的元素的数量。

示例 1：
输入：nums = [5,2,6,1]
输出：[2,1,1,0]
解释：
5 的右侧有 2 个更小的元素 (2 和 1)
2 的右侧仅有 1 个更小的元素 (1)
6 的右侧有 1 个更小的元素 (1)
1 的右侧有 0 个更小的元素
示例 2：
输入：nums = [-1]
输出：[0]
示例 3：
输入：nums = [-1,-1]
输出：[0,0]

提示：
1 <= nums.length <= 105
-104 <= nums[i] <= 104
*/

/*
counts[i] 的值是 nums[i] 右侧小于nums[i] 的元素的数量
暴力求解两个for循环能解决，也不难不写了，但求别超时

归并，心态爆炸，写不下去了，下次再写
*/
var tmp []int

func countSmaller(nums []int) []int {
	tmp = make([]int, len(nums))
	left := 0
	right := len(nums) - 1
	sortArray315(nums, left, right)
	return nil
}

// 类似912那样写的话，没法在merge方法中知道当前的数在原数组的index
func sortArray315(nums []int, left, right int) {
	if left == right {
		return
	}
	mid := (right-left)/2 + left
	sortArray315(nums, left, mid)
	sortArray315(nums, mid+1, right)
	merge315(nums, left, mid, right)
}

// 合并left~mid 和mid+1～right
func merge315(nums []int, left, mid, right int) []int {

	return nil
}

func TestAl315(t *testing.T) {
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
}
