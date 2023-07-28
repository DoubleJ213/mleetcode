package fxxk

import (
	"fmt"
	"testing"
)

/*
700. 二叉搜索树中的搜索
给定二叉搜索树（BST）的根节点root和一个整数值val。
你需要在 BST 中找到节点值等于val的节点。 返回以该节点为根的子树。 如果节点不存在，则返回null。

示例 1:
输入：root = [4,2,7,1,3], val = 2
输出：[2,1,3]
示例 2:
输入：root = [4,2,7,1,3], val = 5
输出：[]

提示：
数中节点数在[1, 5000]范围内
1 <= Node.val <= 107
root是二叉搜索树
1 <= val <= 107
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var getNode *TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	getNode = nil
	traverse700(root, val)
	return getNode
}

func traverse700(root *TreeNode, val int) {
	if root == nil {
		return
	}
	if root.Val > val {
		traverse700(root.Left, val)
	} else if root.Val < val {
		traverse700(root.Right, val)
	} else if root.Val == val {
		getNode = root
		return
	}
}

func TestAl700(t *testing.T) {
	root5 := &TreeNode{3, nil, nil}
	root4 := &TreeNode{7, nil, nil}
	root3 := &TreeNode{6, root5, root4}
	root2 := &TreeNode{4, nil, nil}
	root := &TreeNode{5, root2, root3}
	a := searchBST(root, 6)
	fmt.Println(a)
}
