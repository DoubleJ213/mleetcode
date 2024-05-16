package fxxk

import (
	"fmt"
	"testing"
)

// 104. 二叉树的最大深度
func maxDepth2(root *TreeNode) int {
	//BFS
	if root == nil {
		return 0
	}
	var level int
	listTree := make([]*TreeNode, 0)
	listTree = append(listTree, root)
	for level = 0; len(listTree) > 0; level++ {
		tmpList := make([]*TreeNode, 0)
		for i := 0; i < len(listTree); i++ {
			if listTree[i].Left != nil {
				tmpList = append(tmpList, listTree[i].Left)
			}
			if listTree[i].Right != nil {
				tmpList = append(tmpList, listTree[i].Right)
			}
		}
		listTree = tmpList
	}
	return level
}

func TestAl104(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	fmt.Println(maxDepth2(root))
	fmt.Println("done")
}
