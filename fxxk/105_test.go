package fxxk

import (
	"fmt"
	"testing"
)

//105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 递归
	if len(preorder) == 0 || len(inorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	//找到当前两个数组的 root节点，分成左右子树，再次递归
	root := &TreeNode{preorder[0], nil, nil}
	index := 0
	for index = 0; index < len(inorder); index++ {
		if inorder[index] == preorder[0] {
			break
		}
	}
	// 中序遍历左子树为 最开始到index，不包含index 所以为 inorder[:index]
	// 前序遍历左子树 应该和上面等长，从哪里开始，第一位不算 [1:len(inorder[:index])+1]
	root.Left = buildTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	// 中序遍历右子树为 index往后（不包含index）到最后
	// 前序遍历右子树哪里到哪里，简单想就是头为root 左子树剩余的部分 就是右子树
	root.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}

func TestAl105(*testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	aa := buildTree(preorder, inorder)
	fmt.Println(aa)
	fmt.Println("done")
}
