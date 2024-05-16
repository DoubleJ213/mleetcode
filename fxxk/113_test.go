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

var res113 [][]int

func pathSum113(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return res113
	}
	res113 = make([][]int, 0)
	traverse113(root, make([]int, 0), targetSum)
	return res113
}

func traverse113(root *TreeNode, path []int, sum int) *TreeNode {
	if root == nil {
		return root
	}
	left := traverse113(root.Left, append(path, root.Val), sum-root.Val)
	right := traverse113(root.Right, append(path, root.Val), sum-root.Val)
	if left == nil && right == nil {
		if sum == root.Val {
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, root.Val)
			res113 = append(res113, tmp)
		}
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPathSumIi(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{5, nil, nil}
	root5 := &TreeNode{6, nil, nil}
	root4 := &TreeNode{100, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	head := &TreeNode{-2, nil, &TreeNode{-3, nil, nil}}
	fmt.Println(pathSum113(head, -5))
	fmt.Println(pathSum113(root, 9))
	//	[-2,null,-3]
}
