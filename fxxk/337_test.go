package fxxk

import (
	"fmt"
	"testing"
)

/*
337. 打家劫舍 III
小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
除了 root 之外，每栋房子有且只有一个“父“房子与之相连。
一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额

输入: root = [3,2,3,null,3,null,1]
输出: 7
解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7

	    3
	2       3
	  3       1

输入: root = [3,4,5,1,3,null,1]
输出: 9
解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9

	      3
	  4       5
	1   3       1

提示：
树的节点数在 [1, 10^4] 范围内
0 <= Node.val <= 10^4
*/
var dp337 map[*TreeNode]int

func rob337(root *TreeNode) int {
	dp337 = make(map[*TreeNode]int)
	// 二叉树dp之前还没有接触过, 之前都遍历的数组现在遍历二叉树，要么递归，要么for循环这点倒是没啥疑问。
	// 但是之前直接初始化一个数组，然后一一对应的能递归结果，现在怎么存这个结果，然后递推，难不成再初始化一个树，用来存结果？
	// 好在这题只是需要知道最大的值int，left 和right 的dp值相加 和 root的dp值比较大小
	// 不过这样有点复杂,先不写了。
	// 直接用个map存一下 key 为指针， value为对应用来递推的值
	return dfs337(root)
}

func dfs337(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if val, ok := dp337[root]; ok {
		return val
	}

	money := root.Val

	if root.Left != nil {
		money += dfs337(root.Left.Left) + dfs337(root.Left.Right)
	}

	if root.Right != nil {
		money += dfs337(root.Right.Left) + dfs337(root.Right.Right)
	}
	ans := getMax(money, dfs337(root.Left)+dfs337(root.Right))
	dp337[root] = ans
	return ans
}

func TestAl337(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	fmt.Println(rob337(root))
}
