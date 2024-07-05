package fxxk

import (
	"fmt"
	"testing"
)

/*
给你一个满足下述两条属性的 m x n 整数矩阵：
每行中的整数从左到右按非严格递增顺序排列。
每行的第一个整数大于前一行的最后一个整数。
给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false
*/
/*二维数组展开成一维数组，然后二分法
左上角 0，0 右下角 len(matrix) len(matrix[0]) 因为是m*n的矩形
第一个数0表示 第二行第一个数3表示即 1*3+0（ x=2 (2-1)*y + 0） 最后一个  x*y-1
比如2行*3列矩阵 012 345 678*/
//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	x := len(matrix)    //多少行
	y := len(matrix[0]) //多少列
	l := 0
	r := x*y - 1
	for l <= r {
		m := (r-l)>>1 + l // 需要转换成合适的（x,y）
		mid := matrix[m/len(matrix[0])][m%len(matrix[0])]
		if mid == target {
			return true
		} else if mid > target {
			r = m - 1
		} else if mid < target {
			l = m + 1
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSearchA2dMatrix(t *testing.T) {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 7))
}
