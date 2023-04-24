package fxxk

import (
	"fmt"
	"testing"
)

/*
1094. 拼车
车上最初有capacity个空座位。车只能向一个方向行驶（也就是说，不允许掉头或改变方向）
给定整数capacity和一个数组 trips,
trip[i] = [numPassengersi, fromi, toi]表示第 i 次旅行有numPassengersi乘客，
接他们和放他们的位置分别是fromi和toi。这些位置是从汽车的初始位置向东的公里数。
当且仅当你可以在所有给定的行程中接送所有乘客时，返回true，否则请返回 false。

示例 1：
输入：trips = [[2,1,5],[3,3,7]], capacity = 4
输出：false
示例 2：
输入：trips = [[2,1,5],[3,3,7]], capacity = 5
输出：true

提示：
1 <= trips.length <= 1000
trips[i].length == 3
1 <= numPassengersi<= 100
0 <= fromi< toi<= 1000
1 <= capacity <= 105
*/

// carPooling
// 把1109的差分数组结构体直接拿来用
// 性能肯定不是最好的，数组长度拉到最长了。
// 算出差分数组，算出全部最终数组才进行判断，不能提前返回false
// 没有考虑到下车的情况，[[2,1,5],[3,5,7]] 比如第5公里 前面2个人下车了，后面3个人上来了。其实是2个人的
// 所以 数组 5】 7】都是开区间
func carPooling(trips [][]int, capacity int) bool {
	input := make([]int, 1000, 1000)
	nd := NewDifference(input)
	for _, trip := range trips {
		nd.inCrease(trip[1], trip[2]-1, trip[0])
	}
	output := nd.getOutPut(0, nd.size-1)
	for _, out := range output {
		if out > capacity {
			return false
		}
	}
	return true
}

func TestAl1094(t *testing.T) {
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 5, 7}}, 3))
}
