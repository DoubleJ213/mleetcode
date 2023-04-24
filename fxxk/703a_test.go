package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

//703. 数据流中的第K大元素

// 如果自己实现最大堆， fuck 没见过这么长的简单题

type MyArray []int

//粗心了，写复杂了。没必要再包一层，k存在KthLargest1 结构体中就好
type MaxHeap struct {
	//0,1,2,3 最小的在最前面随时被替换
	nums *MyArray
	k    int
}

func (mh *MaxHeap) Init() {
	sort.Ints(*mh.nums)
	if len(*mh.nums) > mh.k {
		*mh.nums = (*mh.nums)[len(*mh.nums)-mh.k:]
	}
}

/*TODO:
不需要像数组一样严格排序，是用数组存放的数据，其实是个堆，保证父节点小于子节点就好
可以看heap.Init源码
也可以看703_pic.png
*/
func (mh *MaxHeap) Push(val int) {
	if len(*mh.nums) < mh.k {
		//堆不满，添加，重平衡
		*mh.nums = append(*mh.nums, val)
		for i := 0; i < len(*mh.nums)-1; i++ {
			last := len(*mh.nums) - 1
			if !mh.less(i, last) {
				mh.swap(i, last)
			}
		}
	} else {
		//	堆已满，pop最小的数，再push?
		top := (*mh.nums)[0]
		if top < val {
			mh.Pop()
			mh.Push(val)
		}
	}
}
func (mh *MaxHeap) swap(a, b int) {
	//a,b 为数组index
	(*mh.nums)[a], (*mh.nums)[b] = (*mh.nums)[b], (*mh.nums)[a]
}

func (mh *MaxHeap) less(a, b int) bool {
	//a,b 为数组index
	return (*mh.nums)[a] < (*mh.nums)[b]
}

//0,1,2,3 最小的在最前面随时被替换
func (mh *MaxHeap) Pop() int {
	nums := *mh.nums
	src := nums[1:]
	tmp := make([]int, len(src))
	copy(tmp, src)
	*mh.nums = tmp
	if len(*mh.nums) > 1 {
		return nums[0]
	}
	return 0
}

type KthLargest1 struct {
	MaxHeap
}

func Constructor1(k int, nums []int) KthLargest1 {
	ma := MyArray(nums)
	mh := MaxHeap{&ma, k}
	mh.Init()
	kl := KthLargest1{
		mh,
	}
	return kl
}

func (this *KthLargest1) Add(val int) int {
	if len(*this.nums) < this.k {
		this.Push(val)
		return (*this.nums)[0]
	}
	if val > (*this.nums)[0] {
		this.Pop()
		this.Push(val)
	}
	return (*this.nums)[0]
}

func TestAl703(t *testing.T) {
	k := 3
	var arr = []int{4, 5, 8, 2}
	obj := Constructor1(k, arr)
	fmt.Printf("%v， %v\n", obj, obj.Add(3))
	fmt.Printf("%v， %v\n", obj, obj.Add(5))
	fmt.Printf("%v， %v\n", obj, obj.Add(10))
	fmt.Printf("%v， %v\n", obj, obj.Add(9))
	fmt.Printf("%v， %v\n", obj, obj.Add(4))
}

/*

测试用例
["KthLargest","add","add","add","add","add"]
[[2,[0]],[-1],[1],[-2],[-4],[3]]

["KthLargest","add","add","add","add","add"]
[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]

*/
