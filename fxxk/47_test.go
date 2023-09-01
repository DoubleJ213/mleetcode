package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

/*
47. 全排列 II
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：
输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
示例 2：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

提示：
1 <= nums.length <= 8
-10 <= nums[i] <= 10
*/
var res47 [][]int
var used47 []bool
var path47 []int

/*
相比于47，因为要对重复元素进行比较，如果在每次递归的时候删除一个数，新的数组就没有之前的数了。
那是不是再加个参数，把每次递归的数据传过去
*/
func permuteUnique(nums []int) [][]int {
	res47 = make([][]int, 0)
	used47 = make([]bool, len(nums))
	path47 = make([]int, 0)
	sort.Ints(nums)
	traverse47(nums, 0)
	return res47
}

func traverse47(nums []int, level int) {
	if level == len(nums) {
		tmp := make([]int, len(path47))
		copy(tmp, path47)
		res47 = append(res47, tmp)
		return
	}
	//递归进来不需要判断前一次是否重复，所以递归进来把前一次的数改掉，只需要在for循环中判断是否使用过
	for i := 0; i < len(nums); i++ {
		if i > 0 && !used47[i-1] && nums[i] == nums[i-1] {
			continue
		}
		if !used47[i] {
			used47[i] = true
			// 记录这条树枝上前一次的值
			path47 = append(path47, nums[i])
			traverse47(nums, level+1)
			path47 = path47[:len(path47)-1]
			used47[i] = false
		}
	}
}

func TestAl47(t *testing.T) {
	a := permuteUnique([]int{-1, 2, -1, 2, 1, -1, 2, 1})
	fmt.Printf("%v\n", a)
	fmt.Printf("len %v\n", len(a))
}
