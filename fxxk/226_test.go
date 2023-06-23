package fxxk

import (
	"fmt"
	"testing"
)

/*
226. 翻转二叉树
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

示例 1：
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]
示例 2：
输入：root = [2,1,3]
输出：[2,3,1]
示例 3：
输入：root = []
输出：[]

提示：
树中节点数目范围在 [0, 100] 内
-100 <= Node.val <= 100
*/
func invertTree(root *TreeNode) *TreeNode {
	traverse3(root)
	return root
}

func traverse3(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := traverse3(root.Left)
	right := traverse3(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func TestAl226(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	fmt.Printf("%v", invertTree(root))
	fmt.Printf("11")
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree2(root.Left)
	right := invertTree2(root.Right)
	root.Left = right
	root.Right = left
	return root
}
