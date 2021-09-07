package mleetcode

import (
	"fmt"
	"testing"
)

//876. 链表的中间结点
//给定一个头结点为 head 的非空单链表，返回链表的中间结点。
//如果有两个中间结点，则返回第二个中间结点。

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func TestAl876(*testing.T) {
	node := &ListNode{
		Val:  -1,
		Next: nil,
	}
	tmp := node
	for i := 1; i < 6; i++ {
		tmp.Next = &ListNode{Val: i, Next: nil}
		tmp = tmp.Next
	}
	middleNode(node.Next)
	fmt.Println("done")
}
