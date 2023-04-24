package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

/*
18.四数之和
给你一个由n个整数组成的数组nums，和一个目标值target。
请你找出并返回满足下述全部条件且不重复的四元组[nums[a],nums[b],nums[c],nums[d]]：
0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a]+nums[b]+nums[c]+nums[d]==target
你可以按任意顺序返回答案。
输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
*/

var sumRes [][]int

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	sumRes = make([][]int, 0)
	sumDfs(nums, []int(nil), target, 0)
	return sumRes
}

func sumDfs(nums, cur []int, target, begin int) {
	if cur != nil && len(cur) == 4 {
		if target == 0 {
			tmp := append([]int(nil), cur...)
			combRes = append(combRes, tmp)
			return
		}
	}
	//
	//if target < 0 {
	//	return
	//}

	for i := begin; i < len(nums); i++ {
		num := nums[i]
		//剪枝，当前长度加上长度小于4时
		if len(nums)-i < 4-len(cur) {
			return
		}
		if i != 0 && num == nums[i-1] {
			continue
		}
		sumDfs(nums[i+1:], append(cur, num), target-num, begin)
	}

}

func TestAl18(t *testing.T) {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	//fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}
