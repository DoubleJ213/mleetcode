package fxxk

import (
	"fmt"
	"testing"
)

/*
889. 根据前序和后序遍历构造二叉树
给定两个整数数组，preorder和 postorder ，
其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，
postorder 是同一棵树的后序遍历，重构并返回二叉树。
如果存在多个答案，您可以返回其中 任何 一个。

示例 1：
输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
输出：[1,2,3,4,5,6,7]
示例 2:
输入: preorder = [1], postorder = [1]
输出: [1]

提示：
1 <= preorder.length <= 30
1 <= preorder[i] <= preorder.length
preorder中所有值都 不同
postorder.length == preorder.length
1 <= postorder[i] <= postorder.length
postorder中所有值都 不同
保证 preorder和 postorder是同一棵二叉树的前序遍历和后序遍历
*/
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		preorder[0], nil, nil,
	}
	if len(preorder) == 1 {
		return root
	}

	leftPre, leftPost, rightPre, rightPost := getNums(preorder, postorder)
	root.Left = constructFromPrePost(leftPre, leftPost)
	root.Right = constructFromPrePost(rightPre, rightPost)

	return root
}

// 输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
func getNums(preorder []int, postorder []int) (leftPre, leftPost, rightPre, rightPost []int) {
	//root := preorder[0]
	leftRoot := preorder[1]
	index := 0
	for i, post := range postorder {
		if post == leftRoot {
			index = i
			leftPre = preorder[1 : index+2]
			leftPost = postorder[:index+1]
			break
		}
	}
	rightRoot := postorder[len(postorder)-2]
	if leftRoot != rightRoot {
		for j, pre := range preorder {
			if pre == rightRoot {
				rightPre = preorder[j:]
				rightPost = postorder[index+1 : len(postorder)-1]
				break
			}
		}
	}

	return leftPre, leftPost, rightPre, rightPost
}

func TestAl889(t *testing.T) {
	a := constructFromPrePost([]int{2, 1}, []int{1, 2})
	fmt.Print(a)
}
