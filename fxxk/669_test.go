package fxxk

import (
	"fmt"
	"testing"
)

/*
669. 修剪二叉搜索树
给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。
通过修剪二叉搜索树，使得所有节点的值在[low, high]中。
修剪树 不应该改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。
可以证明，存在唯一的答案。
所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。

示例 1：
输入：root = [1,0,2], low = 1, high = 2
输出：[1,null,2]
示例 2：
     3
  0     4
    2
   1
输入：root = [3,0,4,null,2,null,null,1], low = 1, high = 3
输出：[3,2,null,1]

提示：
树中节点数在范围 [1, 104] 内
0 <= Node.val <= 104
树中每个节点的值都是 唯一 的
题目数据保证输入是一棵有效的二叉搜索树
0 <= low <= high <= 104
*/

/*
二叉查找树（Binary Search Tree），（又：二叉搜索树，二叉排序树）
它或者是一棵空树，或者是具有下列性质的二叉树：
若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值；
若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值.
*/
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	traverse(root, low, high)
	return root
}

func traverse(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return nil
	}
	//if root.Val > high {
	//	return traverse(root.Left, low, high)
	//}
	//if root.Val < low {
	//	return traverse(root.Right, low, high)
	//}
	root.Left = traverse(root.Left, low, high)
	root.Right = traverse(root.Right, low, high)
	return root
}

func TestAl669(t *testing.T) {
	tree1 := &TreeNode{Val: 9, Left: nil, Right: nil}
	tree2 := &TreeNode{Val: 12, Left: nil, Right: nil}
	tree3 := &TreeNode{Val: 10, Left: tree1, Right: tree2}
	tree4 := &TreeNode{Val: 5, Left: nil, Right: nil}
	root := &TreeNode{Val: 8, Left: tree4, Right: tree3}
	a := trimBST(root, 1, 10)
	fmt.Printf("%v", a)
}

/*
func traverse(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = traverse(root.Left, low, high)
	root.Right = traverse(root.Right, low, high)
	return root
}

这样就相当于不做任何修改，直接全部递归完成
*/
