package mleetcode

import (
	"fmt"
	"testing"
)

/*
给你一个整数数组nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：
输入：nums = [0]
输出：[[],[0]]

提示：
1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同
*/

func subsets(nums []int) [][]int {
	subRes := make([][]int, 0)
	subsetsDfs(nums, []int{}, &subRes)
	return subRes
}

func subsetsDfs(nums, levelRes []int, res *[][]int) {
	//每次都加到res中
	*res = append(*res, append([]int{}, levelRes...))

	if len(nums) == 0 {
		return
	}

	for index, value := range nums {
		subsetsDfs(nums[index+1:], append(levelRes, value), res)
	}
}

func TestAl78(t *testing.T) {
	a := subsets([]int{1, 2, 3})
	fmt.Print("abc")
	_ = a
}
