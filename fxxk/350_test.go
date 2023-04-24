package fxxk

import (
	"fmt"
	"testing"
)

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		//先循环小的，大的数组后判断，节约map空间，降低空间复杂度
		return intersect(nums2, nums1)
	}
	//m := map[int]int{}
	resMap := make(map[int]int)
	for _, val := range nums1 {
		resMap[val] += 1
	}
	intersection := make([]int, 0)
	for _, val := range nums2 {
		if resMap[val] > 0 {
			intersection = append(intersection, val)
			resMap[val] -= 1
		}
	}
	return intersection
}

func TestAl350(t *testing.T) {
	fmt.Printf("intersection %v \n", intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Printf("intersection %v \n", intersect([]int{1, 2, 3, 1}, []int{2, 3, 4}))
}
