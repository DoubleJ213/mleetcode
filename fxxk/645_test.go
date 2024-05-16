package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

/*
645. 错误的集合
集合 s 包含从 1 到 n 的整数。不幸的是，因为数据错误，
导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合 丢失了一个数字 并且 有一个数字重复 。
给定一个数组 nums 代表了集合 S 发生错误后的结果。
请你找出重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。

示例 1：
输入：nums = [1,2,2,4]
输出：[2,3]
示例 2：
输入：nums = [1,1]
输出：[1,2]

提示：
2 <= nums.length <= 10^4
1 <= nums[i] <= 10^4
*/

// 首先 搞一个map 记录每个数出现次数。空间复杂度o(n)不是很好
// 数字是从1到n，然后只错了一个 数字从1到n 没说是有序的
func findErrorNums1(nums []int) []int {
	sort.Ints(nums)
	pre := 0
	ans654 := make([]int, 2)
	//重复的整数ans654[0] 丢失的整数  ans654[1]
	for i := 0; i < len(nums); i++ {
		current := nums[i]
		if current != pre {
			if pre+1 != current {
				ans654[1] = pre + 1
			}
			pre = current
		} else {
			ans654[0] = pre
		}
	}
	if ans654[1] == 0 {
		ans654[1] = len(nums)
	}
	return ans654
}

// 这空间降下去了，时间复杂度还很高，异或的思路貌似也可以
func findErrorNums(nums []int) []int {
	ans654 := make([]int, 2)
	//1, 3, 2, 3, 4 输入 异或下 index+1 1 2 3 4 5 最后剩余
	//0, 3, 0, 0, 0 ^5

	//1, 1  0+1 1+1
	// 1 ^1 =0   1 ^ 2
	// 遍历到最后一个元素，如果当前是0当前元素就是多的

	return ans654
}

func TestAl645(t *testing.T) {
	fmt.Println(findErrorNums([]int{1, 3, 2, 3, 4})) // [3,5]
	fmt.Println(findErrorNums([]int{1, 3, 4, 3, 5})) // [3,2]
	fmt.Println(findErrorNums([]int{1, 1}))          // [3,2]
}
