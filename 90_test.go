package mleetcode

import (
	"fmt"
	"sort"
	"testing"
)

/*90. 子集 II
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

示例 1：
输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
示例 2：
输入：nums = [0]
输出：[[],[0]]

提示：
1 <= nums.length <= 10
-10 <= nums[i] <= 10
*/
func subsetsWithDup(nums []int) [][]int {
	res90 := make([][]int, 0)
	//看方法名都知道，这题和78区别就在于输入元素可以重复
	//但是又不是简单去除重复的元素，1，2，2和2，2这种也要保留
	//可以考虑先排序，在加入res90 的时候进行判断，如果已存在不添加, 其余逻辑和78题保持不变。
	//但是这么想虽然简单，复杂度上去了
	//之前我们作此类题目的时候已经知道，for循环是横向展开的，递归是纵向的，所以只需要for循环中去重即可

	sort.Ints(nums)
	swdDfs(nums, []int{}, &res90)
	return res90
}

func swdDfs(nums, levelRes []int, res *[][]int) {
	*res = append(*res, append([]int{}, levelRes...))

	if len(nums) == 0 {
		return
	}
	for index, value := range nums {
		if index > 0 && nums[index-1] == value {
			continue
		}
		swdDfs(nums[index+1:], append(levelRes, value), res)
	}
}

func TestAl90(t *testing.T) {
	a := subsetsWithDup([]int{1, 2, 2})
	fmt.Print("done")
	_ = a
}
