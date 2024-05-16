package fxxk

import (
	"fmt"
	"testing"
)

/*
718. 最长重复子数组
给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。

示例 1：
输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
输出：3
解释：长度最长的公共子数组是 [3,2,1] 。
示例 2：
输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
输出：5

提示：
1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 100
*/

var dp718 [][]int

// 最长重复子数组
func findLengthErr(nums1 []int, nums2 []int) int {
	y := len(nums1)
	x := len(nums2)
	dp718 = make([][]int, x+1)
	for i := 0; i < x+1; i++ {
		dp718[i] = make([]int, y+1)
	}
	for j := 0; j < x; j++ { //nums2
		for k := 0; k < y; k++ { // nums1
			if nums2[j] == nums1[k] {
				dp718[j+1][k+1] = getMax(dp718[j][k+1], dp718[j+1][k]) + 1
			} else {
				dp718[j+1][k+1] = getMax(dp718[j][k+1], dp718[j+1][k])
			}
		}
	}
	/*
		dp718 = make([][]int, x+1)
		for i := 0; i < x+1; i++ {
			dp718[i] = make([]int, y+1)
		}
		结合初始化代码，数组 行长度为y+1，列长度为x+1
		二维数组  arr[纵坐标][横坐标]
	*/
	return dp718[x][y]
}

var max718 int

// 最长重复子数组
func findLength(nums1 []int, nums2 []int) int {
	y := len(nums1)
	x := len(nums2)
	max718 = 0
	dp718 = make([][]int, x+1)
	for i := 0; i < x+1; i++ {
		dp718[i] = make([]int, y+1)
	}
	for j := 0; j < x; j++ { //nums2
		for k := 0; k < y; k++ { // nums1
			//当数据全是0的时候发现不能横纵坐标取最大值再加1
			//而是对角线上的，两个数组都是前一个，即dp[j][k] dp[j-][k-1]有关系
			if nums2[j] == nums1[k] {
				dp718[j+1][k+1] = dp718[j][k] + 1
				if max718 < dp718[j+1][k+1] {
					max718 = dp718[j+1][k+1]
				}
			} else {
				dp718[j+1][k+1] = 0
			}
		}
	}
	return max718
}

func TestAl718(t *testing.T) {
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
	//fmt.Println(findLength([]int{3, 2, 1}, []int{3, 2, 1}))
	fmt.Println(findLength([]int{0, 0, 0, 0}, []int{0, 0, 0, 0}))
}
