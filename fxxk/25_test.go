package fxxk

import (
	"fmt"
	"testing"
)

//双指针技巧主要分为两类：左右指针和快慢指针
//链表由于没有右边起点，较多的是快慢指针，数组有快慢和左右指针两种常用方法

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//25. K 个一组翻转链表

//给你链表的头节点 head ，每k个节点一组进行翻转，请你返回修改后的链表。
//k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换

func reverseKGroup(head *ListNode, k int) *ListNode {
	//myReverse head开始k个
	end := head
	for i := 0; i < k; i++ {
		if end == nil {
			//尾指针指到nil,不够翻转的数量了，直接和newHead 接上
			return head
		} else {
			end = end.Next
		}
	}
	//反转前 k 个元素
	//newHead, preEnd := myReverse(head, end)

	// 递归反转后续链表并连接起来
	//preEnd.Next = reverseKGroup(end, k)

	/*
		为什么是 head.Next 得思考下
		newHead 前面k个被翻转，那原先的首节点，连的就是当前的输入end
		下一段的开始,end起点，k个（子问题）
		head 变成了首节点，接着当前end的链表
		子问题翻转得到的值，重新接到head.Next后面
	*/

	newHead := reverser(head, end)
	head.Next = reverseKGroup(end, k)

	return newHead
}

func myReverse(head *ListNode, end *ListNode) (newHead, res *ListNode) {
	//翻转头结点head,尾节点end的，end不包含
	preEnd := head
	p := end
	for head != end {
		tmp := head.Next
		head.Next = p
		p = head
		head = tmp
	}
	//preEnd 就是原来的头结点，就是head,由外面传进来的，没必要再返回出去
	return p, preEnd
}

func reverser(head *ListNode, end *ListNode) *ListNode {
	//翻转头结点head,尾节点end的，end不包含
	p := end
	for head != end {
		tmp := head.Next
		head.Next = p
		p = head
		head = tmp
	}
	return p

}

func myReverse1(head *ListNode) *ListNode {
	//翻转头结点head,尾节点nil.这个就是常规的链表翻转
	var p *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = p
		p = head
		head = tmp
	}
	return p
}

func TestAl25(t *testing.T) {
	fmt.Printf("fxxk: %v\n", 1)
	head := &ListNode{-1, nil}
	res := head
	for i := 1; i < 8; i++ {
		head.Next = &ListNode{i, nil}
		head = head.Next
	}
	//myReverse1(res.Next)
	//myReverse(res.Next, res.Next.Next.Next.Next)
	a := reverseKGroup(res.Next, 2)
	fmt.Println("123", a)
}
