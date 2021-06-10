package mleetcode

import (
	"fmt"
	"testing"
)

func DisplayNode(head *ListNode, name string) {
	tmp := head
	for tmp != nil {
		fmt.Print(tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println(" value of", name)
}

func TestK(*testing.T) {
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
	aa := reverseBetween1(node.Next, 2, 5)
	fmt.Println(aa)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//92. 反转链表 II
//left、right为索引
func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	//一遍 遍历 完成翻转
	// 1-2-3-4-5-6
	// 1 -2- 3-4-5-6
	// 1 -3- 2-4-5-6 3插入到1之后，原来1之后3之前的接到3之后，3之后的接到3之前的位置
	// 1 -4- 3-2-5-6 4插入到1之后，原来1之后4之前的接到4之后，4之后的接到4之前的位置
	// 1 -5- 4-3-2-6 5插入到1之后，原来1之后5之前的接到5之后，5之后的接到5之前的位置
	tmp := &ListNode{Val: -1, Next: head}
	pre := tmp
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = cur.Next.Next
		//cur.Next = next.Next
		next.Next = pre.Next
		//next.Next = cur
		pre.Next = next
	}
	//return head 有问题
	//[3,5]
	//1
	//2
	return tmp.Next
}
