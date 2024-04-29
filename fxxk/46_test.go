package fxxk

import (
	"fmt"
	"testing"
)

/*
46. 全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

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
/*var res46 [][]int

func permute(nums []int) [][]int {
	res46 = make([][]int, 0)
	traverse46(nums, make([]int, 0))
	return res46
}

func traverse46(nums, current []int) {
	if len(nums) == 0 {
		res46 = append(res46, current)
		return
	}
	for i := 0; i < len(nums); i++ {
		traverse46(getNewNums(nums, i), append(current, nums[i]))
	}
}

func getNewNums(nums []int, i int) []int {
	res := make([]int, len(nums)-1)
	copy(res[0:i], nums[0:i])
	copy(res[i:], nums[i+1:])
	return res
}
*/
func TestAl46(t *testing.T) {
	fmt.Printf("%v\n", permute2([]int{1, 2, 3}))
	fmt.Printf("%v\n", permute2([]int{0, 1}))
}

var res46 [][]int
var path46 []int
var used46 []bool

func permute2(nums []int) [][]int {
	res46 = make([][]int, 0)
	path46 = make([]int, 0)
	used46 = make([]bool, len(nums))
	traverse46_1(nums, 0)
	return res46
}

func traverse46_1(nums []int, level int) {
	if len(nums) == level {
		tmp := make([]int, len(path46))
		copy(tmp, path46)
		res46 = append(res46, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used46[i] {
			continue
		}
		used46[i] = true
		path46 = append(path46, nums[i])
		traverse46_1(nums, level+1)
		path46 = path46[:len(path46)-1]
		used46[i] = false
	}
}

var ans46 [][]int
var level46 []int

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute46a(nums []int) [][]int {
	n := len(nums)
	if n <= 0 {
		return nil
	}
	ans46 = make([][]int, 0)
	level46 = make([]int, 0)
	traverse46(nums)
	return ans46
}

func traverse46(nums []int) {
	if len(nums) == 0 {
		tmp := make([]int, len(level46))
		copy(tmp, level46)
		ans46 = append(ans46, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		newNum := getNewNum(nums, i)
		level46 = append(level46, nums[i])
		traverse46(newNum)
		level46 = level46[:len(level46)-1]
	}
}

func getNewNum(nums []int, skip int) []int {
	result47 := make([]int, len(nums)-1)
	copy(result47[:skip], nums[:skip])
	copy(result47[skip:], nums[skip+1:])
	return result47
}
