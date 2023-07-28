package fxxk

import (
	"fmt"
	"testing"
)

/*
304. 二维区域和检索 - 矩阵不可变

给定一个二维矩阵matrix，以下类型的多个请求：
计算其子矩形范围内元素的总和，该子矩阵的左上角为(row1,col1)，右下角为(row2,col2)。
实现NumMatrix类：
NumMatrix(int[][] matrix)给定整数矩阵matrix进行初始化
int sumRegion(int row1, int col1, int row2, int col2)返回 左上角 (row1,col1)、右下角(row2,col2) 所描述的子矩阵的元素 总和 。
输入:
["NumMatrix","sumRegion","sumRegion","sumRegion"]
[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
输出:
[null, 8, 11, 12]

解释:
NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]);
numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)

提示：
m == matrix.length
n == matrix[i].length
1 <= m,n <=200
-105<= matrix[i][j] <= 105
0 <= row1 <= row2 < m
0 <= col1 <= col2 < n
最多调用 104 次sumRegion 方法
*/

//row行、col列。第一行/第一列 坐标为0
type NumMatrix struct {
	Sum [][]int
}

//前缀和主要适用的场景是原始数组不会被修改的情况下，频繁查询某个区间的累加和。
func Constructor304(matrix [][]int) *NumMatrix {
	if matrix == nil {
		return &NumMatrix{
			Sum: nil,
		}
	}
	rowSize := len(matrix)
	if rowSize == 0 {
		return &NumMatrix{
			Sum: nil,
		}
	}
	colSize := len(matrix[0])
	if colSize == 0 {
		return &NumMatrix{
			Sum: nil,
		}
	}
	sum := make([][]int, rowSize+1)
	for x := 0; x < rowSize+1; x++ {
		sum[x] = make([]int, colSize+1)
		for y := 0; y < colSize+1; y++ {
			if x == 0 {
				sum[0][y] = 0
			}
		}
		sum[x][0] = 0
	}

	for i := 1; i <= rowSize; i++ {
		for j := 1; j <= colSize; j++ {
			sum[i][j] = sum[i][j-1] + sum[i-1][j] + matrix[i-1][j-1] - sum[i-1][j-1]
		}
	}
	return &NumMatrix{
		Sum: sum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if this.Sum == nil {
		return 0
	}
	A := this.Sum[row2+1][col2+1]
	B := this.Sum[row2+1][col1]
	C := this.Sum[row1][col2+1]
	D := this.Sum[row1][col1]
	return A + D - B - C
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor1(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func TestAl304(t *testing.T) {
	sub1 := []int{3, 0, 1, 4, 2}
	sub2 := []int{5, 6, 3, 2, 1}
	sub3 := []int{1, 2, 0, 1, 5}
	sub4 := []int{4, 1, 0, 1, 7}
	sub5 := []int{1, 0, 3, 0, 5}
	hehe := [][]int{sub1, sub2, sub3, sub4, sub5}

	fmt.Printf("len %d \n", len(hehe))
	obj := Constructor304([][]int{})
	//fmt.Println(obj.SumRegion(1, 1, 2, 2))

	fmt.Println(obj.SumRegion(0, 0, 0, 0))

	fmt.Println(obj.SumRegion(2, 1, 4, 3))
	// return 8 (红色矩形框的元素总和)
	fmt.Println(obj.SumRegion(1, 1, 2, 2))
	// return 11 (绿色矩形框的元素总和)
	fmt.Println(obj.SumRegion(1, 2, 2, 4))
	// return 12 (蓝色矩形框的元素总和)

}
