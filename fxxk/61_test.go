package fxxk

import (
	"fmt"
	"testing"
)

func TestG(t *testing.T) {
	node := &ListNode{Val: 0, Next: nil}
	tmp := node
	for i := 1; i < 3; i++ {
		tmp.Next = &ListNode{Val: i, Next: nil}
		tmp = tmp.Next
	}
	aaa := rotateRight(node, 1)
	fmt.Println(aaa)
}

// 61. 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	// 1-2-3-4-5 ，rotate 2
	// 4-5-1-2-3
	step := 0
	length := 1
	lNode, rNode := head, head
	for rNode.Next != nil {
		//不走到最后一个，不然rNode为nil不好组环
		step++
		length++
		rNode = rNode.Next
		if step > k {
			lNode = lNode.Next
		}
	}
	if length < k {
		//当k大于链表长度时，取余数再进行上述的循环操作
		k = k % length
		step := 0
		rNode = head
		for rNode.Next != nil {
			rNode = rNode.Next
			step++
			if step > k {
				lNode = lNode.Next
			}
		}
	}

	fmt.Println(head)
	rNode.Next = head
	head = rNode
	lNode.Next = nil
	return head
}
