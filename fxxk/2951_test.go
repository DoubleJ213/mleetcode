package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 以数组形式返回给定数组中 峰值 的下标，顺序不限
// 峰值 是指一个严格大于其相邻元素的元素。
// 数组的第一个和最后一个元素 不 是峰值。
func findPeaks(mountain []int) []int {
	var ans []int
	for index, m := range mountain {
		if index == 0 || index == len(mountain)-1 {
			continue
		}
		if mountain[index-1] < m && m > mountain[index+1] {
			ans = append(ans, index)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindThePeaks(t *testing.T) {
	//fmt.Println(findPeaks([]int{1, 4, 3, 8, 5}))
	fmt.Println(findPeaks([]int{1, 4, 3}))

}
