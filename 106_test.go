package mleetcode

import (
	"fmt"
	"testing"
)

//106. 从中序与后序遍历序列构造二叉树
//注意:你可以假设树中没有重复的元素。
func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 || len(inorder) != len(postorder) {
		return nil
	}
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	var index int
	for index = 0; index < len(inorder); index++ {
		if postorder[len(postorder)-1] == inorder[index] {
			break
		}
	}
	//后序最后一个值为root
	//通过root找到中序遍历的左右子树
	//后序每次取一样的长度，且每次找根节点即可
	root.Left = buildTree106(inorder[:index], postorder[:index])
	root.Right = buildTree106(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}

func TestAl106(*testing.T) {
	//中序遍历 inorder = s[9,3,15,20,7]
	//后序遍历 postorder = [9,15,7,20,3]
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	aaa := buildTree106(inorder, postorder)
	fmt.Println(aaa)
}
