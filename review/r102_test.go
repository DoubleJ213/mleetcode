package review

import (
	"fmt"
	"testing"
)

func TestAl102(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	//fmt.Println(levelOrder(root))
	fmt.Println(levelOrder2(root))
}

/*var res [][]int

func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)
	traverse1(root, 0)
	return res
}

func traverse1(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if len(res) <= level {
		tmp := make([]int, 0)
		tmp = append(tmp, root.Val)
		res = append(res, tmp)
	} else {
		tmp := res[level]
		tmp = append(tmp, root.Val)
		res[level] = tmp
	}

	//可以优化成下面的代码。长度不足，没有这个元素，可以先补一个空节点。
	//然后对层数对应那个元素进行追加
	//if len(res) <= level {
	//	tmp := make([]int, 0)
	//	res = append(res, tmp)
	//}
	//res[level] = append(res[level], root.Val)

	traverse1(root.Left, level+1)
	traverse1(root.Right, level+1)
	//if res[level] == nil {
	//	lrs := make([]int, 0)
	//	lrs = append(lrs, root.Val)
	//	res[level] = lrs
	//} else {
	//	res[level] = append(res[level], root.Val)
	//}
}*/

func levelOrder2(root *TreeNode) [][]int {
	//	迭代
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	listNode := make([]*TreeNode, 0)
	listNode = append(listNode, root)
	for len(listNode) > 0 {
		size := len(listNode)
		toAddNode := make([]*TreeNode, 0)
		currentRes := make([]int, 0)
		for i := 0; i < size; i++ {
			tmpNode := listNode[i]
			currentRes = append(currentRes, tmpNode.Val)
			if tmpNode.Left != nil {
				toAddNode = append(toAddNode, tmpNode.Left)
			}
			if tmpNode.Right != nil {
				toAddNode = append(toAddNode, tmpNode.Right)
			}
		}
		listNode = toAddNode
		res = append(res, currentRes)
	}
	return res
}
