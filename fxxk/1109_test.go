package fxxk

import (
	"fmt"
	"testing"
)

/*
1109. 航班预订统计
这里有n个航班，它们分别从 1 到 n 进行编号。
有一份航班预订表bookings ，
表中第i条预订记录bookings[i] = [firsti, lasti, seatsi]
意味着在从 firsti到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi个座位。
请你返回一个长度为 n 的数组answer，里面的元素是每个航班预定的座位总数。

示例 1：
输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
输出：[10,55,45,25,25]
解释：
航班编号        1   2   3   4   5
预订记录 1 ：   10  10
预订记录 2 ：       20  20
预订记录 3 ：       25  25  25  25
总座位数：      10  55  45  25  25
因此，answer = [10,55,45,25,25]

示例 2：
输入：bookings = [[1,2,10],[2,2,15]], n = 2
输出：[10,25]
解释：
航班编号        1   2
预订记录 1 ：   10  10
预订记录 2 ：       15
总座位数：      10  25
因此，answer = [10,25]

提示：
1 <= n <= 2 * 104
1 <= bookings.length <= 2 * 104
bookings[i].length == 3
1 <= firsti <= lasti <= n
1 <= seatsi <= 104
*/

//「差分数组」，差分数组的主要适用场景是频繁对原始数组的某个区间的元素进行增减。
//思路看    ![1109.png](.1109.png)
func corpFlightBookings(bookings [][]int, n int) []int {
	//默认数组全0，长度未知，长度可以按照n来
	nd := NewDifference(make([]int, n, n))
	for _, booking := range bookings {
		nd.inCrease(booking[0]-1, booking[1]-1, booking[2])
	}
	return nd.getOutPut(0, n-1)
}

type Difference struct {
	input     []int
	different []int
	size      int
}

//NewDifference  通过原数组初始化，差分数组结构体
func NewDifference(input []int) *Difference {
	//8,5,9,6,1
	//8,-3,4,-3,-5
	size := len(input)
	different := make([]int, size)
	different[0] = input[0]
	for i := 1; i < size; i++ {
		different[i] = input[i] - input[i-1]
	}
	return &Difference{
		input:     input,
		different: different,
		size:      size,
	}
}

//inCrease
//暂时减也调用这个方法，num变成负的就行
//这里i,j 表示数组索引，外部不是从0开始记数的自己转换
func (d *Difference) inCrease(i, j, num int) {
	//8,5,9,6,1
	//8,-3,4,-3,-5
	if 0 <= i && i <= j && j < d.size {
		d.different[i] += num
		if j+1 < d.size {
			d.different[j+1] -= num
		}
	} else {
		fmt.Println("input parma error check them")
	}
	return
}

//getOutPut  i，j 为索引，区间
func (d *Difference) getOutPut(i, j int) []int {
	output := make([]int, d.size, d.size)
	output[0] = d.different[0]
	for i := 1; i < d.size; i++ {
		output[i] = d.different[i] + output[i-1]
	}
	return output[i : j+1]
}

func TestAl1109(t *testing.T) {
	fmt.Println(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))
	//[10,55,45,25,25]
}
