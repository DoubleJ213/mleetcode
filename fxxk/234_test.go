package fxxk

import (
	"fmt"
	"testing"
)

//给你一个单链表的头节点head，请你判断该链表是否为回文链表。如果是，返回true；否则，返回false。
//回文链表是指从头遍历与从后遍历得到的结点顺序一致,子回文串
//这题是让你判断是否为回文链表，不是问你是否,包含回文链表，所以仅仅子回文链表的不算，如下
//1-1-2-1 false  1-2-1 true

//进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func TestAl234(t *testing.T) {
	head := &ListNode{-1, nil}
	res := head
	head.Next = &ListNode{1, nil}
	head = head.Next
	head.Next = &ListNode{2, nil}
	head = head.Next
	head.Next = &ListNode{3, nil}
	head = head.Next
	head.Next = &ListNode{4, nil}
	head = head.Next
	fmt.Println(isPalindrome(res.Next))
}

func isPalindrome(head *ListNode) bool {
	/*
		能想到的方法：
		1.一遍遍历，所有数都找出来，放入数组，双指针判断空间复杂度O(n),不满足进阶要求不写了先
		2.快慢指针，找到中间点，然后翻转一半，比较。快慢指针遍历一次o(n)，直接一次遍历计算？拿不到中间节点，还得再走一遍。
	*/
	if head.Next == nil {
		return true
	}
	fast, slow := head, head
	//1-2-1-nil
	//s   f
	//  s
	//1-2-2-1-nil
	//偶数 fast == nil
	//奇数 fast.Next == nil 也可以说 fast != nil
	//奇数
	post := head
	//提前一次结束循环，不然等slow往后走一个，不好找slow前面的一个节点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		if fast.Next.Next == nil {
			//偶数
			post = slow
			break
		} else if fast.Next.Next.Next == nil {
			//	奇数
			post = slow.Next
			break
		}
		fast = fast.Next.Next
	}

	reverseHead := helper(head, slow)
	for post != nil {
		if post.Val != reverseHead.Val {
			return false
		}
		post = post.Next
		reverseHead = reverseHead.Next
	}
	return true
}

func helper(head, end *ListNode) *ListNode {
	//左开右闭的
	list := end
	for head != end {
		head.Next, list, head = list, head, head.Next
	}
	return list
}
