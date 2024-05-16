package fxxk

import (
	"fmt"
	"testing"
)

/*
1038. 从二叉搜索树到更大和树
给定一个二叉搜索树root(BST)，请将它的每个节点的值替换成树中大于或者等于该节点值的所有节点值之和。
提醒一下， 二叉搜索树 满足下列约束条件：
节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。

示例 1：
输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
示例 2：
输入：root = [0,null,1]
输出：[1,null,1]

提示：
树中的节点数在[1, 100]范围内。
0 <= Node.val <= 100
树中的所有值均 不重复。
*/

// 和538一样啊
var total int

func bstToGst(root *TreeNode) *TreeNode {
	total = 0
	traverse1038(root)
	return root
}

func traverse1038(root *TreeNode) {
	if root == nil {
		return
	}
	traverse1038(root.Right)
	total += root.Val
	root.Val = total
	traverse1038(root.Left)
}

func TestAl1038(*testing.T) {
	root6 := &TreeNode{1, nil, nil}
	root5 := &TreeNode{4, nil, nil}
	root4 := &TreeNode{2, root6, nil}
	root3 := &TreeNode{6, nil, nil}
	root2 := &TreeNode{3, root4, root5}
	root := &TreeNode{5, root2, root3}
	a := bstToGst(root)
	fmt.Println(a)
	fmt.Println("abc")
}
