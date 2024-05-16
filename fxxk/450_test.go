package fxxk

import (
	"fmt"
	"github.com/mlcPractice/leetcode/editor/cn"
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

/*
节点值唯一
root 是合法的二叉搜索树
要求算法时间复杂度为 O(h)，h 为树的高度。
*/
var head, parent *test.TreeNode

func deleteNode(root *test.TreeNode, key int) *test.TreeNode {
	head, parent = nil, nil
	if root == nil {
		return root
	}
	dfs450(root, root, key)
	if head == nil { //没找到直接返回
		return root
	}
	if head == parent {
		//根节点
		return subDeleteTree(head)
	} else {
		if head.Val < parent.Val {
			// 左子节点
			parent.Left = subDeleteTree(head)
		} else {
			//	右子节点
			parent.Right = subDeleteTree(head)
		}
		return root
	}
}

func subDeleteTree(node *test.TreeNode) *test.TreeNode {
	if node.Right == nil {
		return node.Left
	}
	if node.Left == nil {
		return node.Right
	}
	//左右子树都有值，重建node
	if node.Right.Left == nil {
		node, node.Right.Left = node.Right, node.Left
		return node
	}
	//node.Right.Left不为空，右子树不动，左子树找个最小的，变成新的root
	index := node.Right
	for index.Left != nil {
		index = index.Left
	}
	//index 指向 没有左子树的位置了
	node.Left, index.Left, node = nil, node.Left, node.Right
	return node
}

func dfs450(root *test.TreeNode, pre *test.TreeNode, key int) {
	if root == nil {
		return
	}
	if root.Val > key {
		dfs450(root.Left, root, key)
	}
	if root.Val < key {
		dfs450(root.Right, root, key)
	}
	if root.Val == key {
		head, parent = root, pre
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDeleteNodeInABst(t *testing.T) {
	var root = &test.TreeNode{5,
		&test.TreeNode{2,
			&test.TreeNode{1, nil, nil},
			&test.TreeNode{4,
				&test.TreeNode{3, nil, nil}, nil}},
		&test.TreeNode{7,
			&test.TreeNode{6, nil, nil},
			&test.TreeNode{8, nil, nil}}}
	a := deleteNode(root, 2)
	fmt.Println(a)

}
