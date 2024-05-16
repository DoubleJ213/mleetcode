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
进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
前面有插入排序链表的，这题要求nLogn时间 o1空间 快排 归并都满足条件，
但是本题是链表如果新搞一个链表，那空间就是o(n)了所以只能原链表上操作。
*/
func sortList(head *test.ListNode) *test.ListNode {
	/*过一下几个 nlogn 时间复杂度的
	快排？相当于二叉树的前序遍历，把当前的数找到合适的位置，然后再排他的左\右子树
	需要两端像中间逼近，单链表倒序遍历。常数空间做不到
	归并排序？好像就挺满足的，先分割，然后合并
	相当于二叉树的后序遍历，左\右子树都排好了，然后合并结果
	空间可能不满足，不管了先写一个
	*/
	ans, _ := mergeSort148(head, nil)
	return ans
}

func mergeSort148(head *test.ListNode, end *test.ListNode) (*test.ListNode, *test.ListNode) {
	if head == end || head == nil || head.Next == end {
		return head, end
	}
	mid := getMid148(head, end)
	left, leftEnd := mergeSort148(head, mid)
	right, rightEnd := mergeSort148(mid, end)
	return merge148(left, leftEnd, right, rightEnd)
}

func merge148(left, leftEnd, right, rightEnd *test.ListNode) (*test.ListNode, *test.ListNode) {
	//其实不需要把末尾元素传进来right 就是 left 的末尾
	if left == nil || left == right {
		return right, rightEnd
	}
	if right == nil {
		return left, leftEnd
	}
	//这样一写 nlogn 空间？应该不是吧只多一个-1 的头指针，其余还是原来的指针
	ans := &test.ListNode{-1, nil}
	res := ans
	for left != nil && left != leftEnd && right != nil && right != rightEnd {
		if left.Val <= right.Val {
			ans.Next = left
			left = left.Next
		} else {
			ans.Next = right
			right = right.Next
		}
		ans = ans.Next
	}
	if left == leftEnd {
		ans.Next = right
		for right.Next != rightEnd {
			right = right.Next
		}
		right.Next = rightEnd
	}
	if right == rightEnd {
		ans.Next = left
		for left.Next != leftEnd {
			left = left.Next
		}
		left.Next = rightEnd
	}
	return res.Next, rightEnd
}

func getMid148(head *test.ListNode, end *test.ListNode) *test.ListNode {
	slow := head
	fast := head
	for fast != end && fast.Next != end {
		slow = slow.Next
		if fast.Next.Next != end {
			fast = fast.Next.Next
		} else {
			break
		}
	}
	return slow
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSortList(t *testing.T) {
	root := &test.ListNode{-1, nil}
	dump := root
	for i := 5; i > 1; i-- {
		dump.Next = &test.ListNode{i, nil}
		dump = dump.Next
	}
	sortList(root.Next)
}
