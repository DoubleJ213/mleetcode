package fxxk

import (
	"fmt"
	"testing"
)

/*
100. 相同的树
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1：
输入：p = [1,2,3], q = [1,2,3]
输出：true
示例 2：
输入：p = [1,2], q = [1,null,2]
输出：false
示例 3：

输入：p = [1,2,1], q = [1,1,2]
输出：false
*/
var isSame bool

func isSameTree1(p *TreeNode, q *TreeNode) bool {
	isSame = true
	sameTree(p, q)
	return isSame
}

func sameTree(p *TreeNode, q *TreeNode) {
	if p == nil || q == nil {
		if p != q {
			isSame = false
		}
		return
	}
	if p.Val != q.Val {
		isSame = false
		return
	}
	sameTree(p.Left, q.Left)
	sameTree(p.Right, q.Right)
}

func TestAl100(t *testing.T) {
	tree1 := &TreeNode{Val: 9, Left: nil, Right: nil}
	tree11 := &TreeNode{Val: 8, Left: nil, Right: nil}
	tree2 := &TreeNode{Val: 12, Left: nil, Right: nil}
	tree3 := &TreeNode{Val: 10, Left: tree1, Right: tree2}
	tree4 := &TreeNode{Val: 5, Left: nil, Right: nil}
	root := &TreeNode{Val: 8, Left: tree4, Right: tree3}
	root2 := &TreeNode{Val: 8, Left: tree4, Right: tree11}
	root3 := &TreeNode{Val: 9, Left: tree4, Right: tree3}
	fmt.Println(isSameTree2(root, root2))
	fmt.Println(isSameTree(root, root3))
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return isSameTree2(p.Left, q.Left) && isSameTree2(p.Right, q.Right)
}
