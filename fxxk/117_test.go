package fxxk

import (
	"fmt"
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

var preNode []*Node

// 给定二叉树  填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。
func connect177(root *Node) *Node {
	//进阶： 你只能使用常量级额外空间。用递归解题也符合要求，本题中递归程序的隐式栈空间不计入额外空间复杂度。
	// o(n)的好写，o(1)怎么来?
	//先写on 吧
	// 首先，同一个root上面的左右子树互相连接比较简单。难点就是不同父节点，左右相连
	preNode = make([]*Node, 0)
	//preNode 存层数为index的前一个node节点
	myCon177(root, 0)
	return root
}

func myCon177(root *Node, level int) {
	if root == nil {
		return
	}

	if len(preNode) > level {
		pre := preNode[level]
		if pre != nil {
			pre.Next = root
		}
		preNode[level] = root
	} else {
		preNode = append(preNode, root)
	}
	//
	//pre := root.Left
	//if pre != nil {
	//	// 同root上面的左右子树互相连接
	//	pre.Next = root.Right
	//}

	myCon177(root.Left, level+1)
	myCon177(root.Right, level+1)

}

func connect2(root *Node) *Node {
	ans := root
	myDfs177(root)
	return ans
}

func myDfs177(root *Node) {
	if root == nil {
		return
	}
	var first *Node
	var last *Node
	for root != nil {
		left := root.Left
		right := root.Right
		if left == nil && right == nil {
			root = root.Next
			continue
		}
		if first == nil {
			//	找下一层头节点
			if left == nil {
				//两个都nil已经跳过了，必然至少存在一个
				first = right
			} else {
				first = left
			}
		}
		if last == nil {
			//last用来连接下一层链表,顺便表示尾结点
			last = first
		}
		//到这里last已经存在了
		if last == left {
			last.Next = right
			last = last.Next
		} else if last == right {
			root = root.Next
			continue
		} else {
			if left != nil {
				last.Next = left
				last = last.Next
			}
			if right != nil {
				last.Next = right
				last = last.Next
			}
		}
		root = root.Next
	}
	myDfs177(first)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPopulatingNextRightPointersInEachNodeIi(t *testing.T) {
	root7 := &Node{7, nil, nil, nil}
	root6 := &Node{6, nil, nil, nil}
	root5 := &Node{5, nil, nil, nil}
	root4 := &Node{4, nil, nil, nil}
	root3 := &Node{3, root6, root7, nil}
	root2 := &Node{2, root4, root5, nil}
	root := &Node{1, root2, root3, nil}
	fmt.Println(connect177(root))

}
