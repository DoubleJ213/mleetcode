package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findMinimumTime(tasks [][]int) int {
	//对于每一个任务，我们从结束时间 end 开始，从后往前选择尽可能靠后的点，
	//这样这些点更有可能被后面的任务重复利用。
	res := 0
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][1] < tasks[j][1] })
	maxSize := tasks[len(tasks)-1][1]
	//visited := make([]bool, maxSize+1)
	visited := make([]int, maxSize+1) //发现用0 1 表示更容易直接统计有多少个直接结果相加 0 表示未用过 1 表示用过
	//visited[1] 表示时间1 是否被前面的任务唤醒使用。0 没有实际意义
	for i := 0; i < len(tasks); i++ {
		start, end, duration := tasks[i][0], tasks[i][1], tasks[i][2]
		for j := start; j <= end; j++ {
			//duration 剩余需要的时间片
			duration -= visited[j]
		}
		//需要的时间片大于已经有的时间片 新开启时间片计入总结果
		if duration > 0 {
			res += duration
		}
		for k := end; k >= start && duration > 0; k-- {
			//visited 数组标记为最新状态方便后续继续判断
			if visited[k] != 1 {
				visited[k] = 1
				duration--
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumTimeToCompleteAllTasks(t *testing.T) {
	//findMinimumTime([][]int{{4, 7, 1}, {1, 5, 2}, {2, 3, 1}})

	//fmt.Println(findMinimumTime([][]int{{1, 3, 2}, {2, 5, 3}, {5, 6, 2}, {4, 7, 1}, {1, 5, 2}}))
	//fmt.Println(findMinimumTime([][]int{{2, 13, 2}, {6, 18, 5}, {2, 13, 3}}))
	//	[[2,13,2],[6,18,5],[2,13,3]]
	fmt.Println(findMinimumTime([][]int{{2, 3, 1}, {4, 5, 1}, {1, 5, 2}}))
}
