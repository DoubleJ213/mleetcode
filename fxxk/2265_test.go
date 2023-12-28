package fxxk

import (
	"fmt"
	"testing"
)

/*
2265. 统计值等于子树平均值的节点数
给你一棵二叉树的根节点 root ，找出并返回满足要求的节点数，要求节点的值等于其 子树 中值的 平均值 。
注意：
n 个元素的平均值可以由 n 个元素 求和 然后再除以 n ，并 向下舍入 到最近的整数。
root 的 子树 由 root 和它的所有后代组成。

输入：root = [4,8,5,0,1,null,6]
输出：5
解释：
对值为 4 的节点：子树的平均值 (4 + 8 + 5 + 0 + 1 + 6) / 6 = 24 / 6 = 4 。
对值为 5 的节点：子树的平均值 (5 + 6) / 2 = 11 / 2 = 5 。
对值为 0 的节点：子树的平均值 0 / 1 = 0 。
对值为 1 的节点：子树的平均值 1 / 1 = 1 。
对值为 6 的节点：子树的平均值 6 / 1 = 6 。

输入：root = [1]
输出：1
解释：对值为 1 的节点：子树的平均值 1 / 1 = 1。

提示：
树中节点数目在范围 [1, 1000] 内
0 <= Node.val <= 1000
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//统计多少个 节点值等于子树和自己和的平均值的
var total2265 int

func averageOfSubtree(root *TreeNode) int {
	total2265 = 0
	if root != nil {
		dfs2265(root)
		//dfs2265(root, 0 子树有数的个数, 0 子树所有的值和)
	}
	return total2265
}

/*
dfs2265
return 的值 count 子树有数的个数  sum 子树所有的值和
*/
func dfs2265(root *TreeNode) (count, sum int) {
	//二叉树
	if root == nil {
		return 0, 0
	}
	//, count+1, sum+root.Val
	//root.Val
	leftCount, leftSum := dfs2265(root.Left)
	rightCount, rightSum := dfs2265(root.Right)

	if isEqual(root.Val, leftCount+rightCount+1, root.Val+leftSum+rightSum) {
		total2265++
	}
	return leftCount + rightCount + 1, root.Val + leftSum + rightSum
}

func isEqual(val, count, sum int) bool {
	if val == sum/count {
		return true
	}
	return false
}

func TestAl2265(t *testing.T) {
	root7 := &TreeNode{6, nil, nil}
	//root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{1, nil, nil}
	root4 := &TreeNode{0, nil, nil}
	root3 := &TreeNode{5, nil, root7}
	root2 := &TreeNode{8, root4, root5}
	root := &TreeNode{4, root2, root3}
	fmt.Println(averageOfSubtree(root))
}
