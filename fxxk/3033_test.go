package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func modifiedMatrix(matrix [][]int) [][]int {
	ans := make([][]int, len(matrix))
	replace := make([][]int, 0)
	maxList := make([]int, len(matrix[0])) //2 <= m, n <= 50
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			maxList[j] = getMax(maxList[j], matrix[i][j])
			if matrix[i][j] == -1 {
				replace = append(replace, []int{i, j})
			}
			ans[i] = make([]int, len(matrix[i]))
			copy(ans[i], matrix[i])
		}
	}
	for k := 0; k < len(replace); k++ {
		x, y := replace[k][0], replace[k][1]
		ans[x][y] = maxList[y]
	}
	return ans
}

//可以修改原数组

//leetcode submit region end(Prohibit modification and deletion)

func TestModifyTheMatrix(t *testing.T) {
	//输入：matrix = [[1,2,-1],[4,-1,6],[7,8,9]] 输出：[[1,2,9],[4,8,6],[7,8,9]]
	fmt.Println(modifiedMatrix([][]int{{1, 2, -1}, {4, -1, 6}, {7, 8, 9}, {9, -1, 4}}))
}
