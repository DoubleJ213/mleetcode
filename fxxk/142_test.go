package fxxk

import (
	"fmt"
	"testing"
)

func TestH(t *testing.T) {
	node := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := node
	for i := 1; i < 5; i++ {
		tmp.Next = &ListNode{Val: i, Next: nil}
		tmp = tmp.Next
		if i == 4 {
			tmp.Next = node.Next.Next
		}
	}
	pos := detectCycle(node)
	fmt.Println(pos)
}

// 142. 环形链表 II
// 给定一个链表，返回链表开始入环的第一个节点。如果链表无环，则返回 null。
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
// 说明：不允许修改给定的链表

func detectCycle(head *ListNode) *ListNode {
	// 为什么快慢指针差1个步长。
	// 快指针可以为3为4为5等等，但是为2时更容易相遇，设置为2可以降低时间复杂度
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	fast, slow := head.Next.Next, head.Next
	for {
		if fast == slow {
			break
		}
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 如果跳出循环证明有环，下面是取头结点，当前slow、fast已经指向环内了
	// 因为之前fast比slow快一步，fast走过的总步数是slow的2倍
	// D+S1+n(S2+S1)= 2(D+S1) D为起点到环入口 S1为环入口到相遇，S2为相遇到环入口
	// n假设为1 D=S2  假设 n为2 S1+S2+S2=D 不管多转多少圈，一定在入口相遇
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
