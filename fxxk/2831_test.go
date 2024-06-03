package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func longestEqualSubarray(nums []int, k int) int {
	position := make(map[int][]int) //key  num  value array of index
	for index, num := range nums {
		if km, ok := position[num]; ok {
			position[num] = append(km, index)
		} else {
			position[num] = []int{index}
		}
	}
	ans := 1 // 1 <= nums.length <= 10^5  0 <= k <= nums.length
	for _, indexList := range position {
		//	遍历map
		left := 0
		tmp := 0
		for right := 1; right < len(indexList); right++ {
			//indexList[right] - indexList[right-1] - 1 两个索引间有几个元素要删除
			tmp += indexList[right] - indexList[right-1] - 1
			for tmp > k && left < right {
				tmp -= indexList[left+1] - indexList[left] - 1
				left++
			}
			ans = getMax(ans, right-left+1) //ans = max(ans, r-l+1) golang 1.21
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindTheLongestEqualSubarray(t *testing.T) {
	//fmt.Println(longestEqualSubarray([]int{1, 3, 2, 3, 1, 3}, 3))
	fmt.Println(longestEqualSubarray([]int{1}, 0))
}
