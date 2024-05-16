package fxxk

import (
	"fmt"
	"testing"
)

/*
654. 最大二叉树
给定一个不重复的整数数组nums 。最大二叉树可以用下面的算法从nums 递归地构建:

创建一个根节点，其值为nums 中的最大值。
递归地在最大值左边的子数组前缀上构建左子树。
递归地在最大值 右边 的子数组后缀上构建右子树。
返回nums 构建的 最大二叉树 。

示例 1：
输入：nums = [3,2,1,6,0,5]
输出：[6,3,5,null,2,0,null,null,1]
解释：递归调用如下所示：
- [3,2,1,6,0,5] 中的最大值是 6 ，左边部分是 [3,2,1] ，右边部分是 [0,5] 。
  - [3,2,1] 中的最大值是 3 ，左边部分是 [] ，右边部分是 [2,1] 。
  - 空数组，无子节点。
  - [2,1] 中的最大值是 2 ，左边部分是 [] ，右边部分是 [1] 。
  - 空数组，无子节点。
  - 只有一个元素，所以子节点是一个值为 1 的节点。
  - [0,5] 中的最大值是 5 ，左边部分是 [0] ，右边部分是 [] 。
  - 只有一个元素，所以子节点是一个值为 0 的节点。
  - 空数组，无子节点。

示例 2：
输入：nums = [3,2,1]
输出：[3,null,2,null,1]

提示：
1 <= nums.length <= 1000
0 <= nums[i] <= 1000
nums 中的所有整数 互不相同
*/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	left, right, val := getMaxNum(nums)
	leftNode := constructMaximumBinaryTree(left)
	rightNode := constructMaximumBinaryTree(right)
	return &TreeNode{Left: leftNode, Right: rightNode, Val: val}
}

func getMaxNum(nums []int) (a, b []int, c int) {
	var max int
	var index int
	for i, num := range nums {
		if num > max {
			max = num
			index = i
		}
	}
	return nums[0:index], nums[index+1:], max
}

func TestAl654(t *testing.T) {
	a := constructMaximumBinaryTree([]int{1, 6, 2, 3, 7, 5, 4})
	fmt.Print(a)
	fmt.Println("easy as abc")
}
