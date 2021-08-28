package mleetcode

import (
	"fmt"
	"testing"
)

// 二叉树，BFS

func levelOrder2(root *TreeNode) [][]int {
	// 迭代 实现
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	listTree := make([]*TreeNode, 0)
	listTree = append(listTree, root)
	for len(listTree) > 0 {
		currTreeList := make([]*TreeNode, 0)
		res := make([]int, 0)
		for i := 0; i < len(listTree); i++ {
			res = append(res, listTree[i].Val)
			if listTree[i].Left != nil {
				currTreeList = append(currTreeList, listTree[i].Left)
			}
			if listTree[i].Right != nil {
				currTreeList = append(currTreeList, listTree[i].Right)
			}
		}
		result = append(result, res)
		listTree = currTreeList
	}

	return result
}

func TestAl102_2(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	fmt.Println(levelOrder2(root))
	fmt.Println("done")

}
