package fxxk

import (
	"fmt"

	"testing"
)

//114. 二叉树展开为链表

/*
给你二叉树的根结点root，请你将它展开为一个单链表：
展开后的单链表应该同样使用TreeNode,
其中right子指针指向链表中下一个结点，而左子指针始终为null。
展开后的单链表应该与二叉树先序遍历顺序相同。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//我提问了，看看有没有人回复

func flatten2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	flattenDfs(root, &res)
	tree := &TreeNode{Val: -1, Left: nil, Right: nil}
	node := tree
	for _, value := range res {
		node.Right = &TreeNode{Val: value, Left: nil, Right: nil}
		node.Left = nil
		node = node.Right
	}
	return tree.Right
}

func flattenDfs(root *TreeNode, i *[]int) {
	if root == nil {
		return
	}
	*i = append(*i, root.Val)
	flattenDfs(root.Left, i)
	flattenDfs(root.Right, i)
}

func flatten(root *TreeNode) {
	list := preorderTraversal4(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
	fmt.Println("aaa")
}

func preorderTraversal4(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal4(root.Left)...)
		list = append(list, preorderTraversal4(root.Right)...)
	}
	return list
}

func TestAl114(*testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	//aa := testPreOrder(root)
	//aa := flatten2(root)
	//fmt.Printf("114 %v", aa)
	//flatten(root)
	//flatten3(root)
	preorderTraversal3(root)
}

func flatten3(root *TreeNode) *TreeNode {
	head := root
	var max *TreeNode

	for root != nil {
		if root.Left == nil {
			root = root.Right
		} else {
			max = root.Left
			for max.Right != nil {
				max = max.Right
			}

			root.Right, max.Right = root.Left, root.Right
			root.Left = nil
		}
	}
	return head
}

//先序遍历是将左子树遍历完成后才会遍历右子树，
//而左子树最后遍历的位置其实是左子树的最右端的那个节点，
//因此将右子树挂到左子树的最右端的节点下面就ok了。递归重复此过程。

func flatten3Dfs(root *TreeNode) {
	//https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by--26/
	if root == nil {
		return
	}
	flatten3Dfs(root.Right)
	flatten3Dfs(root.Left)

}

func testPreOrder(root *TreeNode) []int {
	res := make([]int, 0)
	//testDfs(root, &res)
	fmt.Println(res)
	testZhi(res)
	fmt.Println(res)
	testYinyong(&res)
	fmt.Println(res)
	return res
}

func testYinyong(i *[]int) {
	//切片的指针传过来，可以修改到原来的切片的内容
	*i = append(*i, 2)
}

func testZhi(i []int) {
	//切片是个引用类型的作为参数传过来是当前切片的副本，能取值，不能赋值
	i = append(i, 1)
}

func testDfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	testDfs(root.Left, res)
	testDfs(root.Right, res)
}

func preorderTraversal3(root *TreeNode) []int {
	/*
		Morris 遍历无需使用栈，空间复杂度为 O(1)。核心思想是：
		遍历二叉树节点，
		若当前节点 root 的左子树为空，将当前节点值添加至结果列表 res 中，并将当前节点更新为 root.right
		若当前节点 root 的左子树不为空，找到左子树的最右节点 pre（也即是 root 节点在中序遍历下的前驱节点）：
		若前驱节点 pre 的右子树为空，将当前节点值添加至结果列表 res 中，然后将前驱节点的右子树指向当前节点 root，并将当前节点更新为 root.left。
		若前驱节点 pre 的右子树不为空，将前驱节点右子树指向空（即解除 pre 与 root 的指向关系），并将当前节点更新为 root.right。
		循环以上步骤，直至二叉树节点为空，遍历结束
	*/
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
