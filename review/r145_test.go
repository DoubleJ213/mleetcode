package review

import (
	"fmt"
	"testing"
)

// 二叉树，前中后序遍历
//迭代，多少遍才能熟练

func preorderTraversal(root *TreeNode) []int {
	res145 := make([]int, 0)
	toDoList := make([]*TreeNode, 0)
	for root != nil || len(toDoList) > 0 {
		for root != nil {
			res145 = append(res145, root.Val)
			toDoList = append(toDoList, root)
			root = root.Left
		}
		root = toDoList[len(toDoList)-1]
		root = root.Right
		toDoList = toDoList[:len(toDoList)-1]
	}
	return res145
}

func inorderTraversal(root *TreeNode) []int {
	res145 := make([]int, 0)
	toDoList := make([]*TreeNode, 0)
	for root != nil || len(toDoList) > 0 {
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
			res145, pre, root = append(res145, root.Val), root, nil
		} else {
			toDoList, root = append(toDoList, root), root.Right
		}
	}
	return res145
}

//2023.08.07 删除再写一遍
func TestAl145(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	//[1 2 4 5 3 6 7]
	//fmt.Println(preorderTraversal(root))
	//[4 2 5 1 6 3 7]
	//fmt.Println(inorderTraversal(root))
	//[4 5 2 6 7 3 1]
	fmt.Println(postorderTraversal(root))
	fmt.Println("done")
}
