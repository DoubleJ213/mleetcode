package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

//40. 组合总和 II
//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//candidates 中的每个数字在每个组合中只能使用一次。
//注意：解集不能包含重复的组合。
//
//示例 1:
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出: [[1,1,6], [1,2,5], [1,7], [2,6]]
//示例 2:
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:[ [1,2,2], [5]]

func combinationSum3(candidates []int, target int) [][]int {
	data := make([]int, 0)
	csRes := make([][]int, 0)
	sort.Ints(candidates)
	csDfs(candidates, target, data, &csRes)
	return csRes
}

func csDfs(candidates []int, target int, data []int, csRes *[][]int) {
	//for循环横向遍历，这个时候同一层去重处理 、dfs纵向递归遍历各个子树
	for index, value := range candidates {
		//这边value 和 candidates[index-1] 在递归进来的时候的判断要理解
		//candidates元素递归进来在不断变少
		if index > 0 && candidates[index-1] == value {
			continue
		} else if target < candidates[index] {
			//如果目标值剩余的数已经小于当前数，后面就可以不判断了
			return
		} else if target == candidates[index] {
			//如果目标剩余的值和当前值相等则结果合并，结束本次递归
			data = append(data, candidates[index])
			*csRes = append(*csRes, data)
			return
		} else {
			csDfs(candidates[index+1:], target-value, append(data, value), csRes)
		}
	}
}

func TestAl40(t *testing.T) {
	a := combinationSum3([]int{2, 5, 2, 1, 2}, 5)
	fmt.Print("abc")
	_ = a
}
