package mleetcode

import (
	"fmt"
	"testing"
)

//单调队列
type queue struct {
	value []int
}

func (q *queue) IsEmpty() bool {
	return len(q.value) == 0
}

func (q *queue) push(num int) {
	q.value = append(q.value, num)
}

func (q *queue) popFront() int {
	if q.IsEmpty() == true {
		return -1
	}
	tmp := q.value[0]
	q.value = q.value[1:]
	return tmp
}

func (q *queue) front() int {
	return q.value[0]
}

func (q *queue) popBack() int {
	if q.IsEmpty() == true {
		return -1
	}
	tmp := q.value[len(q.value)-1]
	q.value = q.value[:len(q.value)-1]
	return tmp
}

func (q *queue) back() int {
	if q.IsEmpty() == true {
		return -1
	}
	return q.value[len(q.value)-1]
}

type maxQueue struct {
	data queue
}

func (m *maxQueue) push(n int) {
	for m.data.IsEmpty() == false && m.data.back() < n {
		m.data.popBack()
	}
	m.data.push(n)
}

func (m *maxQueue) max() int {
	return m.data.front()
}

func (m *maxQueue) pop(n int) {
	if m.data.IsEmpty() == false && m.data.front() == n {
		m.data.popFront()
	}
}

func maxSlidingWindow2(nums []int, k int) []int {
	var res []int
	window := maxQueue{}
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			window.push(nums[i])
		} else {
			window.push(nums[i])
			res = append(res, window.max())
			window.pop(nums[i-k+1])
		}
	}
	return res
}

func TestMaxSlidingWindow2(t *testing.T) {
	nums2 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	nums := []int{7, 2, 4}
	nums3 := []int{9, 10, 9, -7, -4, -8, 2, -6}

	aaa := maxSlidingWindow2(nums2, 3)
	bbb := maxSlidingWindow2(nums, 2)
	ccc := maxSlidingWindow2(nums3, 5)

	fmt.Printf("%v", aaa)
	fmt.Printf("%v", bbb)
	fmt.Printf("%v", ccc)
}
