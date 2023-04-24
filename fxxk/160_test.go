package fxxk

import (
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//题目数据 保证 整个链式结构中不存在环。
//注意，函数返回结果后，链表必须保持其原始结构。

//进阶：你能否设计一个时间复杂度 O(m + n) 、仅用 O(1) 内存的解决方案？
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 先遍历第一个链表，把所有节点都存到一个map中
	// 然后再遍历第二个链表
	if headA == nil || headB == nil {
		return nil
	}
	nodeMap := make(map[*ListNode]bool)
	tmpA := headA
	for tmpA != nil {
		nodeMap[tmpA] = true
		tmpA = tmpA.Next
	}
	tmpB := headB
	for tmpB != nil {
		_, ok := nodeMap[tmpB]
		if ok {
			return tmpB
		}
		tmpB = tmpB.Next
	}

	//fmt.Println("")
	return nil
}

//进阶：你能否设计一个时间复杂度 O(m + n) 、仅用 O(1) 内存的解决方案？
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	//内存o1,那就是不能借助map来解决问题
	//链表很多题目的思路就是搞2个指针，移来移去比较
	//这里类似于判断链表有环，如果有环返回环的节点那块的逻辑
	//如两个指针为la 和lb，链表a、b同时遍历
	//la lb 如果不同时遍历，时间复杂度做不到o(m+n),存在嵌套
	//需要同时遍历找到重复都点

	//假设 链表长度 分别为m、n，不相交的节点，分别为a、b个，相交的节点为c个
	//那m=a+c n=b+c。
	//如果 a=b,那么两个指针一次遍历，走长度a或者b返回即可。
	//如果a!=b，随便a大还是b大，我们让a走完走b、b走完走a。确保指针走的长度是相同的。
	// 都走了m+n的长度，a走的顺序是 a+c+b+c。b走的顺序是b+c+a+c
	// 规律来了。最后公共的c找到了。

	//本题还能把链表都翻转后再比较，复杂度好像符合要求？

	if headA == nil || headB == nil {
		return nil
	}
	tmpA, tmpB := headA, headB
	tmpC, tmpD := headA, headB
	for tmpC != nil || tmpD != nil {
		if tmpA == nil {
			tmpA = headB
		}
		if tmpB == nil {
			tmpB = headA
		}
		if tmpA == tmpB {
			return tmpB
		} else {
			tmpA = tmpA.Next
			tmpB = tmpB.Next
		}
		if tmpC != nil {
			tmpC = tmpC.Next
		} else {
			tmpD = tmpD.Next
		}
	}
	return nil
}

func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	//2 写得有些啰嗦
	if headA == nil || headB == nil {
		return nil
	}
	tmpA, tmpB := headA, headB
	//当指针走完一遍 a和b的长度之后，应该会同时为nil。tmpA == tmpB
	for tmpA != tmpB {
		//if tmpA == tmpB {
		//	return tmpB
		//}
		if tmpA == nil {
			tmpA = headB
		} else {
			tmpA = tmpA.Next
		}
		if tmpB == nil {
			tmpB = headA
		} else {
			tmpB = tmpB.Next
		}
	}
	return tmpA
	//	题目描述有毒
}

func TestAl160(t *testing.T) {
	root4 := &ListNode{
		Val:  5,
		Next: nil,
	}
	root3 := &ListNode{
		Val:  4,
		Next: root4,
	}
	root2 := &ListNode{
		Val:  8,
		Next: root3,
	}
	root1 := &ListNode{
		Val:  1,
		Next: root2,
	}
	root := &ListNode{
		Val:  4,
		Next: root1,
	}

	head5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	head4 := &ListNode{
		Val:  4,
		Next: head5,
	}
	head3 := &ListNode{
		Val:  8,
		Next: head4,
	}
	head2 := &ListNode{
		Val:  1,
		Next: head3,
	}
	head1 := &ListNode{
		Val:  6,
		Next: head2,
	}
	head := &ListNode{
		Val:  5,
		Next: head1,
	}

	//上面这种例子，永远不想等
	//所以题目描述也很蛋疼

	//请你找出并返回两个单链表相交的起始节点。
	//不是相交一下就结束，是某个点开始都相交了。

	//abc := getIntersectionNode(root.Next, root.Next.Next.Next)
	//xyz := getIntersectionNode2(root.Next, root.Next.Next.Next)
	//xyz1 := getIntersectionNode3(root, head)
	getIntersectionNode3(root, head)
	//fmt.Println(abc, xyz, xyz1)
}
