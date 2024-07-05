package fxxk

import (
	"fmt"
	"testing"
)

/*
111. 二叉树的最小深度
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：2
示例 2：
输入：root = [2,null,3,null,4,null,5,null,6]
输出：5

提示：
树中节点数的范围在 [0, 10^5] 内
-1000 <= Node.val <= 1000
*/

var min111 int

/*
很多人写出的代码都不符合 1,2 这个测试用例，是因为没搞清楚题意
题目中说明:叶子节点是指没有子节点的节点，这句话的意思是 1 不是叶子节点
题目问的是到叶子节点的最短距离，所以所有返回结果为 1 当然不是这个结果
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min111 = 10001
	traverse4(root, 1)
	return min111
}

func traverse4(root *TreeNode, level int) {
	if root == nil {
		return
	}
	traverse4(root.Left, level+1)
	traverse4(root.Right, level+1)
	if root.Left == nil && root.Right == nil {
		min111 = getMin(min111, level)
	}
	return
}

func TestAl111(t *testing.T) {
	root2 := &TreeNode{2, nil, nil}
	root := &TreeNode{1, root2, nil}
	fmt.Println(minDepth(root))
}
