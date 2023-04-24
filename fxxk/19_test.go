package fxxk

import (
	"fmt"
	"testing"
)

func TestE(t *testing.T) {
	node := &ListNode{
		Val:  1,
		Next: nil,
	}
	tmp := node
	// node 赋值给tmp，如果node改变
	// 如果tmp改变，node也改变
	for i := 2; i <= 3; i++ {
		tmp.Next = &ListNode{Val: i, Next: nil}
		//tmp改成其他值，node之前的指针不变
		tmp = tmp.Next
	}
	removeNthFromEnd(node, 3)
	fmt.Println(node)
}

//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	leftNode := head
	rightNode := head

	for i := 0; i < n; i++ {
		if rightNode.Next == nil {
			//n超出head长度
			return head.Next
		}
		rightNode = rightNode.Next
	}
	for rightNode.Next != nil {
		rightNode = rightNode.Next
		leftNode = leftNode.Next
	}
	// 这步是关键 修改left的指针，改变head
	leftNode.Next = leftNode.Next.Next
	return head
}
