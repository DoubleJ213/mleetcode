package mleetcode

import (
	"container/heap"
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
type IntList []int

func (intList IntList) Push(x interface{}) {
	panic("implement me")
}

func (intList IntList) Pop() interface{} {
	panic("implement me")
}

func (intList IntList) Len() int {
	return len(intList)
}

func (intList IntList) Less(i, j int) bool {
	//reports whether the element with index i should sort before the element with index j
	//返回i是否应该在j前面 ,如果是< 表示从小到大排序
	return intList[i] < intList[j]
}

func (intList IntList) Swap(i, j int) {
	intList[i], intList[j] = intList[j], intList[i]
}

type KthLargest struct {
	k   int
	arr *IntList
}

func Constructor(k int, nums []int) KthLargest {
	sort.Sort(IntList(nums))
	//heap.Init()
	if len(nums) > k {
		nums = nums[len(nums)-k:]
	}
	MaxKNums := IntList(nums)
	heap.Init(&MaxKNums)
	return KthLargest{k, &MaxKNums}

}

func (this *KthLargest) Add(val int) int {

	return 0
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func TestKthLargest(t *testing.T) {
	k := 3
	var arr = []int{4, 5, 8, 2}

	obj := Constructor(k, arr)
	obj.Add(3)

}
