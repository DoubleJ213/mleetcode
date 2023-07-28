package fxxk

import (
	"fmt"
	"testing"
)

/*
95. 不同的二叉搜索树 II
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同二叉搜索树。
可以按 任意顺序 返回答案。

示例 1：
输入：n = 3
输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
示例 2：

输入：n = 1
输出：[[1]]

提示：
1 <= n <= 8
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func generateTrees(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{buildTree95(1)}
	}
	//闭区间构造二叉树
	return traverse95(1, n)
}

func traverse95(left, right int) []*TreeNode {
	node95 := make([]*TreeNode, 0)
	if left > right {
		node95 = append(node95, nil)
		return node95
	}
	for i := left; i <= right; i++ {
		//构造左子树所有可能性
		leftList := traverse95(left, i-1)
		//构造右子树所有可能性
		rightList := traverse95(i+1, right)
		for _, lNode := range leftList {
			for _, rNode := range rightList {
				root := buildTree95(i)
				root.Left = lNode
				root.Right = rNode
				node95 = append(node95, root)
			}
		}
	}
	return node95
}

func buildTree95(n int) *TreeNode {
	return &TreeNode{n, nil, nil}
}

func TestAl95(t *testing.T) {
	a := generateTrees(3)
	fmt.Println(a)
}
