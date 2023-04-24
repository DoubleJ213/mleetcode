package fxxk

import (
	"fmt"
	"testing"
)

func TestJ(t *testing.T) {
	//obj := Constructor()
	//obj.AddAtHead(1)
	//obj.DeleteAtIndex(0)
	//fmt.Println(obj.Get(0))

	//obj := Constructor()
	//obj.AddAtHead(1)
	//obj.AddAtTail(3)
	//obj.AddAtIndex(1, 2)
	//fmt.Println(obj.Get(1))
	//obj.DeleteAtIndex(0)
	//fmt.Println(obj.Get(0))

	//obj := Constructor()
	//obj.AddAtHead(1)
	//obj.AddAtTail(3)
	//obj.AddAtIndex(1, 2)
	//fmt.Println(obj.Get(1))
	//obj.DeleteAtIndex(1)
	//fmt.Println(obj.Get(1))

	obj := JConstructor()
	obj.AddAtIndex(0, 10)
	obj.AddAtIndex(0, 20)
	obj.AddAtIndex(1, 30)
	fmt.Println(obj.Get(0))

	//obj := Constructor()
	//obj.AddAtTail(1)
	//obj.AddAtTail(2)
	//fmt.Println(obj.Get(0))
}

//707. 设计链表
//设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val和next。val是当前节点的值，next是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性prev以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。
//在链表类中实现这些功能：
//get(index)：获取链表中第index个节点的值。如果索引无效，则返回-1。
//addAtHead(val)：在链表的第一个元素之前添加一个值为val的节点。插入后，新节点将成为链表的第一个节点。
//addAtTail(val)：将值为val 的节点追加到链表的最后一个元素。
//addAtIndex(index,val)：在链表中的第index个节点之前添加值为val 的节点。如果index等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
//deleteAtIndex(index)：如果索引index 有效，则删除链表中的第index 个节点

type MyLinkedList struct {
	node *ListNode
}

/** Initialize your data structure here. */
func JConstructor() MyLinkedList {
	return MyLinkedList{
		node: &ListNode{Next: nil, Val: 1001},
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	i := 0
	tmp := this.node
	for tmp != nil {
		if index == i {
			return tmp.Val
		}
		tmp = tmp.Next
		i++
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.node.Val == 1001 && this.node.Next == nil {
		//	初始化的情况
		this.node.Val = val
		return
	}
	tmp := &ListNode{Val: val, Next: nil}
	tmp.Next = this.node
	this.node = tmp
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.node.Val == 1001 && this.node.Next == nil {
		this.node.Val = val
		return
	}
	tmp := this.node
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &ListNode{Val: val, Next: nil}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	i := 0
	tmp := this.node
	for tmp != nil {
		if i == index {
			this.AddAtTail(val)
		}

		if i+1 == index {
			tmp.Next = &ListNode{Val: val, Next: tmp.Next}
			return
		}
		tmp = tmp.Next
		i++
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.node = this.node.Next
	} else if index > 0 {
		i := 0
		p := this.node
		pre := &ListNode{}
		for p != nil {
			if index == i {
				if p.Next == nil {
					pre.Next = nil
				} else {
					pre.Next = p.Next
					pre.Next.Val = p.Next.Val
				}

			}
			i++
			pre = p
			p = p.Next
		}
	}

}
