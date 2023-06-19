package fxxk

import (
	"container/heap"
	"fmt"
	"testing"
)

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

//func (pq *PriorityQueue) Update(item *Item, value int, priority int) {
//	item.value = value
//	item.priority = priority
//	heap.Fix(pq, item.index)
//}

func (pq *PriorityQueue) Remove(value int) {
	for i, v := range *pq {
		if v != nil && v.value == value {
			heap.Remove(pq, i)
			break
		}
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	n := len(nums)
	ans := make([]int, n-k+1)
	for i := 0; i < n; i++ {
		start := i - k
		if start >= 0 {
			pq.Remove(nums[start])
		}
		item := &Item{
			value:    nums[i],
			priority: nums[i],
		}

		heap.Push(&pq, item)
		if pq.Len() == k {
			ans[start+1] = pq[0].value
		}
	}
	return ans
}

func TestMaxSlidingWindow(t *testing.T) {
	nums2 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	nums := []int{7, 2, 4}
	nums3 := []int{9, 10, 9, -7, -4, -8, 2, -6}

	aaa := maxSlidingWindow(nums2, 3)
	bbb := maxSlidingWindow(nums, 2)
	ccc := maxSlidingWindow(nums3, 5)

	fmt.Printf("%v", aaa)
	fmt.Printf("%v", bbb)
	fmt.Printf("%v", ccc)
}

func maxSlidingWindow4(nums []int, k int) []int {
	if nums == nil || len(nums) == 0 || len(nums) < k {
		return nil
	}
	window := make([]int, 0)
	result := make([]int, 0)
	pushIntoWindow := func(i int) {
		for len(window) > 0 && nums[i] >= nums[window[len(window)-1]] {
			//窗口中已经有元素。且窗口最后一个没有自己大，把他干掉就行
			window = window[:len(window)-1]
		}
		//把自己加进去
		window = append(window, i)
	}
	for i := 0; i < k; i++ {
		//开始构造滑动窗口。前面K个元素，直接往里面丢就行。不需要考虑最前面元素过期的问题
		pushIntoWindow(i)
	}
	result = append(result, nums[window[0]])
	for i := k; i < len(nums); i++ {
		pushIntoWindow(i)
		if window[0] <= i-k {
			window = window[1:]
		}
		result = append(result, nums[window[0]])
	}
	return result
}
