package fxxk

import (
	"fmt"
	"testing"
)

/*
124. 二叉树中的最大路径和
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。
示例 1：

	1

2    3
输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
示例 2：

	   -10
	9      20
	    15     7

输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
提示：
树中节点数目范围是 [1, 3 * 104]
-1000 <= Node.val <= 1000
*/
var max int

func maxPathSum(root *TreeNode) int {
	//考虑val取值范围，0已经不是最小的数咯
	max = -1000
	getSideMax(root)
	return max

}

func getSideMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := getMax(0, getSideMax(root.Left))
	rightMax := getMax(0, getSideMax(root.Right))
	max = getMax(max, leftMax+rightMax+root.Val)
	return getMax(leftMax, rightMax) + root.Val
}

func getMax(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func TestAl(t *testing.T) {
	//tree1 := &TreeNode{Val: 15, Left: nil, Right: nil}
	//tree2 := &TreeNode{Val: 7, Left: nil, Right: nil}
	//tree3 := &TreeNode{Val: 20, Left: tree1, Right: tree2}
	//tree4 := &TreeNode{Val: 9, Left: nil, Right: nil}
	root := &TreeNode{Val: -1, Left: nil, Right: nil}
	fmt.Printf("%v", maxPathSum(root))
}
