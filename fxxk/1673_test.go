package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func mostCompetitive(nums []int, k int) []int {
	ans := make([]int, 0)
	for i, num := range nums {
		for len(ans) > 0 && num < ans[len(ans)-1] && len(ans)+len(nums)-i > k {
			//栈内元素个数+剩余元素个数 > k，此时才能弹出，否则凑不齐k个元素
			ans = ans[:len(ans)-1]
		}
		if len(ans) < k {
			ans = append(ans, num)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

// 超时，之前还写了递归，剪枝后也超时
func mostCompetitive2(nums []int, k int) []int {
	ans := make([]int, 0)
	index := 0
	for i := 0; i < k; i++ {
		next := index
		for left := next; left < len(nums)-k+i+1; left++ {
			if len(ans) < i+1 {
				ans = append(ans, nums[left])
				index = left + 1
			} else if nums[left] < ans[len(ans)-1] {
				ans[len(ans)-1] = nums[left]
				index = left + 1
			}
		}
	}
	return ans
}

func TestFindTheMostCompetitiveSubsequence(t *testing.T) {
	//fmt.Println(mostCompetitive([]int{3, 2, 6, 4, 5}, 3))
	//fmt.Println(mostCompetitive([]int{2, 4, 3, 3, 5, 4, 9, 6}, 4))
	fmt.Println(mostCompetitive([]int{11, 52, 57, 91, 47, 95, 86, 46, 87, 47,
		70, 56, 54, 61, 89, 44, 3, 73, 1, 7, 87, 48, 17, 25, 49,
		54, 6, 72, 97, 62, 16, 11, 47, 34, 68, 58, 14, 36, 46, 65, 2, 15}, 18))
	//fmt.Println(mostCompetitive([]int{70, 3, 34, 32, 18, 67, 51, 79, 44, 69, 16, 0, 20, 89, 43, 4, 13, 22, 62, 54, 61, 64, 18, 53, 98, 84, 48, 17, 73}, 26))
	//fmt.Println(mostCompetitive([]int{2, 10, 3, 5, 9, 4, 2, 0, 6, 7, 8, 0, 6, 5, 8,
	//	1, 6, 1, 5, 5, 2, 10, 9, 5, 7, 7, 3, 2, 1, 4, 0, 7, 0, 3, 10, 10, 5, 10,
	//	4, 7, 0, 2, 10, 9, 0, 2, 6, 10, 6, 9, 2, 1, 9, 8, 7, 2, 0, 7, 3, 6, 2, 1, 8,
	//	0, 0, 0, 10, 4, 3, 5, 0, 8, 1, 8, 5, 1, 6, 0, 4, 4, 10, 2, 0, 5, 1, 1, 3, 3,
	//	5, 2, 6, 5, 6, 0, 3, 8, 0, 1, 7, 0, 0, 9, 6, 10, 5, 9, 8, 9, 8, 7, 8, 10, 6,
	//	3, 8, 0, 5, 7, 4, 3, 5, 7, 7, 0, 3, 10, 1, 3, 10, 2, 10, 3, 2, 6, 3, 10, 8,
	//	10, 6, 0, 7, 6, 2, 10, 4, 0, 7, 4, 8, 8, 1, 7, 1, 4, 9, 7, 7, 8, 9, 8, 7, 2,
	//	4, 9, 8, 8, 0, 8, 2, 10, 7, 3, 10, 8, 5, 1, 1, 3, 0, 5, 1, 7, 1, 7, 9, 2, 6,
	//	9, 6, 10, 6, 1, 7, 8, 3, 6, 9, 3, 5, 9, 0, 9, 3, 5, 8, 4, 6, 8, 10, 8, 0, 9,
	//	3, 7, 10, 4, 4, 2, 3, 7, 2, 10, 3, 5, 4, 9, 9, 2, 1, 2, 10, 4, 4, 4, 3, 5, 9,
	//	7, 2, 0, 3, 6, 6, 7, 3, 9, 4, 6, 9, 7, 1, 3, 2, 3, 6, 6, 1, 7, 10, 0, 4, 10,
	//	3, 5, 0, 10, 3, 10, 3, 0, 0, 1, 6, 6, 5, 9, 10, 5, 5, 9, 0, 5, 4, 1, 10, 2, 3,
	//	1, 7, 9, 10, 10, 4, 3, 5, 9, 5, 4, 4, 8, 0, 1, 8, 1, 4, 6, 5, 6, 0, 6, 8, 6, 5, 6, 5,
	//	7, 9, 5, 8, 8, 4, 2, 0, 0, 2, 9, 4, 9, 2, 6, 5, 2, 2, 8, 5, 4, 10, 8, 7, 7, 3, 4,
	//	2, 0, 4, 3, 8, 6, 1, 7, 10, 10, 7, 4, 0, 6, 6, 0, 5, 6, 10, 3, 8, 3, 2, 4, 10, 4,
	//	3, 0, 4, 10, 7, 6, 0, 4, 7, 0, 5, 2, 5, 2, 10, 9, 1, 10, 9, 6, 6, 5, 9, 10, 1, 3,
	//	5, 2, 0, 6, 8, 5, 6, 3, 4, 8, 4, 0, 7, 0, 7, 9, 9, 1, 4, 6, 4, 5, 7, 3, 0, 4, 4,
	//	9, 10, 5, 10, 3, 9, 6, 6, 2, 9, 4, 0, 4, 3, 3, 1, 7, 2, 1, 0, 2, 6, 7, 1, 1, 0,
	//	3, 9, 8, 9, 4, 6, 3, 10, 7, 3, 1, 5, 2, 0, 3, 9, 5, 3, 3, 3, 1, 7, 5, 8, 10, 10,
	//	8, 0, 2, 3, 3, 2, 9, 3, 1, 3, 9, 0, 1, 8, 2, 1, 6, 0, 6, 3, 1, 3, 1, 10, 5, 6, 0, 4, 7, 10}, 79))
}

func moansCompetitive2(nums []int, k int) []int {
	for i := 0; i < len(nums)-k; i++ {
		//第一个元素确认，下面开始找最小的第二个元素？
		for j := i + k; j < len(nums); j++ {

		}
	}
	return nil
}
