package fxxk

import (
	"fmt"
	"testing"
)

/*
给定一个非空且只包含非负数的整数数组 nums，数组的 度 的定义是指数组里任一元素出现频数的最大值。
你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度

Your task is to find the smallest possible length of a (contiguous) subarray of nums,
that has the same degree as nums.
*/
//leetcode submit region begin(Prohibit modification and deletion)
type m697 struct {
	count int    //出现次数
	index [2]int //长度为2
}

func findShortestSubArray(nums []int) int {
	m := make(map[int]*m697)
	for i, num := range nums {
		if _, ok := m[num]; ok {
			if m[num].count >= 1 {
				m[num].index[1] = i
			}
			m[num].count += 1
		} else {
			m[num] = &m697{1, [2]int{i, 0}}
		}
	}
	var deep int
	ans := len(nums)
	for _, v := range m {
		if v.count == 1 && deep == 0 {
			ans = 1
			deep = v.count
		} else if v.count > deep {
			ans = v.index[1] - v.index[0] + 1
			deep = v.count
		} else if v.count == deep && deep != 1 {
			ans = getMin(ans, v.index[1]-v.index[0]+1)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDegreeOfAnArray(t *testing.T) {
	//fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
	//fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	fmt.Println(findShortestSubArray([]int{1, 2, 3}))

}
