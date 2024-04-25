package fxxk

import (
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

/*
首先，中序遍历，就是个递增的数组.
再有 用a 和 b 表示两个需要交换的节点。 额外使用last 记录上一个node。需要用last 和当前的值进行判断，
当 last > 当前root时 说明 这两个值不对。但是不一定就是这两个进行交换。
如果 a 为空 说明之前没有节点不正常.当前两个节点 last肯定是a root先暂时等于b 不一定是最终的b
如果a 不为空 说明之前有个不正常的节点，再和当前这个不正常的 偏小的  root 节点互换就可以了。
*/
var a99 *test.TreeNode
var b99 *test.TreeNode
var last99 *test.TreeNode

func recoverTree(root *test.TreeNode) {
	a99 = nil
	b99 = nil
	last99 = nil
	myRecover99(root)
	if a99 != nil && b99 != nil {
		a99.Val, b99.Val = b99.Val, a99.Val
	}
}

func myRecover99(root *test.TreeNode) {
	if root == nil {
		return
	}
	myRecover99(root.Left)
	if last99 != nil && last99.Val > root.Val {
		if a99 == nil {
			a99 = last99
			b99 = root
		} else {
			b99 = root
		}
	}
	last99 = root
	myRecover99(root.Right)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRecoverBinarySearchTree(t *testing.T) {
	root1 := &test.TreeNode{Val: 1}
	root2 := &test.TreeNode{Val: 2}
	root := &test.TreeNode{Val: 3, Left: root1, Right: root2}
	recoverTree(root)
}
