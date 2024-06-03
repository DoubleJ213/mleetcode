package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func kthLargestValue(matrix [][]int, k int) int {
	//矩阵中坐标 (a, b) 的 目标值
	//可以通过对所有元素 matrix[i][j] 执行异或运算得到，
	//其中 i 和 j 满足 0 <= i <= a < m 且 0 <= j <= b < n（下标从 0 开始计数）。
	res := make([][]int, len(matrix))
	ansList := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		res[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 && j == 0 {
				res[0][0] = matrix[0][0]
			} else if i == 0 {
				res[0][j] = res[0][j-1] ^ matrix[0][j]
			} else if j == 0 {
				res[i][0] = res[i-1][0] ^ matrix[i][0]
			} else {
				res[i][j] = res[i-1][j] ^ res[i][j-1] ^ res[i-1][j-1] ^ matrix[i][j]
			}
			ansList = append(ansList, res[i][j])
		}
	}
	sort.Ints(ansList)
	return ansList[len(ansList)-k] //1 <= k <= m * n
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindKthLargestXorCoordinateValue(t *testing.T) {
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 2))
}
