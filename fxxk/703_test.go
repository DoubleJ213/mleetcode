package fxxk

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

//703. 数据流中的第K大元素
//设计一个找到数据流中第K大元素的类（class）。注意是排序后的第K大元素，不是第K个不同的元素。
//你的KthLargest类需要一个同时接收整数k 和整数数组nums的构造器，它包含数据流中的初始元素。
//每次调用KthLargest.add，返回当前数据流中第K大的元素。
// 示例
//int k = 3;
//int[] arr = [4,5,8,2];
//KthLargest kthLargest = new KthLargest(3, arr);
//kthLargest.add(3); // returns 4
//kthLargest.add(5); // returns 5
//kthLargest.add(10); // returns 5
//kthLargest.add(9); // returns 8
//kthLargest.add(4); // returns 8

//说明:
//你可以假设 nums 的长度≥ k-1 且k ≥ 1。

// 自定义类型和类型别名
// 将MaxList定义为[]int类型
// 通过Type关键字的定义，MaxList就是一种新的类型，它具有[]int的特性，但是他们不是同一种类型。
type IntList []int

func (intList *IntList) Push(x interface{}) {
	*intList = append(*intList, x.(int))
}

func (intList *IntList) Pop() interface{} {
	//删除最后一个元素，并返回
	pre := *intList
	*intList = pre[:len(pre)-1]
	return pre[len(pre)-1]
}

func (intList *IntList) Len() int {
	return len(*intList)
}

// 为什么不能直接写 *intList[i]
// 可能是因为*的优先级没有[]高，所以你这里相当于 *(m[1])了
// 指针符号*不能直接使用在已经声明好的变量名上
// (*intList)[0]
func (intList *IntList) Less(i, j int) bool {
	//reports whether the element with index i should sort before the element with index j
	//返回i是否应该在j前面 ,如果是< 表示从小到大排序
	nums := *intList
	return nums[i] < nums[j]
}

func (intList *IntList) Swap(i, j int) {
	nums := *intList
	nums[i], nums[j] = nums[j], nums[i]
}

type KthLargest struct {
	k   int
	arr *IntList
}

func Constructor703(k int, nums []int) KthLargest {
	//TODO：如果这里不排序，是不是一个效果？不是一个效果，因为只维护k个数，如果不先排序就会丢掉原来比较大的数

	//对于相同底层结构类型的变量，go语言可以使用简单类型转换的方式
	// 字符类型
	//var a byte = 'y'
	// 转换为字符串类型
	//b := string(a)

	aaa := IntList(nums)
	sort.Sort(&aaa)

	//sort.Sort(IntList(nums)) 接受的参数是 interface类型，这里传aaa或者aaa的指针都可以
	//另外这个interface类型具体要求就是
	//type Interface interface {
	//	Len() int
	//	Less(i, j int) bool
	//	Swap(i, j int)
	//} 所以IntList这种结构体符合要求

	if len(nums) > k {
		//限制堆的大小
		nums = nums[len(nums)-k:]
	}
	MaxKNums := IntList(nums)
	// init的参数可以不是指针，语法上没问题，但是如果不是指针，将无法修改heap中的内容
	heap.Init(&MaxKNums)
	//heap.Init(&MaxKNums) 也是接受的interface类型的参数，，这个interface类型具体要求：
	//type Interface interface {
	//	sort.Interface  先实现sort.Interface
	//	Push(x interface{}) // add x as element Len()
	//	Pop() interface{}   // remove and return element Len() - 1.
	//}
	return KthLargest{k, &MaxKNums}
}

func (this *KthLargest) Add(val int) int {
	if len(*this.arr) < this.k {
		// heap.Push会重新平衡，没必要fix
		heap.Push(this.arr, val)
		//this.arr.Push(val)
		//heap.Fix(this.arr, 0)
		return (*this.arr)[0]
	}
	if val > (*this.arr)[0] {
		heap.Pop(this.arr)
		heap.Push(this.arr, val)
	}
	return (*this.arr)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func TestKthLargest(t *testing.T) {
	k := 2
	var arr = []int{0, 7, 3, 5}
	obj := Constructor703(k, arr)
	fmt.Printf("%v， %v\n", obj, obj.Add(-1))
	fmt.Printf("%v， %v\n", obj, obj.Add(1))
	fmt.Printf("%v， %v\n", obj, obj.Add(-2))
	fmt.Printf("%v， %v\n", obj, obj.Add(-4))
	fmt.Printf("%v， %v\n", obj, obj.Add(3))

	hehe := IntList([]int{0, 7, 3, 5})
	// init的参数可以不是指针，语法上没问题，但是如果不是指针，将无法修改heap中的内容
	heap.Init(&hehe)
	fmt.Printf("%v", hehe)
}
