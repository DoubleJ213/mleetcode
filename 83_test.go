package mleetcode

import (
	"fmt"
	"testing"
)

func TestAl83(*testing.T) {
	node7 := &ListNode{Next: nil, Val: 5}
	node6 := &ListNode{Next: node7, Val: 5}
	//node5 := &ListNode{Next: node6, Val: 4}
	node4 := &ListNode{Next: node6, Val: 2}
	node3 := &ListNode{Next: node4, Val: 2}
	node2 := &ListNode{Next: node3, Val: 2}
	node := &ListNode{Next: node2, Val: 1}

	deleteDuplicates(node)
	fmt.Println(node.Next)
}

//存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
//返回同样按升序排列的结果链表。
//注意审题，只保留没有重复出现过的数字
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	res := &ListNode{Val: -1, Next: head}
	node := res
	for node.Next != nil && node.Next.Next != nil {
		if node.Next.Val == node.Next.Next.Val {
			repeatVal := node.Next.Val
			for node.Next != nil && node.Next.Val == repeatVal {
				//体会这题的巧妙之处，不停的比对第二和第三个节点，如果相同一直让后面的节点覆盖第二个节点
				node.Next = node.Next.Next
			}
		} else {
			node = node.Next
		}
	}

	return res.Next
}
