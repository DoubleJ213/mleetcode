package fxxk

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{-1, nil}
	head := result
	for list1 != nil || list2 != nil {
		if list1 == nil {
			result.Next = list2
			break
		}
		if list2 == nil {
			result.Next = list1
			break
		}
		if list1.Val <= list2.Val {
			result.Next = list1
			list1 = list1.Next
			result.Next.Next = nil
		} else {
			result.Next = list2
			list2 = list2.Next
			result.Next.Next = nil
		}
		result = result.Next
	}
	return head.Next
}

func TestAl21(t *testing.T) {
	list1B := &ListNode{Val: 4, Next: nil}
	list1A := &ListNode{Val: 2, Next: list1B}
	_ = &ListNode{Val: 1, Next: list1A}

	list2B := &ListNode{Val: 4, Next: nil}
	list2A := &ListNode{Val: 3, Next: list2B}
	_ = &ListNode{Val: 1, Next: list2A}
	mergeTwoLists(nil, nil)
}
