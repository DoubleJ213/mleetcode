package review

import (
	"fmt"
	"testing"
)

/*

 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	//分解问题,拉平整个树，就是拉平左边，然后拉平右边
	flatten(root.Left)
	flatten(root.Right)
	//后序遍历的地方，可以看到树的全貌
	preRight := root.Right
	root.Left, root.Right = nil, root.Left

	//原先右子树，接到当前右子树最下面
	for root.Right != nil {
		root = root.Right
	}
	root.Right = preRight
}

func TestAl114(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	flatten(root)
	fmt.Println("a")
}
