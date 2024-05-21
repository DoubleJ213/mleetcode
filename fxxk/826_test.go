package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
/*
输入: difficulty = [2,4,6,8,10], profit = [10,20,30,40,50], worker = [4,5,6,7]
输出: 100 解释: 工人被分配的工作难度是 [4,4,6,6] ，分别获得 [20,20,30,30] 的收益。
*/
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	/*
		遍历worker中每个工人的工作能力，二分法找当前工人的最大收益，然后计算对应收益的值？
		造测试用例的时候，发现：不一定是难度更高的任务收益更高。目的是找最高收益
		那改一下计算方法。整一个map，key value key为价值，value为对应的难度，如果40这个价值，对应2个难度，比如1 和2 都是40收益。那value等于1好了
		可以方便后续计算，只要能干难度1的活，就可以获得40收益
	*/
	profitMap := make(map[int]int)
	res826 := 0
	if len(difficulty) != len(profit) {
		return 0
	}
	for i := 0; i < len(profit); i++ {
		if diff, ok := profitMap[profit[i]]; ok {
			if diff < difficulty[i] {
				//不更新的情况，存在，且已经存的难度小于当前的
				continue
			}
		}
		profitMap[profit[i]] = difficulty[i]
	}
	sort.Ints(worker)
	sort.Ints(profit)
	x := len(profit) - 1
	y := len(worker) - 1
	for x >= 0 && y >= 0 {
		minDiff := profitMap[profit[x]] //找大于等于这个数的worker
		if worker[y] >= minDiff {
			//此时最大的worker 能干这个难度，价值加上去，继续判断下一个worker
			y--
			res826 += profit[x]
		} else {
			//此时最大的worker已经干不了这个活了。判断下一个低价值的
			x--
		}
	}
	return res826
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMostProfitAssigningWork(t *testing.T) {
	fmt.Println(maxProfitAssignment([]int{10}, []int{10}, []int{4, 5, 6, 7}))
}
