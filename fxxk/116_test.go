package fxxk

import (
	"fmt"
	"testing"
)

/*
116. 填充每个节点的下一个右侧节点指针
给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。
如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有next 指针都被设置为 NULL。

示例 1：
输入：root = [1,2,3,4,5,6,7]
输出：[1,#,2,3,#,4,5,6,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，
以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，
同一层节点由 next 指针连接，'#' 标志着每一层的结束。
示例 2:
输入：root = []
输出：[]

提示：
树中节点的数量在[0, 2^12- 1]范围内
-1000 <= node.val <= 1000
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	//待刷
	return root
}

func TestAl116(t *testing.T) {
	root7 := &Node{7, nil, nil, nil}
	root6 := &Node{6, nil, nil, nil}
	root5 := &Node{5, nil, nil, nil}
	root4 := &Node{4, nil, nil, nil}
	root3 := &Node{3, root6, root7, nil}
	root2 := &Node{2, root4, root5, nil}
	root := &Node{1, root2, root3, nil}
	fmt.Printf("%v", connect(root))
	fmt.Println("1")
}
