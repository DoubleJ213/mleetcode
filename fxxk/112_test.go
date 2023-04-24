package fxxk

import (
	"fmt"

	"testing"
)

/*
给你二叉树的根节点root 和一个表示目标和的整数targetSum ，
判断该树中是否存在 根节点到叶子节点 的路径，
这条路径上所有节点值相加等于目标和targetSum 。

叶子节点 是指没有子节点的节点

*/

//[1,2]
//1
//注意审题，这题是需要确认，根到叶子路径和是否等于target，
//不是下面解法，下面的解法是找路径是否满足target，不是根到叶子的路径

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var status bool
	hasPathDfs(root, targetSum, &status)
	return status
}

func hasPathDfs(root *TreeNode, sum int, status *bool) {
	if root == nil {
		return
	}

	if root.Left != nil {
		hasPathDfs(root.Left, sum-root.Val, status)
	}
	if root.Right != nil {
		hasPathDfs(root.Right, sum-root.Val, status)
	}
	//if sum == root.Val {
	//	*status = true
	//}
	if sum == root.Val && root.Right == nil && root.Left == nil {
		*status = true
	}
}

func TestAl112(*testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	aa := hasPathSum(root, 7)
	_ = aa
	fmt.Println("aa  ")
}
