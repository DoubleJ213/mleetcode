package fxxk

import (
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

/*
没有太多特殊的地方。 递增。
//head [5,2,7,5]  ans [2,3,5,5] 为返回结果
第一个链表中遍历。然后再循环遍历ans中的结果 。
head=5-2-7-5  ans 为空 所以 ans 为 5
head 往后移动 head = 2-7-5
head 2 < ans 5 时 ans递增，我比你的头还小，那ans直接接在后面
tmp := head.next 存起来 方便后续 往后移动 head.next = ans   ans = head   head = tmp
继续往后移动 head=7 ans为 2-5 得找到一个合适7 插入的地方
head 7 >= ans 2 时 此时不能直接把7接到2后面，得确定2后面要么是nil 要么这个数比7大
如果 head不能直接直接接到2后面，ans往后移动继续判断，知道找到7能摆得下的地方。
*/
func insertionSortList(head *test.ListNode) *test.ListNode {
	var ans *test.ListNode
	for head != nil {
		if ans == nil {
			ans = &test.ListNode{Val: head.Val, Next: nil}
			head = head.Next
		}
		index := ans
		for head != nil && index != nil {
			if head.Val <= index.Val {
				head, head.Next, ans = head.Next, ans, head
				break
			} else {
				if index.Next == nil || index.Next.Val >= head.Val {
					//插入
					head, index.Next, head.Next = head.Next, head, index.Next
					break
				} else {
					index = index.Next
				}
			}
		}
	}
	return ans
	//	写完构造了个长度5000的链表，过了
}

//leetcode submit region end(Prohibit modification and deletion)

func TestInsertionSortList(t *testing.T) {
	//dump := &ListNode{Val: -1, Next: nil}
	//root := dump
	//for i := 5; i > 0; i-- {
	//	dump.Next = &ListNode{Val: i, Next: nil}
	//	dump = dump.Next
	//}
	//node4 := &ListNode{Next: nil, Val: 5}
	//node3 := &ListNode{Next: node4, Val: 7}
	//node2 := &ListNode{Next: node3, Val: 2}
	node := &test.ListNode{Next: nil, Val: 5}
	//insertionSortList(root.Next)
	insertionSortList(node)
}
