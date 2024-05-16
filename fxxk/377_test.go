package fxxk

import (
	"fmt"
	"math"
	"testing"
)

/*
377. 组合总和 Ⅳ
给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target
请你从 nums 中找出并返回总和为 target 的元素组合的个数。
题目数据保证答案符合 32 位整数范围。

示例 1：
输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。
示例 2：
输入：nums = [9], target = 3
输出：0
*/

var (
	res377  [][]int
	path377 []int
)

func combinationSum4(nums []int, target int) int {
	//无重可复
	res377 = make([][]int, 0)
	path377 = make([]int, 0)
	traverse377(nums, target)
	fmt.Printf("%v\n", res377)
	return len(res377)
}

func traverse377(nums []int, target int) {
	if target == 0 {
		//[4, 2, 1]  32
		//out of memory
		tmp := make([]int, len(path377))
		copy(tmp, path377)
		res377 = append(res377, tmp)
	}
	if target < 0 {
		return
	}
	for i := 0; i < len(nums); i++ {
		path377 = append(path377, nums[i])
		if len(path377) > math.MaxInt32 {
			return
		}
		traverse377(nums, target-nums[i])
		path377 = path377[:len(path377)-1]
	}
}

/*
上面暴力求解，递归的是一个满二叉树，可能会超时的，提交后发现内存溢出了。
优化成自底向上的循环，带备忘录，子问题
target = i 所有可能的组合

	[1,2,3]

dp[i] = dp[i-1]  + dp[i-2] + dp[i-3]

dp[0] = 1
dp[1] = 1
dp[2] = 2 ([1,1] [0,2]) dp[2-1] dp[2-2]
dp[3] = dp[3-1]+dp[3-2]+dp[3-3] = 1+1+2 =4
dp[4] = dp[4-3]+ dp[4-2] +dp[4-1] = 1+2+4 = 7
*/
func combinationSum4_1(nums []int, target int) int {
	//无重可复
	dp377 := make([]int, target+1)
	dp377[0] = 1
	for i := 0; i < len(dp377); i++ {
		for j := 0; j < len(nums); j++ {
			if i-nums[j] >= 0 {
				dp377[i] = dp377[i] + dp377[i-nums[j]]
			}
		}
	}
	return dp377[target]
}

func TestAl377(t *testing.T) {
	//fmt.Printf("%v", combinationSum4([]int{1, 2, 3}, 4))
	//fmt.Printf("%v", combinationSum4_1([]int{4, 2, 1}, 10))
	fmt.Printf("%v", combinationSum4_1([]int{4, 2, 1}, 32))
}

//[4,2,1]
//32
