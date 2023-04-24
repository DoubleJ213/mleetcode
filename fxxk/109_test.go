package fxxk

import (
	"fmt"

	"testing"
)

//109. 有序链表转换二叉搜索树
//给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
//本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

func sortedListToBST(head *ListNode) *TreeNode {
	//找出链表中位数节点
	//在找出了中位数节点之后，我们将其作为当前根节点的元素。
	//并递归地构造其左侧部分的链表对应的左子树，以及右侧部分的链表对应的右子树。
	root := buildMyTree(head, nil)
	return root
}

func buildMyTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getMiddle(left, right)
	root := &TreeNode{mid.Val, nil, nil}
	root.Left = buildMyTree(left, mid)
	root.Right = buildMyTree(mid.Next, right)
	return root
}

func getMiddle(left, right *ListNode) *ListNode {
	if left == nil || left.Next == nil {
		return left
	}
	fast, slow := left, left
	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func TestAl109(*testing.T) {
	node := &ListNode{
		Val:  -1,
		Next: nil,
	}
	tmp := node
	for i := 0; i < 1; i++ {
		tmp.Next = &ListNode{Val: i, Next: nil}
		tmp = tmp.Next
	}
	sortedListToBST(node.Next)
	fmt.Println("done")
}
