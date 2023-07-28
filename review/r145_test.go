package review

import (
	"fmt"
	"testing"
)

// 二叉树，前中后序遍历
//迭代，多少遍才能熟练
//前序遍历，根节点在前面
func preorderTraversal(root *TreeNode) []int {
	res145 := make([]int, 0)
	if root == nil {
		return res145
	}
	toDoList := make([]*TreeNode, 0)
	for len(toDoList) > 0 || root != nil {
		for root != nil {
			res145 = append(res145, root.Val)
			toDoList = append(toDoList, root)
			root = root.Left
		}
		root = toDoList[len(toDoList)-1].Right
		toDoList = toDoList[:len(toDoList)-1]
	}
	return res145
}

//中序遍历 根节点在中间
func inorderTraversal(root *TreeNode) []int {
	res145 := make([]int, 0)
	toDoList := make([]*TreeNode, 0)
	for len(toDoList) > 0 || root != nil {
		for root != nil {
			toDoList = append(toDoList, root)
			root = root.Left
		}
		root = toDoList[len(toDoList)-1]
		res145 = append(res145, root.Val)
		root = root.Right
		toDoList = toDoList[:len(toDoList)-1]
	}
	return res145
}

func postorderTraversal(root *TreeNode) []int {
	res145 := make([]int, 0)
	toDoList := make([]*TreeNode, 0)
	var pre *TreeNode
	for root != nil || len(toDoList) > 0 {
		for root != nil {
			toDoList = append(toDoList, root)
			root = root.Left
		}
		root = toDoList[len(toDoList)-1]
		toDoList = toDoList[:len(toDoList)-1]
		if root.Right == nil || root.Right == pre {
			res145 = append(res145, root.Val)
			pre = root
			root = nil
		} else {
			toDoList = append(toDoList, root)
			root = root.Right
		}
	}
	return res145
}

func TestAl145(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	//fmt.Println(preorderTraversal(root))
	//fmt.Println(inorderTraversal(root))
	fmt.Println(postorderTraversal(root))
	fmt.Println("done")
}
