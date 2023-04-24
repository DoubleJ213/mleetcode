package fxxk

import (
	"testing"
)

//24. 两两交换链表中的节点

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 24 两两翻转
func swapPairs(head *ListNode) *ListNode {
	head = &ListNode{
		Val:  0,
		Next: head,
	}
	prev := head
	for prev.Next != nil && prev.Next.Next != nil {
		//涉及4个部分变化， a1,a2,a3,a4
		a1, a2, a3, a4 := prev, prev.Next, prev.Next.Next, prev.Next.Next.Next
		a1.Next, a2.Next, a3.Next = a3, a4, a2
		prev = prev.Next.Next
	}
	return head.Next
}

func swapPairs1(head *ListNode) *ListNode {
	tmp := &ListNode{
		Val:  -1,
		Next: head,
	}
	res := tmp
	for tmp.Next != nil && tmp.Next.Next != nil {
		pre1 := tmp
		pre := tmp.Next
		pos := tmp.Next.Next
		pos1 := tmp.Next.Next.Next
		pre1.Next, pos.Next, pre.Next = pos, pre, pos1

		tmp = tmp.Next.Next

	}

	return res.Next
}

func TestSwapPairs(t *testing.T) {
	res := &ListNode{0, nil}
	temp := res
	for i := 1; i < 6; i++ {
		temp.Next = &ListNode{i, nil}
		temp = temp.Next
	}
	//swapPairs(res.Next)
	swapPairs1(res.Next)
}
