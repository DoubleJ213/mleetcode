package fxxk

import (
	"fmt"
	"testing"
)

/*
589. N 叉树的前序遍历
给定一个 n叉树的根节点 root，返回 其节点值的 前序遍历 。
n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

示例 1：
输入：root = [1,null,3,2,4,null,5,6]
输出：[1,3,5,6,2,4]
示例 2：
输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
输出：[1,2,3,6,7,11,14,4,8,12,5,9,13,10]

提示：
节点总数在范围[0, 104]内
0 <= Node.val <= 104
n 叉树的高度小于或等于 1000

进阶：递归法很简单，你可以使用迭代法完成此题吗?
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

var res589 []int

func preorder(root *Node) []int {
	res589 = make([]int, 0)
	if root == nil {
		return res589
	}
	toTraverseList := [][]*Node{{root}}
	for len(toTraverseList) > 0 {
		currentChild := toTraverseList[len(toTraverseList)-1]
		toTraverseList = toTraverseList[:len(toTraverseList)-1]
		for len(currentChild) > 0 {
			current := currentChild[0]
			res589 = append(res589, current.Val)
			otherChilds := currentChild[1:]
			if len(otherChilds) > 0 {
				toTraverseList = append(toTraverseList, otherChilds)
			}
			currentChild = current.Children
		}
	}
	return res589
}

//func preorder589(root *Node) []int {
//	res589 = make([]int, 0)
//	traverse589(root)
//	return res589
//}
//
//func traverse589(root *Node) {
//	if root == nil {
//		return
//	}
//	res589 = append(res589, root.Val)
//	for _, child := range root.Children {
//		traverse589(child)
//	}
//}

func TestAl589(t *testing.T) {
	root6 := &Node{6, nil}
	root5 := &Node{5, nil}
	root4 := &Node{4, nil}
	root3 := &Node{3, []*Node{root5, root6}}
	root2 := &Node{2, nil}
	_ = &Node{1, []*Node{root3, root2, root4}}
	fmt.Println(preorder(nil))
}
