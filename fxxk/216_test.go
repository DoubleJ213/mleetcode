package fxxk

import (
	"fmt"
	"testing"
)

/*
216. 组合总和 III
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
只使用数字1到9 每个数字 最多使用一次 返回 所有可能的有效组合的列表 。
该列表不能包含相同的组合两次，组合可以以任何顺序返回。

示例 1:
输入: k = 3, n = 7
输出: [[1,2,4]]
解释:
1 + 2 + 4 = 7
没有其他符合的组合了。
示例 2:
输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
解释:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
没有其他符合的组合了。
示例 3:
输入: k = 4, n = 1
输出: []
解释: 不存在有效的组合。
在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。

提示:
2 <= k <= 9
1 <= n <= 60
*/

var (
	res216     [][]int
	path216    []int
	pathSum216 int
	//used216    []bool
)

/*
k个数，凑到和为n
*/
func combinationSum216(k int, n int) [][]int {
	res216 = make([][]int, 0)
	path216 = make([]int, 0)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//used216 = make([]bool, len(nums))
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	if sum > n {
		return res216
	}
	traverse216(nums, pathSum216, 0, k, n, 0)
	return res216
}

func traverse216(nums []int, pathSum216, depth, count, target, start int) {
	if depth == count {
		if pathSum216 == target {
			tmp := make([]int, len(path216))
			copy(tmp, path216)
			res216 = append(res216, tmp)
		}
		return
	}

	for i := start; i < len(nums); i++ {
		path216 = append(path216, nums[i])
		pathSum216 += nums[i]
		traverse216(nums, pathSum216, depth+1, count, target, i+1)
		//traverse216(nums, pathSum216, depth+1, count, target)
		pathSum216 -= nums[i]
		path216 = path216[:len(path216)-1]
		//上一次递归的数不满足要求，改成false
	}
}

func getNewNums(nums []int, i int) []int {
	res := make([]int, len(nums)-1)
	copy(res[0:i], nums[0:i])
	copy(res[i:], nums[i+1:])
	return res
}

func TestAl216(t *testing.T) {
	fmt.Printf("%v", combinationSum216(3, 9))
}
