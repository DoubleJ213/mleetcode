package fxxk

import "testing"

//给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，
//使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//你应当 保留 两个分区中每个节点的初始相对位置。

//示例 1：
//
//输入：head = [1,4,3,2,5,2], x = 3
//输出：[1,2,2,4,3,5]
//示例 2：
//
//输入：head = [2,1], x = 2
//输出：[1,2]
//

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func partition(head *ListNode, x int) *ListNode {
	dummyA := &ListNode{-1, nil}
	dummyB := &ListNode{-1, nil}
	tmpA := dummyA
	tmpB := dummyB
	for head != nil {
		if head.Val < x {
			dummyA.Next = head
			head = head.Next
			dummyA.Next.Next = nil
			dummyA = dummyA.Next
		} else {
			dummyB.Next = head
			head = head.Next
			dummyB.Next.Next = nil
			dummyB = dummyB.Next
		}
	}
	dummyA.Next = tmpB.Next
	return tmpA.Next
}

func TestAl86(t *testing.T) {
	list1D := &ListNode{Val: 5, Next: nil}
	list1C := &ListNode{Val: 2, Next: list1D}
	list1B := &ListNode{Val: 4, Next: list1C}
	list1A := &ListNode{Val: 2, Next: list1B}
	head := &ListNode{Val: 1, Next: list1A}
	partition(head, 3)
}
