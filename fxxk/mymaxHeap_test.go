package fxxk

import (
	"fmt"
	"testing"
)

func TestF1(t *testing.T) {
	input := []int{4, 5, 8, 2}
	nums := MyArray(input)
	myHeap := &MaxHeap{&nums, 3}
	myHeap.Init()
	fmt.Printf("init  myHeap: %v\n", myHeap.nums)
	myHeap.Push(3)
	fmt.Printf("add 3 myHeap: %v\n", myHeap.nums)
	myHeap.Push(5)
	fmt.Printf("add 5 myHeap: %v\n", myHeap.nums)
	myHeap.Push(10)
	fmt.Printf("add 10 myHeap: %v\n", myHeap.nums)
	myHeap.Push(9)
	fmt.Printf("add 9 myHeap: %v\n", myHeap.nums)
	myHeap.Push(4)
	fmt.Printf("add 4 myHeap: %v\n", myHeap.nums)

}
