package fxxk

import (
	"fmt"
	"math"
	"testing"
)

/*
'nums' must consist of values from 1 to 1000000 only
因为数字都是大于1的，所以只要相同，直接无脑加上就行
重点来了，如果不相同，怎么处理？
当当前数大于x也可以加上，顺便就改了后续求值的可能性，当当前数小于x呢
*/

// leetcode submit region begin(Prohibit modification and deletion)

func maxScore2(nums []int, x int) int64 {
	dp := make([]int64, len(nums))
	oddIndex, evenIndex := -1, -1 //奇数最大坐标 偶数最大坐标
	dp[0] = int64(nums[0])
	ans := dp[0]
	if nums[0]&1 != 1 { //！= 1 表示偶数
		evenIndex = 0
	} else {
		oddIndex = 0
	}

	for i := 1; i < len(nums); i++ {
		if nums[i]&1 != 1 { //偶数
			if evenIndex != -1 { //前面有偶数
				dp[i] = dp[evenIndex] + int64(nums[i])
			} else {
				// 前面没有偶数，前面又有数字，所以前面一定是有奇数的
				dp[i] = dp[oddIndex] + int64(nums[i]-x)
			}
			if oddIndex != -1 { //前面有奇数
				dp[i] = getMax64(dp[i], dp[oddIndex]+int64(nums[i]-x))
			}
		} else { //奇数
			if oddIndex != -1 { //前面有奇数
				dp[i] = dp[oddIndex] + int64(nums[i])
			} else {
				// 前面没有奇数，前面又有数字，所以前面一定是有偶数的
				dp[i] = dp[evenIndex] + int64(nums[i]-x)
			}
			if evenIndex != -1 { //前面有偶数
				dp[i] = getMax64(dp[i], dp[evenIndex]+int64(nums[i]-x))
			}
		}
		//最后更新下坐标，方便后面数计算
		if nums[i]&1 != 1 { //！= 1 表示偶数
			evenIndex = i
		} else {
			oddIndex = i
		}
		ans = getMax64(dp[i], ans)
	}
	return ans
}

func getMax64(x, y int64) int64 {
	if x >= y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)

func TestVisitArrayPositionsToMaximizeScore(t *testing.T) {
	//整数数组 nums 和一个正整数 x
	//[85,12] 79
	fmt.Println(maxScore([]int{79, 3, 6, 26, 72, 53, 78, 45, 34, 20}, 97))
	//输出 230...  预计 221...
	//fmt.Println(maxScore([]int{2, 3, 6, 1, 9, 2}, 5))
	//fmt.Println(maxScore([]int{85, 12}, 79))
}

// 进一步搞一个数组 max max := []int{奇数最大值，偶数最大值} max[0] 奇数最大值 max[1] 偶数最大值 最后就是求max[0] max[1]两个最大值
func maxScore(nums []int, x int) int64 {
	max := [2]int{math.MinInt + 1000000, math.MinInt + 1000000} //max[v&1^1]-x 防止溢出，不能定义成最小的数 max[v&1^1]-x
	max[nums[0]&1] = nums[0]                                    //这里为了方便代码编写，与等于0 偶数 与等于1奇数 ，前面偶数 后面奇数
	for _, v := range nums[1:] {
		max[v&1] = getMax(max[v&1], max[v&1^1]-x) + v
		//v&1^1	0变1，1变0。就是奇偶的转换
	}
	return int64(getMax(max[0], max[1]))
}
