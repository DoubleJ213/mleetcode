package fxxk

import (
	"fmt"
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 平衡 二叉搜索树  nums 按 严格递增 顺序排列
func sortedArrayToBST108(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0], Left: nil, Right: nil}
	}
	right := len(nums) - 1
	left := 0
	return mySort108(nums, left, right)
}

func mySort108(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (right-left)/2 + left
	leftNode := mySort108(nums, left, mid-1)
	rightNode := mySort108(nums, mid+1, right)

	return &TreeNode{Val: nums[mid], Left: leftNode, Right: rightNode}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestConvertSortedArrayToBinarySearchTree(t *testing.T) {
	fmt.Println(sortedArrayToBST108([]int{-10, -3, 0, 5, 9}))
	//fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9, 10}))

}
