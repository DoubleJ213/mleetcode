package mleetcode

import "testing"

/**
* 给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//
//func rangeSumBST(root *TreeNode, low int, high int) int {
//	root.Val
//}
//
//func TreeCreate(i int, arr []int) *TreeNode {
//	t := &TreeNode{nil,arr[i], nil}
//	if i < len(arr) && 2*i+1 < len(arr) {
//		t.Left = TreeCreate(2*i+1, arr)
//	}
//	if i < len(arr) && 2*i+2 < len(arr) {
//		t.Right = TreeCreate(2*i+2, arr)
//	}
//	return t
//}

func TestRange(t *testing.T) {

	//rangeSumBST(TreeCreate(1, []int{1, 2, 3, 4, 5}), 2, 4)
}
