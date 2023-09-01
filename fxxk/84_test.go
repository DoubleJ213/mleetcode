package fxxk

import (
	"fmt"
	"testing"
)

/*84. 柱状图中最大的矩形
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。

示例 1:
输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10

示例 2：
输入： heights = [2,4]
输出： 4

提示：
1 <= heights.length <=105
0 <= heights[i] <= 104
*/

/*
	stack中存索引，通过索引计算索引对应的高度
	找到当前高度最大的面积，除了要知道高度，还得知道宽度
	宽度怎么来，从当前索引的位置往左右两边找，找到左右两边都比自己小的两个索引
	最左边和最右边当成高度为0
	右边可以通过和stack顶端元素比较大小判断，左端怎么处理呢，比如2,1,5,6,2,3
	0先push进stack，高度2的最左边是0，index=0开始，到右边 index1结束
	1位置最大的面积，遍历到1的时候，右边一直比1大，所以一直求不出以1为高度的最大面积，直到遍历到最后自己加的元素index6 value0
	以1为高度的最大面积，最右端才能定下来，最左端怎么定，很明显知道是最左边高度为0的那个。
	那对应到stack中怎么理解
	左边界
	由于我们这个stack是个单调递增的stack，所以找左边比自己小的，肯定是pop完后新的顶部是第一个比自己小的
	右边界呢
	遍历的时候当前位置如果被push进栈，说明当前超过了之前的高度。
	如果当前元素不能加入stack，那他就是前一个元素的右边界。
*/
func largestRectangleArea(heights []int) int {
	stack84 := newStack()
	var maxArea int
	for i := 0; i < len(heights); i++ {
		for !stack84.Empty() && heights[stack84.Peek()] > heights[i] {
			high := heights[stack84.Pop()]
			right := i
			var left int
			var wide int
			if !stack84.Empty() {
				left = stack84.Peek()
				wide = right - left - 1
			} else {
				wide = right
			}
			area := wide * high
			if area > maxArea {
				maxArea = area
			}
		}
		stack84.Push(i)
	}
	//走到这里表示heights遍历完成，但是stack中可以还存在数据，还得继续计算面积
	for !stack84.Empty() {
		high := heights[stack84.Pop()]
		right := len(heights)
		var left, wide int
		if !stack84.Empty() {
			left = stack84.Peek()
			wide = right - left - 1
		} else {
			wide = right
		}
		area := wide * high
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func TestAl84(t *testing.T) {
	fmt.Printf("%v", largestRectangleArea([]int{1, 2, 2, 3, 4}))
	fmt.Printf("%v", largestRectangleArea2([]int{1, 2, 2, 3, 4}))
}

func largestRectangleArea2(heights []int) int {
	if len(heights) == 1 {
		return heights[0]
	}
	stack84 := make([]int, 0)
	var maxArea int
	for i := 0; i < len(heights); i++ {
		for len(stack84) != 0 && heights[stack84[len(stack84)-1]] > heights[i] {
			high := heights[stack84[len(stack84)-1]]
			stack84 = stack84[:len(stack84)-1]
			right := i
			var left, wide int
			if len(stack84) != 0 {
				left = stack84[len(stack84)-1]
				wide = right - left - 1
			} else {
				wide = right
			}
			area := wide * high
			if area > maxArea {
				maxArea = area
			}
		}
		stack84 = append(stack84, i)
	}
	//走到这里表示heights遍历完成，但是stack中可以还存在数据，还得继续计算面积
	for len(stack84) != 0 {
		high := heights[stack84[len(stack84)-1]]
		stack84 = stack84[:len(stack84)-1]
		right := len(heights)
		var left, wide int
		if len(stack84) != 0 {
			left = stack84[len(stack84)-1]
			wide = right - left - 1
		} else {
			wide = right
		}
		area := wide * high
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
