package fxxk

import (
	"fmt"
	"math"
	"testing"
)

/*
假定每组输入只存在恰好一个解
所有三数和的组合都列出来，然后排序，二分法找个离得最近的？
*/
//leetcode submit region begin(Prohibit modification and deletion)

// ans 当前的和 dis 和target的差值
var ans16 int
var dis16 float64

func threeSumClosest(nums []int, target int) int {
	ans16 = nums[0] + nums[1] + nums[2] //初始化
	dis16 = math.Abs(float64(target - ans16))
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				dis := math.Abs(float64(target - sum))
				if dis < dis16 {
					if target == sum {
						return sum
					}
					ans16 = sum
					dis16 = dis
				}
			}
		}
	}
	return ans16
}

//leetcode submit region end(Prohibit modification and deletion)

func TestThreeSumClosest(t *testing.T) {
	//fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{40, -53, 36, 89, -38, -51, 80, 11, -10, 76, -30, 46, -39, -15, 4, 72, 83, -25, 33, -69, -73, -100, -23, -37, -13, -62, -26, -54, 36, -84, -65, -51, 11, 98, -21, 49, 51, 78, -58, -40, 95, -81, 41, -17, -70, 83, -88, -14, -75, -10, -44, -21, 6, 68, -81, -1, 41, -61, -82, -24, 45, 19, 6, -98, 11, 9, -66, 50, -97, -2, 58, 17, 51, -13, 88, -16, -77, 31, 35, 98, -2, 0, -70, 6, -34, -8, 78, 22, -1, -93, -39, -88, -77, -65, 80, 91, 35, -15, 7, -37, -96, 65, 3, 33, -22, 60, 1, 76, -32, 22}, 292))

}
