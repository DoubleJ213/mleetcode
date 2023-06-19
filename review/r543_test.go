package review

import (
	"fmt"
	"testing"
)

/*
543. 二叉树的直径

给你一棵二叉树的根节点，返回该树的 直径 。
二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
两节点之间路径的 长度 由它们之间边数表示。
    1
  2   3
4  5
示例 1：
输入：root = [1,2,3,4,5]
输出：3
解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
示例 2：
输入：root = [1,2]
输出：1
提示：
树中节点数目在范围 [1, 104] 内
-100 <= Node.val <= 100
*/

var diameter int

//104 按照labuladong思路 ，不要遍历，通过子树的结果推导 重写下
/*
    3
   / \
  9  20
    /  \
   15   7
*/
func diameterOfBinaryTree(root *TreeNode) int {
	diameter = 0
	abc := maxDepth2(root)
	fmt.Println(abc)
	return diameter
	// 如果写成这样就出错了,因为return的是深度，不是直接和
	// return maxDepth2(root)
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth2(root.Left)
	right := maxDepth2(root.Right)

	/*在后序遍历的时候，能看到完整的二叉树，
	所以如果此时进行回溯，计算当前节点的最大直径
	左子树的深度算好了，右子树的深度也算好了，还没有开始算自己的深度。
	所以还没到return +1的步骤，如果return了，那return的是该节点的深度，不是直径和
	我们是用深度算直径和的，直径和正好等于左子树的深度加上右子树的深度*/
	diameter = getMax(left+right, diameter)
	return getMax(left, right) + 1
}

func TestAl543(t *testing.T) {
	//tree1 := &TreeNode{Val: 15, Left: nil, Right: nil}
	//tree2 := &TreeNode{Val: 7, Left: nil, Right: nil}
	//tree3 := &TreeNode{Val: 20, Left: tree1, Right: tree2}
	tree5 := &TreeNode{Val: 3, Left: nil, Right: nil}
	tree4 := &TreeNode{Val: 2, Left: tree5, Right: nil}
	root := &TreeNode{Val: 1, Left: tree4, Right: nil}
	fmt.Printf("%d", diameterOfBinaryTree(root))
}
