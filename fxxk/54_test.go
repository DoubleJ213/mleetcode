package fxxk

import "testing"

/*
54. 螺旋矩阵
给你一个 m 行 n 列的矩阵matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/

func spiralOrder(matrix [][]int) []int {
	depth := len(matrix)
	if depth == 0 {
		return nil
	}
	wide := len(matrix[0])
	if wide == 0 {
		return nil
	}
	data := make([]int, 0)
	up := 0
	down := depth - 1
	left := 0
	right := wide - 1
	for len(data) < depth*wide {
		if up <= down {
			for i := left; i <= right; i++ {
				data = append(data, matrix[up][i])
			}
			up++
		}
		if left <= right {
			for j := up; j <= down; j++ {
				data = append(data, matrix[j][right])
			}
			right--
		}
		if up <= down {
			for k := right; k >= left; k-- {
				data = append(data, matrix[down][k])
			}
			down--
		}

		if left <= right {
			for l := down; l >= up; l-- {
				data = append(data, matrix[l][left])
			}
			left++
		}
	}
	return data
}

func spiralOrder2(matrix [][]int) []int {
	//内存占用小了，但是cpu占用高了
	depth := len(matrix)
	if depth == 0 {
		return nil
	}
	wide := len(matrix[0])
	if wide == 0 {
		return nil
	}
	data := make([]int, depth*wide)
	up := 0
	down := depth - 1
	left := 0
	right := wide - 1
	count := 0
	for count < depth*wide {
		if up <= down {
			for i := left; i <= right; i++ {
				data[count] = matrix[up][i]
				count++
			}
			up++
		}
		if left <= right {
			for j := up; j <= down; j++ {
				data[count] = matrix[j][right]
				count++
			}
			right--
		}
		if up <= down {
			for k := right; k >= left; k-- {
				data[count] = matrix[down][k]
				count++
			}
			down--
		}

		if left <= right {
			for l := down; l >= up; l-- {
				data[count] = matrix[l][left]
				count++
			}
			left++
		}
	}
	return data
}

func TestAl54(t *testing.T) {
	//spiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}})
	spiralOrder([][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}})
}
