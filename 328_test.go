package mleetcode

import (
	"fmt"
	"testing"
)

func TestL(*testing.T) {
	node := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := node
	for i := 1; i < 7; i++ {
		tmp.Next = &ListNode{Val: i}
		tmp = tmp.Next
	}
	//DisplayNode(node, "node")
	aa := oddEvenList(node.Next)
	fmt.Println(aa)
}

//给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。
//请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
//请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

//说明
//应当保持奇数节点和偶数节点的相对顺序。
//链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
func oddEvenList(head *ListNode) *ListNode {
	odd := &ListNode{-1, nil}
	even := &ListNode{-1, nil}
	odd_res := odd
	even_res := even
	index := 1
	for head != nil {
		if index%2 == 0 {
			a := head.Next
			head.Next = nil
			even.Next = head
			even = even.Next
			head = a
		} else {
			a := head.Next
			head.Next = nil
			odd.Next = head
			odd = odd.Next
			head = a
		}
		index++
	}
	odd.Next = even_res.Next
	return odd_res.Next
}
