package fxxk

import (
	"fmt"
	"github.com/mlcPractice/leetcode/editor/cn"
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
var res199 []int

func rightSideView(root *test.TreeNode) []int {
	res199 = make([]int, 0)
	if root == nil {
		return res199
	}
	traverse199(root, 0)
	return res199
}

func traverse199(root *test.TreeNode, level int) {
	if root == nil {
		return
	}

	if len(res199)-1 < level {
		res199 = append(res199, root.Val)
	} else {
		res199[level] = root.Val
	}

	traverse199(root.Left, level+1)
	traverse199(root.Right, level+1)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeRightSideView(t *testing.T) {
	//dfs 一直right？
	var root = &test.TreeNode{1,
		&test.TreeNode{2, test.root4, test.root5},
		&test.TreeNode{3, nil, nil}}
	fmt.Println(rightSideView(root))
}
