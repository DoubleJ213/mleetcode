package fxxk

import (
	"fmt"
	"testing"
)

/*
LCR 083. 全排列
与46题相同
给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。

示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：
输入：nums = [1]
输出：[[1]]

提示：
1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/

var pathL83 [][]int

func permute(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return pathL83
	}
	pathL83 = make([][]int, 0)
	path := make([]int, 0)
	dfsL83(nums, path)
	return pathL83
}

func dfsL83(nums, path []int) {
	//无重不可复
	if len(nums) == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		pathL83 = append(pathL83, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		path = append(path, nums[i])
		newNums := make([]int, len(nums)-1)
		copy(newNums[:i], nums[:i])
		copy(newNums[i:], nums[i+1:])
		dfsL83(newNums, path)
		path = path[:len(path)-1]
	}
}

func TestLrc83(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}
