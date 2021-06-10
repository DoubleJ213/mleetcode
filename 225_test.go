package mleetcode

import (
	"fmt"
	"testing"
)

//225 LIFO FILO 栈
type MyStack struct {
	num []int
}

/** Initialize your data structure here. */
func qConstructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.num = append(this.num, x)
}

/** Removes the element on top of the stack and returns that element. */
//返回并删除栈顶元素
func (this *MyStack) Pop() int {
	if this.Empty() {
		return -1
	}
	length := len(this.num)
	data := this.num[length-1]
	this.num = this.num[0 : length-1]
	return data
}

/** Get the top element. */
//返回栈顶元素
func (this *MyStack) Top() int {
	if this.Empty() {
		return -1
	}
	return this.num[len(this.num)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.num) == 0 {
		return true
	}
	return false
}

func TestD(t *testing.T) {
	obj := qConstructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
