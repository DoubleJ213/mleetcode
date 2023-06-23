package fxxk

import (
	"fmt"
	"testing"
)

/*
543. 二叉树的直径

给你一棵二叉树的根节点，返回该树的 直径 。
二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
两节点之间路径的 长度 由它们之间边数表示。
    1
  2   3
4  5
示例 1：
输入：root = [1,2,3,4,5]
输出：3
解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
示例 2：
输入：root = [1,2]
输出：1
提示：
树中节点数目在范围 [1, 104] 内
-100 <= Node.val <= 100
*/

var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
	diameter = 0
	if root == nil {
		return diameter
	}
	maxDepth(root)
	return diameter
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//左边子树最大值
	leftMax := maxDepth(root.Left)
	//右边子树最大值
	rightMax := maxDepth(root.Right)
	//整棵树的最大深度等于左右子树的最大深度取最大值，
	//然后再加上根节点自己
	diameter = getMax(diameter, leftMax+rightMax)
	return getMax(leftMax, rightMax) + 1
}

func TestAl543(t *testing.T) {
	tree1 := &TreeNode{Val: 15, Left: nil, Right: nil}
	tree2 := &TreeNode{Val: 7, Left: nil, Right: nil}
	tree3 := &TreeNode{Val: 20, Left: tree1, Right: tree2}
	tree4 := &TreeNode{Val: 9, Left: nil, Right: nil}
	root := &TreeNode{Val: 3, Left: tree4, Right: tree3}
	fmt.Printf("%d", diameterOfBinaryTree(root))
}
