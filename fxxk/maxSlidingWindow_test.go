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
