package mleetcode

import "testing"

//leetcode 1431
func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for _, eachItem := range candies {
		if eachItem > max {
			max = eachItem
		}
	}
	var ret []bool
	for _, eachItem := range candies {
		if eachItem+extraCandies >= max {
			ret = append(ret, true)
		} else {
			ret = append(ret, false)
		}
	}
	return ret
}

func TestKidsWithCandies(t *testing.T) {
	var ca []int
	ca = append(ca, 2, 2, 3, 4)
	var caa = make([]int, 0)
	caa = append(caa, 1)

	var ka = []int{1, 2, 3, 4, 5}
	kidsWithCandies(ka, 3)
}
