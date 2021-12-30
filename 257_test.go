package mleetcode

import (
	"fmt"
	"testing"
)

//257. 二叉树的所有路径
//给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//叶子节点 是指没有子节点的节点。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//var result := make([]string, 0)

var pathRes []string

func binaryTreePaths(root *TreeNode) []string {
	pathRes = make([]string, 0)
	var tmp string
	pathDfs(root, tmp)
	return pathRes
}

func pathDfs(root *TreeNode, tmp string) {
	if root == nil {
		return
	}
	if len(tmp) == 0 {
		tmp = fmt.Sprintf("%d", root.Val)
	} else {
		tmp = fmt.Sprintf("%s->%d", tmp, root.Val)
	}
	if root.Left == nil && root.Right == nil {
		pathRes = append(pathRes, tmp)
	} else {
		pathDfs(root.Left, tmp)
		pathDfs(root.Right, tmp)
	}
}

func TestAl257(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	fmt.Print(binaryTreePaths(root))
}
