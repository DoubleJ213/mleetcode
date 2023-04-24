package fxxk

import (
	"fmt"
	"testing"
)

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

	reverseBetween(node.Next, 1, 2)

	//翻转区间的，先看翻转前N个，n< 链表长度
	//reversePreN(node.Next, 3)
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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	//	一次遍历难以处理，那就2次遍历吧
	//	先把链表按照left、right切割成三段
	//  然后翻转中间那段
	//	最后合并
	//  0-1-2-3-4-5-6  0-1-4-3-2-5-6
	pre, res, post := getDetail(head, left, right)
	var data *ListNode
	if pre != nil {
		data = pre
	}
	for pre != nil && pre.Next != nil {
		pre = pre.Next
	}
	if data != nil {
		pre.Next = res
	} else {
		data = res
	}
	if res == nil {
		return data
	}
	for res.Next != nil {
		res = res.Next
	}
	res.Next = post
	return data
}

func getDetail(head *ListNode, left int, right int) (*ListNode, *ListNode, *ListNode) {
	i := 1
	tmp := head
	pre := head
	var toReverse *ListNode
	var post *ListNode
	getPre := false
	for tmp != nil {
		if left == 1 && toReverse == nil {
			pre = nil
			toReverse = tmp
			getPre = false
		}
		if left <= i+1 && toReverse == nil {
			toReverse = tmp.Next
			getPre = true
			tmp.Next = nil
		}

		if i >= right {
			post = tmp.Next
			tmp.Next = nil
			break
		}
		i++
		if getPre {
			tmp = toReverse
			getPre = false
		} else {
			tmp = tmp.Next
		}
	}
	var res *ListNode
	for toReverse != nil {
		a := toReverse.Next
		toReverse.Next = res
		res = toReverse
		toReverse = a
	}

	return pre, res, post
}

var after *ListNode

func reversePreN(head *ListNode, i int) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil || i == 1 {
		//1-2-3-4-5-6-nil
		//3-2-1-4-5-6-nil
		after = head.Next
		return head
	}
	newHead := reversePreN(head.Next, i-1)
	head.Next.Next = head
	head.Next = after
	return newHead
}
