package fxxk

import (
	"fmt"
	"math"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	//abs(i - j) >= indexDifference
	//abs(nums[i] - nums[j]) >= valueDifference
	for i := 0; i < len(nums)-indexDifference; i++ {
		for j := i + indexDifference; j < len(nums); j++ {
			//0 <= valueDifference <= 50
			if int(math.Abs(float64(nums[i]-nums[j]))) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindIndicesWithIndexAndValueDifferenceI(t *testing.T) {
	fmt.Println(findIndices([]int{5, 1, 4, 1}, 2, 4))

}
