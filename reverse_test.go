package mleetcode

import (
	"fmt"
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
func reverseList1(head *ListNode) *ListNode {
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
	reverseList(res.Next)
	//reverseList2(res.Next)
}

func reverseList(head *ListNode) {
	var res *ListNode
	//这题的精髓是构造了一个空的节点，每次将这个节点追加到当前的next，不断遍历成倒序
	for head != nil {
		a := head.Next
		head.Next = res
		res = head
		head = a
	}
	fmt.Println(res)
}
