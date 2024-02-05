package fxxk

import (
	"fmt"
	"testing"
)

/*
1457. 二叉树中的伪回文路径
给你一棵二叉树，每个节点的值为 1 到 9 。我们称二叉树中的一条路径是 「伪回文」的，当它满足：
路径经过的所有节点值的排列中，存在一个回文序列。
请你返回从根到叶子节点的所有路径中 伪回文 路径的数目。

示例 1：
输入：root = [2,3,1,3,1,null,1]
        2
    3        1
  3   1   nil  1
输出：2
解释：上图为给定的二叉树。总共有 3 条从根到叶子的路径：红色路径 [2,3,3] ，绿色路径 [2,1,1] 和路径 [2,3,1] 。
     在这些路径中，只有红色和绿色的路径是伪回文路径，因为红色路径 [2,3,3] 存在回文排列 [3,2,3] ，绿色路径 [2,1,1] 存在回文排列 [1,2,1] 。
示例 2：
输入：root = [2,1,1,1,3,null,null,null,null,null,1]
输出：1
解释：上图为给定二叉树。总共有 3 条从根到叶子的路径：绿色路径 [2,1,1] ，路径 [2,1,3,1] 和路径 [2,1] 。
     这些路径中只有绿色路径是伪回文路径，因为 [2,1,1] 存在回文排列 [1,2,1] 。
示例 3：
输入：root = [9]
输出：1

提示：
给定二叉树的节点数目在范围 [1, 105] 内
1 <= Node.val <= 9
*/

var nums1457 [][]int

func pseudoPalindromicPaths(root *TreeNode) int {
	nums1457 = make([][]int, 0)
	ans := 0
	dfs1457(root, make([]int, 0))
	for i := 0; i < len(nums1457); i++ {
		nums := nums1457[i]
		//长度为偶数，即出现次数为奇数的字符个数为 0 个, 长度为奇数，即出现次数为奇数的字符个数为 1 个（位于中间）
		//它最多只有一个字符的数量是奇数，其余字符数量一定是偶数，不满足条件的就返回false
		if isTrue1457(nums) {
			ans++
		}
	}
	return ans
}

func isTrue1457(nums []int) bool {
	hTable := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if value, ok := hTable[num]; ok {
			hTable[num] = value + 1
		} else {
			hTable[num] = 1
		}
	}

	res := 0
	for x := 1; x <= 9; x++ {
		if value, ok := hTable[x]; ok {
			if value%2 != 0 {
				res++
			}
		}
	}
	if res > 1 {
		return false
	}
	return true
}

func dfs1457(root *TreeNode, path []int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		tmp := make([]int, len(path))
		copy(tmp, path)
		tmp = append(tmp, root.Val)
		nums1457 = append(nums1457, tmp)
		return
	}
	if root.Left != nil {
		dfs1457(root.Left, append(path, root.Val))
	}
	if root.Right != nil {
		dfs1457(root.Right, append(path, root.Val))
	}
}

func TestAl1457(t *testing.T) {
	root7 := &TreeNode{7, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{3, root6, root7}
	root2 := &TreeNode{2, root4, root5}
	root := &TreeNode{1, root2, root3}
	fmt.Println(pseudoPalindromicPaths2(root))
}

/*
int类型的变量 cnt 来统计各数值的出现次数奇偶性：
若cnt的第k位为1，说明数值k的出现次数为奇数，否则说明数值k出现次数为偶数或没出现过。
cnt=(0001010)2  代表数值 1 和数值 3 出现次数为奇数次，其余数值没出现过或出现次数为偶数次。
改变一个数 对应位数的值  cnt ^= 1 << k
判断是否最多只有一个字符出现奇数次的操作，也就是判断一个二进制数字是为全为 0 或仅有一位 1，
可配合 lowbit 来做，若 cnt 与 lowbit(cnt) = cnt & -cnt 相等，说明满足要求。
lowbit(x) 表示 x 的二进制表示中最低位的 1 所在的位置对应的值，即仅保留从最低位起的第一个 1，其余位均以 0 填充：
x = 6，其二进制表示为 (110)2 ，那么 lowbit(6)=(010)2=2
x = 12，其二进制表示为 (1100)2 ，那么 lowbit(12)=(100)2=4
*/
var ans1457 int

func pseudoPalindromicPaths2(root *TreeNode) int {
	ans1457 = 0
	dfs1457Pp(root, 0)
	return ans1457
}

func dfs1457Pp(root *TreeNode, cnt int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		cnt ^= 1 << root.Val //本层的值翻转
		//	翻转完后开始判断
		if cnt == cnt&-cnt {
			ans1457++
		}
	}

	if root.Left != nil {
		dfs1457Pp(root.Left, cnt^(1<<root.Val))
	}
	if root.Right != nil {
		dfs1457Pp(root.Right, cnt^(1<<root.Val))
	}
}
