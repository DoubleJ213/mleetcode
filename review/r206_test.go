package review

import "testing"

func TestR206(t *testing.T) {
	res := &ListNode{0, nil}
	temp := res
	for i := 1; i < 6; i++ {
		temp.Next = &ListNode{i, nil}
		temp = temp.Next
	}
	reverse206(res.Next)
	//reverse206_2(res.Next)
}

func reverse206(head *ListNode) *ListNode {
	var tmp *ListNode
	for head != nil {
		a := head.Next
		head.Next = tmp
		tmp = head
		head = a
	}
	return tmp
}
