package fxxk

import (
	"fmt"
	"testing"
)

/*739. 每日温度
给定一个整数数组temperatures，表示每天的温度，
返回一个数组answer，其中answer[i]是指对于第 i 天，下一个更高温度出现在几天后。
如果气温在这之后都不会升高，请在该位置用0 来代替。

示例 1:
输入: temperatures = [73,74,75,71,69,72,76,73]
输出:[1,1,4,2,1,1,0,0]
示例 2:
输入: temperatures = [30,40,50,60]
输出:[1,1,1,0]
示例 3:
输入: temperatures = [30,60,90]
输出: [1,1,0]

提示：
1 <=temperatures.length <= 105
30 <=temperatures[i]<= 100
*/

/*
496是单调栈 刷的第一题，复用496定义的数据结构
单调栈从后往前遍历，整一个map，key记录index，value记录温度
*/
func dailyTemperatures(temperatures []int) []int {
	stack := newStack()
	//之前这里都存的是更大的数，这里可以考虑存索引，比较大小的时候翻译一下就可以了
	length := len(temperatures)
	hotter := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		this := temperatures[i]
		//这次温度可以重复了
		for !stack.Empty() && temperatures[stack.Peek()] <= this {
			stack.Pop()
		}
		if stack.Empty() {
			hotter[i] = -1
		} else {
			hotter[i] = stack.Peek()
		}
		stack.Push(i)
	}
	result := make([]int, 0)
	for index, value := range hotter {
		if value == -1 {
			result = append(result, 0)
		} else {

			result = append(result, value-index)
		}
	}
	return result
}

func Test739(*testing.T) {
	fmt.Printf("%v", dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
