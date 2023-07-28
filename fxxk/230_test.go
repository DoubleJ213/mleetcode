package fxxk

import (
	"fmt"
	"testing"
)

/*
230. 二叉搜索树中第K小的元素
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第k个最小元素（从 1 开始计数）。

示例 1：
输入：root = [3,1,4,null,2], k = 1
输出：1
示例 2：
输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3

提示：
树中的节点数为 n 。
1 <= k <= n <= 104
0 <= Node.val <= 104

进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？
//*/

/*
再来个结构体，记录以自己为根时所有节点总数，左小右大。
然后通过遍历左子树，知道当前节点左子树节点总数m
当个数m等于k，返回
当个数m大于k，左子树中继续找
当个数m小于k，右子树中继续找
只是满足进阶要求，就一次查询的情况下，貌似没有优势，这里先不写了，下次可以按照这个思路写
转到538题，用上面思路写
*/

var res230 []int

/*
二叉搜索树特性要牢记，根的值小于右子树的值大于左子树的值
考虑到上面这个特性，二叉搜索树前序遍历完，第k小的值，就是第k个元素
*/
func kthSmallest(root *TreeNode, k int) int {
	res230 = make([]int, 0)
	traverse230(root)
	return res230[k-1]
}

func traverse230(root *TreeNode) {
	if root == nil {
		return
	}
	traverse230(root.Left)
	res230 = append(res230, root.Val)
	traverse230(root.Right)
}

func TestAl230(t *testing.T) {
	root6 := &TreeNode{1, nil, nil}
	root5 := &TreeNode{4, nil, nil}
	root4 := &TreeNode{2, root6, nil}
	root3 := &TreeNode{6, nil, nil}
	root2 := &TreeNode{3, root4, root5}
	root := &TreeNode{5, root2, root3}
	fmt.Println(kthSmallest(root, 2))
	//fmt.Println(kthSmallest1(root, 2))
}

/*func kthSmallest1(root *TreeNode, k int) int {
	res230 = make([]int, 0)
	traverse230_1(root, k)
	return res230[k-1]
}

func traverse230_1(root *TreeNode, k int) {
	if root == nil || len(res230) >= k {
		return
	}
	traverse230_1(root.Left, k)
	res230 = append(res230, root.Val)
	traverse230_1(root.Right, k)
}*/
