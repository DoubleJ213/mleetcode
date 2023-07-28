package fxxk

import (
	"fmt"
	"testing"
)

/*
98. 验证二叉搜索树
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效 二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

示例 1：
输入：root = [2,1,3]
输出：true
示例 2：
输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。

提示：
树中节点数目范围在[1, 104] 内
-2^31 <= Node.val <= 2^31 - 1
*/

func isValidBST(root *TreeNode) bool {
	//默认为true，当有一个判断不满足就直接为false
	return traverse98(root, nil, nil)
}

/*
本题要注意的两种
【2，2，2】 false 没人保证输入的根值不能重复
      5
  4      6
       3   7
*/
func traverse98(root *TreeNode, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}

	leftStatus := traverse98(root.Left, min, root)
	rightStatus := traverse98(root.Right, root, max)

	return leftStatus && rightStatus
}

func TestAl98(*testing.T) {
	root5 := &TreeNode{3, nil, nil}
	root4 := &TreeNode{7, nil, nil}
	root3 := &TreeNode{6, root5, root4}
	root2 := &TreeNode{4, nil, nil}
	root := &TreeNode{5, root2, root3}
	fmt.Println(isValidBST(root))
}
