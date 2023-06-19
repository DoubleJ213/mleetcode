package review

import (
	"fmt"
	"testing"
)

/*
104. 二叉树的最大深度
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明:叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度3 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestAl104(t *testing.T) {
	tree1 := &TreeNode{Val: 15, Left: nil, Right: nil}
	tree2 := &TreeNode{Val: 7, Left: nil, Right: nil}
	tree3 := &TreeNode{Val: 20, Left: tree1, Right: tree2}
	tree4 := &TreeNode{Val: 9, Left: nil, Right: nil}
	root := &TreeNode{Val: 3, Left: tree4, Right: tree3}
	root2 := &TreeNode{Val: 3, Left: nil, Right: nil}
	//fmt.Printf("%d", maxDepth(root))
	fmt.Printf("%d\n", maxDepth1(root))
	fmt.Printf("%d\n", maxDepth1(root2))
	fmt.Printf("%d\n", maxDepth1(nil))
}

var max int
var depth int

func maxDepth(root *TreeNode) int {
	max = 0
	depth = 0
	//traverse
	traverse(root)
	return max
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	depth++
	max = getMax(max, depth)
	traverse(root.Left)
	traverse(root.Right)
	depth--
	return
}

func getMax(m int, d int) int {
	if m >= d {
		return m
	}
	return d
}

/*
1、是否可以通过遍历一遍二叉树得到答案？如果可以，用一个 traverse 函数配合外部变量来实现。
2、是否可以定义一个递归函数，通过子问题（子树）的答案推导出原问题的答案？如果可以，写出这个递归函数的定义，并充分利用这个函数的返回值。
3、无论使用哪一种思维模式，你都要明白二叉树的每一个节点需要做什么，需要在什么时候（前中后序）做。
*/

//定义：输入根节点，返回这棵二叉树的最大深度
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 利用定义，计算左右子树的最大深度
	leftMax := maxDepth1(root.Left)
	rightMax := maxDepth1(root.Right)
	// 整棵树的最大深度等于左右子树的最大深度取最大值，
	// 然后再加上根节点自己
	//root2 := &TreeNode{Val: 3, Left: nil, Right: nil} 这个树深度1
	return getMax(leftMax, rightMax) + 1
}

/*
    3
   / \
  9  20
    /  \
   15   7
*/
