package fxxk

import (
	"fmt"
	"math"
	"testing"
)

/*
515. 在每个树行中找最大值
给定一棵二叉树的根节点root ，请找出该二叉树中每一层的最大值。

示例1：
输入: root = [1,3,2,5,3,null,9]
输出: [1,3,9]
示例2：
输入: root = [1,2,3]
输出: [1,3]

提示：
二叉树的节点个数的范围是 [0,104]
-231<= Node.val <= 231- 1
*/
var data []int

func largestValues(root *TreeNode) []int {
	//data = make([][]int, 0)
	data = make([]int, 0)
	if root == nil {
		return nil
	}
	traverse2(root, 0)
	return data
}

func traverse2(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if len(data) <= level {
		data = append(data, int(math.Pow(-2, 31)))
	}
	if root.Val > data[level] {
		data[level] = root.Val
	}

	traverse2(root.Left, level+1)
	traverse2(root.Right, level+1)
}

func TestAl515(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}

	fmt.Printf("%v", largestValues(root))
}
