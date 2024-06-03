package fxxk

import (
	"fmt"
	"testing"
)

/*
先一个map找重复。
然后^[索引] 长度+1
其中的值在 [1, n2] 范围内
[重复，缺失]
*/
//leetcode submit region begin(Prohibit modification and deletion)
func findMissingAndRepeatedValues(grid [][]int) []int {
	a := 0 //a 重复
	b := 0 //b 缺失
	numMap := make(map[int]bool)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			b ^= grid[i][j]
			if _, ok := numMap[grid[i][j]]; ok {
				a = grid[i][j]
			} else {
				numMap[grid[i][j]] = true
			}
		}
	}
	b ^= a
	for x := 1; x < len(grid)*len(grid[0])+1; x++ {
		//数字从1开始，如果3*2 6个数字，那就是1,2,3,4,5,6 这6个 x < m*n+1
		b ^= x
	}
	return []int{a, b}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindMissingAndRepeatedValues(t *testing.T) {
	fmt.Println(findMissingAndRepeatedValues([][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}))
}
