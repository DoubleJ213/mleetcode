package fxxk

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

// leetcode 206
// 翻转，不是俩俩翻转
func reverseList1(head *ListNode) *ListNode {
	var tmp *ListNode
	for head != nil {
		head.Next, tmp, head = tmp, head, head.Next
	}
	return tmp
}

func reverseList3(head *ListNode) *ListNode {
	return help(nil, head)
}

// nil,  1,2,3,4,5
// nil,1, 2,3,4,5
func help(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return help(head, next)
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
	//reverseList(res.Next)
	//reverseList1(res.Next)
	reverseList3(res.Next)
	//reverseList2(res.Next)
	reverseList4(res.Next)
}

func reverseList4(root *ListNode) *ListNode {
	//肯定先走到 root.Next 为空, 即最后一个递归6就行
	//if root.Next == nil {
	//但是 得考虑到root为空的情况
	//if root ==nil || root.Next == nil {
	if root == nil || root.Next == nil {
		return root
	}
	//那么输入 reverseList4(root *ListNode)后，会在这里进行递归：
	//所以reverseList4(root.Next)
	//当递归到最后，末尾的节点变成新的头结点
	newHead := reverseList4(root.Next)
	//第一轮出栈，head为5，head.next为空，返回5
	//第二轮出栈，head为4，head.next为5，执行head.next.next=head也就是5.next=4，
	//把当前节点的子节点的子节点指向当前节点
	//此时链表为1->2->3->4<->5，由于4与5互相指向，所以此处要断开4.next=null
	//此时链表为1->2->3->4<-5
	//返回节点5
	root.Next.Next = root
	root.Next = nil
	return newHead
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
