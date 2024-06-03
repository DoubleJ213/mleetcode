package fxxk

import (
	"fmt"
	"testing"
)

//	给第一个小朋友 1 颗糖果，第二个小朋友 2 颗，依此类推，直到给最后一个小朋友 n 颗糖果。

// leetcode submit region begin(Prohibit modification and deletion)
func distributeCandies1103(candies int, num_people int) []int {
	//num_people 排队的人数  candies 糖果总数
	res := make([]int, num_people)
	index := 0
	count := 1
	for candies > 0 {
		if count > candies {
			res[index] += candies
			break
		} else {
			res[index] += count
			candies -= count
		}
		index++
		count++
		if index == num_people {
			index = 0
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDistributeCandiesToPeople(t *testing.T) {
	//fmt.Println(distributeCandies(7, 4))
	fmt.Println(distributeCandies1103(10, 3))

}
