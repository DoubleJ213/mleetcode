package fxxk

import "testing"

//Morris

func TestPreMorris(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	preorderMorris(root)
}

func preorderMorris(root *TreeNode) []int {
	var res []int
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			pre := root.Left
			for pre.Right != nil && pre.Right != root {
				pre = pre.Right
			}
			if pre.Right == nil {
				res = append(res, root.Val)
				pre.Right = root
				root = root.Left
			} else {
				pre.Right = nil
				root = root.Right
			}
		}
	}
	return res
}
