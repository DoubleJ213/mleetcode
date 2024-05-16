package fxxk

import (
	"fmt"
	"github.com/mlcPractice/leetcode/editor/cn"
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *test.ListNode) {
	//先翻转，然后遍历拼接
	//结果发现翻转期间，原来的链表结构被改坏了。那就一个List 存每个节点信息，然后拼接
	//再或者，不影响原来链表进行翻转.这个想了想没想起来
	//翻转后半部分，然后接在一起就行了
	mid := getMid(head)
	tail := reverse143(mid)
	root := head
	for tail != mid {
		rootTmp := root.Next
		tailTmp := tail.Next
		root.Next = tail
		tail.Next = rootTmp
		tail = tailTmp
		root = rootTmp
	}
}

/*
输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]
*/
func getMid(head *test.ListNode) *test.ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		if fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			break
		}
		fmt.Println("1")
	}
	return slow
}

func reverse143(node *test.ListNode) *test.ListNode {
	var dump *test.ListNode
	for node != nil {
		node.Next, dump, node = dump, node, node.Next
	}
	return dump
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReorderList(t *testing.T) {
	//node7 := &ListNode{Next: nil, Val: 7}
	//node6 := &ListNode{Next: node7, Val: 6}
	node6 := &test.ListNode{Next: nil, Val: 6}
	node5 := &test.ListNode{Next: node6, Val: 5}
	node4 := &test.ListNode{Next: node5, Val: 4}
	node3 := &test.ListNode{Next: node4, Val: 3}
	node2 := &test.ListNode{Next: node3, Val: 2}
	node := &test.ListNode{Next: node2, Val: 1}
	reorderList(node)
	fmt.Println("1")
}
