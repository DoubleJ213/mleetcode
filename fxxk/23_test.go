package fxxk

import (
	"container/heap"
	"testing"
)

//给你一个链表数组，每个链表都已经按升序排列。
//请你将所有链表合并到一个升序链表中，返回合并后的链表。

//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6

//示例 2：
//
//输入：lists = []
//输出：[]
//示例 3：
//
//输入：lists = [[]]
//输出：[]

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type MyPQ []*ListNode

//Len() int
//Less(i, j int) bool
//Swap(i, j int)
//Push(x interface{}) // add x as element Len()
//Pop() interface{}   // remove and return element Len() - 1.

func (pq *MyPQ) Less(i, j int) bool {
	return (*pq)[i].Val < (*pq)[j].Val
}

func (pq *MyPQ) Len() int {
	return len(*pq)
}

func (pq *MyPQ) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *MyPQ) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *MyPQ) Pop() interface{} {
	tmp := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return tmp
}

//MaxHeap 自己实现的堆性能还可以优化,这里不用了
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	//空节点
	dummy := &ListNode{-1, nil}
	p := dummy
	//初始化队列
	myQueue := make(MyPQ, 0)
	for _, list := range lists {
		if list != nil {
			myQueue = append(myQueue, list)
		}
	}
	heap.Init(&myQueue)
	for myQueue.Len() > 0 {
		node := heap.Pop(&myQueue).(*ListNode)
		p.Next = node
		if node != nil && node.Next != nil {
			node = node.Next
			p.Next.Next = nil
			heap.Push(&myQueue, node)
		}
		p = p.Next
	}

	return dummy.Next
}

func TestAl23(t *testing.T) {
	root2 := &ListNode{
		Val:  5,
		Next: nil,
	}
	root1 := &ListNode{
		Val:  4,
		Next: root2,
	}
	_ = &ListNode{
		Val:  1,
		Next: root1,
	}

	head2 := &ListNode{
		Val:  4,
		Next: nil,
	}
	head1 := &ListNode{
		Val:  3,
		Next: head2,
	}
	_ = &ListNode{
		Val:  1,
		Next: head1,
	}
	head := &ListNode{2, nil}
	in := make([]*ListNode, 0)
	in = append(in, nil)
	in = append(in, head)
	mergeKLists(in)
}

/*
[[],[1]]

[[],[]]

[[]]

[]
*/
