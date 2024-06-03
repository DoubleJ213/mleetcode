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
func removeElements(head *test.ListNode, val int) *test.ListNode {
	if head == nil || head.Next == nil && head.Val != val {
		return head
	}
	if head.Val == val {
		return removeElements(head.Next, val)
	}
	cur := head
	//这里 第一个元素不等于 val，且不为空
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveLinkedListElements(t *testing.T) {
	node := &test.ListNode{
		Val: 1,
		Next: &test.ListNode{
			Val: 2,
			Next: &test.ListNode{
				Val: 2,
				Next: &test.ListNode{
					Val: 1,
					Next: &test.ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}
	a := removeElements(node, 2)
	fmt.Println(a)

}
