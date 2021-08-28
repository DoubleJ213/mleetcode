package mleetcode

import (
	"fmt"
	"testing"
)

// 二叉树，前序遍历

func preorderTraversal3(root *TreeNode) []int {
	listNode := make([]*TreeNode, 0)
	list := make([]int, 0)
	for root != nil || len(listNode) > 0 {
		for root != nil {
			list = append(list, root.Val)
			listNode = append(listNode, root)
			root = root.Left
		}
		root = listNode[len(listNode)-1].Right
		listNode = listNode[:len(listNode)-1]
	}
	return list
}

func TestAl114(t *testing.T) {
	//temp := &listNode{
	//	Val:  -1,
	//	Next: nil,
	//}
	//node := temp
	//for i := 0; i < 6; i++ {
	//	temp.Next = &listNode{i, nil}
	//	temp = temp.Next
	//}
	//fmt.Println(node.Next)

	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	fmt.Println(preorderTraversal3(root))

}
