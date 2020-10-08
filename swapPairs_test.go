package mleetcode

import (
	"testing"
)

//24. 两两交换链表中的节点

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func swapPairs(head *ListNode) *ListNode {
	head = &ListNode{
		Val:  0,
		Next: head,
	}
	prev := head
	for prev.Next != nil && prev.Next.Next != nil {
		//涉及4个部分变化， a1,a2,a3,a4
		a1, a2, a3, a4 := prev, prev.Next, prev.Next.Next, prev.Next.Next.Next
		//a1.Next = a3
		//a3.Next = a2
		//a2.Next = a4
		a1.Next, a2.Next, a3.Next = a3, a4, a2
		prev = prev.Next.Next
	}
	return head.Next
}

func TestSwapPairs(t *testing.T) {
	res := &ListNode{0, nil}
	temp := res
	for i := 1; i < 6; i++ {
		temp.Next = &ListNode{i, nil}
		temp = temp.Next
	}
	swapPairs(res.Next)
}
