package mleetcode

import (
	"fmt"
	"testing"
)

// 给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

var result [][]int

// 执行代码与提交代码的结果不同？如果您在C/C++的代码中使用了全局变量，请看这里。

func TestAl107(t *testing.T) {
	//root7 := &TreeNode{7, nil, nil}
	//root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, root5, nil}
	root3 := &TreeNode{3, root4, nil}
	root2 := &TreeNode{2, root3, nil}
	root := &TreeNode{1, root2, nil}

	fmt.Println(levelOrderBottom(root))
	fmt.Println("done")
}

func levelOrderBottom(root *TreeNode) [][]int {
	result = make([][]int, 0)
	dfs2(root, 0)
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return result
}

func dfs2(root *TreeNode, i int) {
	if root == nil {
		return
	}
	if i == len(result) {
		result = append(result, []int{})
	}
	result[i] = append(result[i], root.Val)
	dfs2(root.Left, i+1)
	dfs2(root.Right, i+1)
}
