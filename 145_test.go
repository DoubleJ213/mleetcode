package mleetcode

import (
	"fmt"
	"testing"
)

// 二叉树，后序遍历

func postorderTraversal3(root *TreeNode) []int {
	listTree := make([]*TreeNode, 0)
	res := make([]int, 0)
	//需要保留当前节点的父节点。用于判断退出当前右子节点的逻辑。从当前栈中弹出当前的节点
	var prev *TreeNode
	for root != nil || len(listTree) > 0 {
		for root != nil {
			listTree = append(listTree, root)
			root = root.Left
		}
		root = listTree[len(listTree)-1]
		listTree = listTree[:len(listTree)-1]
		if root.Right == nil || root.Right == prev {
			//如果右子节点为空。那么将左子节点输出，设置右子节点。
			//因为已经准备回到上一层。设置当前root节点为空。从栈顶获取新的节点。

			//root.Right == prev意思为，如果右子树遍历完，打印当前节点
			if root.Val == 2 {
				fmt.Println("22")
			}
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			//如果当前节点的右节点不为空的时候，后序遍历就还需要将右节点的所有左子节点压入栈中。
			//所以重新将当前节点入栈，然后设置当前节点为右子节点
			//另外一种进入此判断的逻辑是
			listTree = append(listTree, root)
			root = root.Right
		}
	}
	return res

}

func TestAl145(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	fmt.Println(postorderTraversal3(root))
	fmt.Println("done")

}
