package review

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
var minR111 int

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minR111 = 10001
	traverseR111(root, 1)
	return minR111
}

func traverseR111(node *TreeNode, level int) *TreeNode {
	if node == nil {
		//	递归下去，发现空了，但是不能表示达到这个树的子树了最低端了，有可能另一个枝丫还有节点
		return nil
	}
	level++
	left := traverseR111(node.Left, level)
	right := traverseR111(node.Right, level)
	level--
	if left == nil && right == nil {
		minR111 = getMin(minR111, level)
	}
	return node
}

func TestR111(t *testing.T) {
	root2 := &TreeNode{2, nil, nil}
	root := &TreeNode{1, root2, nil}
	fmt.Println(minDepth(root))
}
