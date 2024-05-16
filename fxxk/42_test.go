package fxxk

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
1 <= n <= 2 * 104
0 <= height[i] <= 105
*/

/*
[0,1,0,2,1,0,1,3,2,1,2,1]
如102这个0上方能接水，因为两边都比自己大
比如2102 这个0两边都是比自己大大，然后按照最短大边一算，0和上面102一样的能接1*1个单位的水
但是因为1 还不是最高的，1左边有个2 右边也有个2，1也能接1单位的水，然后顺带把0这个坐标多接1单位的水也算出来
所以这个算法的思路有点清楚了。不是每个索引都能直接求出当前索引最大的接水量，而是先算最底下的量，然后再由其他高度的查再算另外的量
算出来的量的和才是最后总的接水的量。
现在再细想左右边界如何通过遍历确认
找到左边界后，如果来个更小的怎么处理，继续入栈，因为要防止存在右边界。
那什么时候才能确定真正的右边界，一次不能真正大确认右边界，需要一次次大确定，一次次大算面积
stack顶部元素已经是相对小的元素了，如果下一个元素比顶部元素大，就是一个右边界。
一边pop，一边算面积，可能一次算不到位，怎么才能右边界算到位了呢，右边界比左边界大或者右边没有新的元素了
所以这个stack不是单调的stack，应该是一个先找大，找到大后如果下一个元素小可以入栈，如果下一个元素更大，直接替换当前顶部元素
*/
func trap(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	stack42 := newStack()
	left := 0
	stack42.Push(left)
	var totalArea int
	for i := 1; i < len(height); i++ {
		current := height[i]
		for !stack42.Empty() && current > height[stack42.Peek()] {
			stackTop := stack42.Pop()
			var myLeft, leftIndex int
			if !stack42.Empty() {
				leftIndex = stack42.Peek()
				myLeft = height[leftIndex]
			}
			h := getMin(current, myLeft)
			high := h - height[stackTop]
			if h != 0 && high != 0 {
				wide := i - leftIndex - 1
				totalArea += wide * high
			}
		}
		if stack42.Empty() {
			left = i
			stack42.Push(i)
		} else if current < height[stack42.Peek()] {
			stack42.Push(i)
		} else if current == height[stack42.Peek()] {
			stack42.Pop()
			stack42.Push(i)
		}
	}
	return totalArea
}

func TestAl42(t *testing.T) {
	fmt.Printf("3 -- %v ", trap([]int{2, 1, 0, 2}))
	fmt.Printf("9 -- %v ", trap([]int{4, 2, 0, 3, 2, 5}))
	fmt.Printf("4 -- %v ", trap([]int{2, 1, 0, 1, 2}))
	fmt.Printf("5 -- %v ", trap([]int{2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Printf("6 -- %v ", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Printf("4 -- %v ", trap([]int{0, 1, 2, 1, 0, 1, 3}))
	fmt.Printf("14 -- %v ", trap([]int{5, 2, 1, 2, 1, 5}))
}
