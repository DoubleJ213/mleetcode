package fxxk

import (
	"fmt"
	"testing"
)

/*
437. 路径总和 III
给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

示例 1：
输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
输出：3
解释：和等于 8 的路径有 3 条，如图所示。
示例 2：

输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：3

提示:
二叉树的节点个数的范围是 [0,1000]
-10^9 <= Node.val <= 10^9
-1000 <= targetSum <= 1000
*/

/*
路径和的数目，知道题目的一些要求，路径和起点不一定是根，也不需要在叶子节点结束，但是方向必须是向下的，同一层的不能算和
另外没提示的，就是路径和肯定路径是某一段连续的，不能跳着求和
这题如果是个数组，然后求targetSum,怎么算？这种题目一般都是dp。
比如[1,2,3,4]
1
12 2
123 23 3
1234 234 34 4 有这么多种路径，然后求和，判等。
这个回头可以想想前缀和的解法能不能有什么优化

二叉树呢
    1
  2   3
 4 5 6 7

1 、[2/3] 、[4\5\6\7] 还不准确

1 2 4
1 2 5
1 3 6
1 3 7
所有从根开始的路径先算出来？
然后和上面类似？
因为只要算个数，所以每次判断和就行，不用记录路径有哪些
*/
var ans437 int

func pathSum(root *TreeNode, targetSum int) int {
	ans437 = 0
	targets := make([]int, 0)
	if root != nil {
		dfs437(root, targetSum, targets)
	}
	return ans437
}

func dfs437(root *TreeNode, targetSum int, targets []int) {
	if root == nil {
		return
	}
	newTargets := make([]int, 0)
	newTargets = append(newTargets, root.Val)

	for i := 0; i < len(targets); i++ {
		total := targets[i] + root.Val
		if total == targetSum {
			ans437++
		}
		newTargets = append(newTargets, total)
	}

	dfs437(root.Left, targetSum, newTargets)
	dfs437(root.Right, targetSum, newTargets)
}

func Test437(t *testing.T) {
	root9 := &TreeNode{9, nil, nil}
	//root6 := &TreeNode{6, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root0 := &TreeNode{0, nil, nil}
	root8 := &TreeNode{8, nil, root9}
	root2 := &TreeNode{2, root0, root4}
	root := &TreeNode{6, root2, root8}
	fmt.Println(pathSum(root, 8))
}
