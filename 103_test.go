package mleetcode

import (
	"fmt"
	"testing"
)

// 103. 二叉树的锯齿形层序遍历

//给定一个二叉树，返回其节点值的锯齿形层序遍历。
//（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	//递归
	res := make([][]int, 0)
	// 判断条件不要忘
	if root == nil {
		return res
	}
	listTree := make([]*TreeNode, 0)
	listTree = append(listTree, root)
	for j := 0; len(listTree) > 0; j++ {
		//不要想当然写成下面这样
		//for j := 0 ; j < len(listTree); j++ {
		tmpTree := make([]*TreeNode, 0)
		curRes := make([]int, 0)
		for i := 0; i < len(listTree); i++ {
			curRes = append(curRes, listTree[i].Val)
			if listTree[i].Left != nil {
				tmpTree = append(tmpTree, listTree[i].Left)
			}
			if listTree[i].Right != nil {
				tmpTree = append(tmpTree, listTree[i].Right)
			}
		}
		if j%2 == 1 {
			for k := 0; k < len(curRes)/2; k++ {
				curRes[k], curRes[len(curRes)-k-1] = curRes[len(curRes)-k-1], curRes[k]
			}
		}
		listTree = tmpTree
		res = append(res, curRes)
	}
	return res
}

func TestAl103(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	//root5 := &TreeNode{5, nil, nil}
	//root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, nil, nil}
	root := &TreeNode{1, root2, root3}
	fmt.Println(zigzagLevelOrder(root))
	fmt.Println("done")

}
