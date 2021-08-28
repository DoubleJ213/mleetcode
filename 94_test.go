package mleetcode

import (
	"fmt"
	"testing"
)

// 二叉树，中序遍历

func inorderTraversal3(root *TreeNode) []int {
	listNode := make([]*TreeNode, 0)
	res := make([]int, 0)
	for root != nil || len(listNode) > 0 {
		for root != nil {
			listNode = append(listNode, root)
			root = root.Left
		}
		root = listNode[len(listNode)-1]
		listNode = listNode[:len(listNode)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

func TestAl94(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	fmt.Println(inorderTraversal3(root))
	fmt.Println("done")

}
