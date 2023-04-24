package fxxk

import "testing"

//leetcode 1431给你一个数组candies和一个整数extraCandies，其中candies[i]代表第 i 个孩子拥有的糖果数目。
//对每一个孩子，检查是否存在一种方案，将额外的extraCandies个糖果分配给孩子们之后，此孩子有 最多的糖果。注意，允许有多个孩子同时拥有 最多的糖果数目。

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for _, eachItem := range candies {
		if eachItem > max {
			max = eachItem
		}
	}
	var ret []bool
	for _, eachItem := range candies {
		if eachItem+extraCandies >= max {
			ret = append(ret, true)
		} else {
			ret = append(ret, false)
		}
	}
	return ret
}

func TestKidsWithCandies(t *testing.T) {
	var ca []int
	ca = append(ca, 2, 2, 3, 4)
	var caa = make([]int, 0)
	caa = append(caa, 1)

	var ka = []int{1, 2, 3, 4, 5}
	kidsWithCandies(ka, 3)
}
