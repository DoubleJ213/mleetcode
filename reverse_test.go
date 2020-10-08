package mleetcode

import (
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode 206
func reverseList(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		temp := head.Next
		head.Next = res
		res = head
		head = temp
	}

	return res
}

func reverseList2(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		head.Next, res, head = res, head, head.Next
	}
	return res
}

// TODO：递归实现

func TestReverseNode(t *testing.T) {
	res := &ListNode{0, nil}
	temp := res
	for i := 1; i < 6; i++ {
		temp.Next = &ListNode{i, nil}
		temp = temp.Next
	}
	reverseList2(res.Next)
}
