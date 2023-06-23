package fxxk

import (
	"fmt"
	"testing"
)

/*
572. 另一棵树的子树
给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。
如果存在，返回 true ；否则，返回 false 。
二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。

示例 1：
输入：root = [3,4,5,1,2], subRoot = [4,1,2]
输出：true
示例 2：
输入：root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
输出：false
*/

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func TestAl573(t *testing.T) {
	tree1 := &TreeNode{Val: 1, Left: nil,
		Right: &TreeNode{Val: 1, Left: nil,
			Right: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil}}}
	tree11 := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil}
	fmt.Printf("%v", isSubtree(tree1, tree11))
}
