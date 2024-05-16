package fxxk

import (
	"fmt"
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var ans129 int

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return ans129
	}
	ans129 = 0
	traverse129(root, 0)
	return ans129
}

func traverse129(root *TreeNode, sum int) *TreeNode {
	if root == nil {
		return root
	}
	left := traverse129(root.Left, sum*10+root.Val)
	right := traverse129(root.Right, sum*10+root.Val)
	if left == nil && right == nil {
		ans129 += sum*10 + root.Val
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSumRootToLeafNumbers(t *testing.T) {
	//root7 := &TreeNode{7, nil, nil}
	//root6 := &TreeNode{5, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{1, nil, nil}
	//root3 := &TreeNode{3, root6, root7}
	root3 := &TreeNode{0, nil, nil}
	root2 := &TreeNode{9, root4, root5}
	root := &TreeNode{4, root2, root3}
	fmt.Println(sumNumbers(root)) //522

}
