package fxxk

import (
	"fmt"
	"reflect"
	"testing"
)

//101. 对称二叉树
//给定一个二叉树，检查它是否是镜像对称的。

//虽然是个简单题，但是这个解题有问题
//比如 [1,2,2,null,3,null,3] 这样的树是有问题的
//         1
//      2     2
//     3     3
//其实题目后面就有解释，这个数是不算对阵的。
//没看题目后续内容，看了个题目标题就开始写代码了，不可取。

func isSymmetric1(root *TreeNode) bool {
	left := root.Left
	right := root.Right
	lNum := make([]int, 0)
	rNum := make([]int, 0)
	symmetricLdfs(left, &lNum)
	symmetricRdfs(right, &rNum)
	return reflect.DeepEqual(lNum, rNum)
}

func symmetricRdfs(right *TreeNode, num *[]int) {
	if right == nil {
		return
	}
	*num = append(*num, right.Val)
	symmetricRdfs(right.Right, num)
	symmetricRdfs(right.Left, num)
}

func symmetricLdfs(left *TreeNode, num *[]int) {
	if left == nil {
		return
	}
	*num = append(*num, left.Val)
	symmetricLdfs(left.Left, num)
	symmetricLdfs(left.Right, num)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	//if left.Val != right.Val || (left == nil && right == nil) {
	// if left.Val != right.Val || (left == nil || right == nil) {
	// 为什么不是left||right为空
	//	其实应该解释下，都为空是true，有一个为空是false
	//}
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	//return (left.Val == right.Val) && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

func TestAl101(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	a := isSymmetric(root)

	fmt.Print("aa")
	_ = a
}
