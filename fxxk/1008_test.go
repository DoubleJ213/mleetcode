package fxxk

import (
	"fmt"
	"testing"
)

/*
1008. 前序遍历构造二叉搜索树
给定一个整数数组，它表示BST(即 二叉搜索树 )的 先序遍历 ，构造树并返回其根。
保证 对于给定的测试用例，总是有可能找到具有给定需求的二叉搜索树。
二叉搜索树 是一棵二叉树，其中每个节点，
Node.left的任何后代的值 严格小于 Node.val,
Node.right的任何后代的值 严格大于 Node.val。
二叉树的 前序遍历 首先显示节点的值，然后遍历Node.left，最后遍历Node.right。

示例 1：
输入：preorder = [8,5,1,7,10,12]
输出：[8,5,10,1,7,null,12]
示例 2:
输入: preorder = [1,3]
输出: [1,null,3]

提示：
1 <= preorder.length <= 100
1 <= preorder[i]<= 10^8
preorder 中的值 互不相同
*/
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}
	left, right := getPreNums(preorder)
	leftRoot := bstFromPreorder(left)
	rightRoot := bstFromPreorder(right)
	return &TreeNode{preorder[0], leftRoot, rightRoot}
}

func getPreNums(preorder []int) (left, right []int) {
	root := preorder[0]
	if preorder[len(preorder)-1] < root {
		//	最后一个数都比root小，树都在左边
		left = preorder[1:]
	} else {
		for i, num := range preorder {
			//单边的树有点小问题。不存在比root大的怎么办
			if num > root {
				left = preorder[1:i]
				right = preorder[i:]
				break
			}
		}
	}
	return left, right
}

func TestAl1008(t *testing.T) {
	a := bstFromPreorder([]int{4, 2})
	fmt.Println(a)

}
