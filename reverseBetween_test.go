package mleetcode

import (
	"fmt"
	"testing"
)

func LinkedListPrint(head *ListNode, name string) {
	tmp := head
	for tmp != nil {
		fmt.Print(" ", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println(" value of ", name)

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

func TestReverseBetween(t *testing.T) {
	node := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := node
	for i := 1; i < 3; i++ {
		tmp.Next = &ListNode{Val: i}
		tmp = tmp.Next
	}
	reverseBetween(node.Next, 1, 2)
	fmt.Println("done")
}
