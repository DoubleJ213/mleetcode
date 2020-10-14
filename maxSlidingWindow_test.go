package mleetcode

import (
	"sort"
	"testing"
)

//239. 滑动窗口最大值
type MyList []int

func (m MyList) Len() int {
        return len(m)
}

func (m MyList) Less(i, j int) bool {
	// 维护最大堆
        return m[i] > m[j]
}

func (m MyList) Swap(i, j int) {
        m[i], m[j] = m[j], m[i]
}

func maxSlidingWindow(nums []int, k int) []int {
	var ret_list []int
	//TODO：堆，优先队列（双端队列）
	//取前k个元素,排序和维护堆，原来nums不改变，方便后续for循环遍历
	new_nums := nums
	sort.Sort(MyList(new_nums))
	length := len(new_nums)
	if len(new_nums) > k {
		new_nums = new_nums[len(new_nums)-k:]
	}

	//heap.Init()
	for i := 0; i < length; i++ {
		if i < k {
			ret_list = append(ret_list, new_nums[0])
		} else {

		}

	}
	return nums
}

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	maxSlidingWindow(nums, k)
}
