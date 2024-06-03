package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func distributeCandies575(candyType []int) int {
	m := make(map[int]bool)
	count := 0
	for i := 0; i < len(candyType); i++ {
		if _, ok := m[candyType[i]]; ok {
			continue
		} else {
			m[candyType[i]] = true
			count++
		}
	}
	if count >= len(candyType)/2 {
		return len(candyType) / 2
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDistributeCandies(t *testing.T) {
	//fmt.Println(distributeCandies575([]int{2, 2, 2, 2, 2, 2, 3, 3}))
	fmt.Println(distributeCandies575([]int{1, 1, 2, 3}))
}
