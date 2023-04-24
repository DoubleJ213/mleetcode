package fxxk

import (
	"fmt"
	"testing"
)

//
//栈（stack）也可以叫堆栈，但是不能叫堆（heap）先入后出

//232
type MyQueue struct {
	num []int
}

//void push(int x) 将元素 x 推到队列的末尾
//int pop() 从队列的开头移除并返回元素
//int peek() 返回队列开头的元素
//boolean empty() 如果队列为空，返回 true ；否则，返回 false

/** Initialize your data structure here. */
func MqConstructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.num = append(this.num, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	data := this.num[0]
	this.num = this.num[1:len(this.num)]
	return data
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	return this.num[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.num) == 0 {
		return true
	}
	return false
}

//数组实现队列，队列先入先出
//dequeue入队/enqueue出队是对应于队列的，队列是先入先出的线性表。
//push/pop是对应于栈的，栈是先入后出的线性表。
func TestC(t *testing.T) {
	obj := MqConstructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	fmt.Println(obj.Pop())
	obj.Push(4)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
