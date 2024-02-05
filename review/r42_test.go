package review

import (
	"fmt"
	"testing"
)

/*
42. 接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

提示：
n == height.length
1 <= n <= 2 * 10^4
0 <= height[i] <= 10^5
*/

func trap(height []int) int {
	var totalArea int
	leftMax := 0
	for i := 0; i < len(height)-1; i++ {
		//最右侧的数一定不能装进去水，其实左侧的也不行，但是我需要知道左侧的边有多大
		current := height[i]
		levelArea := 0
		if leftMax <= current {
			leftMax = current
			levelArea = 0
		} else {
			rightMax := current
			for j := i + 1; j < len(height); j++ {
				right := height[j]
				if right > current {
					if right >= leftMax {
						rightMax = leftMax
						break
					} else {
						rightMax = getMax(rightMax, right)
					}
				}
			}
			levelArea = rightMax - current
		}
		totalArea += levelArea
	}
	return totalArea
}

func TestR42(t *testing.T) {
	fmt.Printf("3 -- %v ", trap([]int{2, 1, 0, 2}))
	fmt.Printf("9 -- %v ", trap([]int{4, 2, 0, 3, 2, 5}))
	fmt.Printf("4 -- %v ", trap([]int{2, 1, 0, 1, 2}))
	fmt.Printf("5 -- %v ", trap([]int{2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Printf("6 -- %v ", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Printf("4 -- %v ", trap([]int{0, 1, 2, 1, 0, 1, 3}))
	fmt.Printf("14 -- %v ", trap([]int{5, 2, 1, 2, 1, 5}))
}
