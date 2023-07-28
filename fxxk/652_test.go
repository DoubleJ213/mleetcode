package fxxk

import (
	"fmt"
	"testing"
)

/*
652. 寻找重复的子树
给你一棵二叉树的根节点 root ，返回所有 重复的子树 。
对于同一类的重复子树，你只需要返回其中任意 一棵 的根结点即可。
如果两棵树具有 相同的结构 和 相同的结点值 ，则认为二者是 重复 的。

示例 1：
输入：root = [1,2,3,4,null,2,4,null,null,4]
输出：[[2,4],[4]]
示例 2：
输入：root = [2,1,1]
输出：[[1]]
示例 3：
输入：root = [2,2,2,3,null,3,null]
输出：[[2,3],[3]]

提示：
树中的结点数在 [1, 5000] 范围内。
-200 <= Node.val <= 200
*/
/*var allNodes []*TreeNode
var duplicateNodes []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	allNodes = make([]*TreeNode, 0)
	duplicateNodes = make([]*TreeNode, 0)
	traverse5(root)
	for i := 0; i < len(allNodes)-1; i++ {
		toAdd := false
		left := allNodes[i]
		if left != nil {
			for j := i + 1; j < len(allNodes); j++ {
				right := allNodes[j]
				if right != nil {
					if isSameTreeNode(left, right) {
						allNodes[j] = nil
						toAdd = true
					}
				}
			}
			if toAdd {
				duplicateNodes = append(duplicateNodes, left)
			}
		}
	}
	return duplicateNodes
}

func traverse5(root *TreeNode) {
	if root == nil {
		return
	}
	traverse5(root.Left)
	traverse5(root.Right)
	allNodes = append(allNodes, root)
}

func isSameTreeNode(root *TreeNode, node *TreeNode) bool {
	if root == nil {
		return node == nil
	}
	if node == nil {
		return false
	}
	if root.Val != node.Val {
		return false
	}
	left := isSameTreeNode(root.Left, node.Left)
	right := isSameTreeNode(root.Right, node.Right)
	return left == true && right == true
}

//写在最后，这个代码写完信心满满提交，超时了
//我还在遍历的时候做了剪枝还超时，看来得优化了
*/

func TestAl652(t *testing.T) {
	node6 := &TreeNode{Val: 0}
	node5 := &TreeNode{Val: 0, Left: node6}
	node4 := &TreeNode{Val: 0}
	node3 := &TreeNode{Val: 0, Left: nil, Right: node5}
	node2 := &TreeNode{Val: 0, Left: node4, Right: nil}
	root := &TreeNode{Val: 0, Left: node2, Right: node3}

	fmt.Println(findDuplicateSubtrees(root))
	fmt.Println("abc")

}

var nodeMap map[string]int
var duplicateNodes []*TreeNode

/*
参考之前序列化的思路，将每个子树进行序列化
然后key为序列化后的值
遍历map，value次数大于1的，反序列化？
这个方法，多了序列化反序列化的过程，少了全部子树去重过程

写的过程中又发现，其实不需要反序列化，只需要序列化后和map已有元素进行比较
当value等1的时候将当前的root加到最终返回的数组就行
*/
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	nodeMap = make(map[string]int)
	duplicateNodes = make([]*TreeNode, 0)
	traverse5(root)
	return duplicateNodes
}

func traverse5(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	leftStr := traverse5(root.Left)
	rightStr := traverse5(root.Right)
	//rootStr := fmt.Sprintf("%s,%d,%s", leftStr, root.Val, rightStr)
	//这里后序遍历，root在最后
	rootStr := fmt.Sprintf("%s,%s,%d", leftStr, rightStr, root.Val)
	if times, ok := nodeMap[rootStr]; ok {
		nodeMap[rootStr] = times + 1
		if times == 1 {
			duplicateNodes = append(duplicateNodes, root)
		}
	} else {
		nodeMap[rootStr] = 1
	}
	return rootStr
}
