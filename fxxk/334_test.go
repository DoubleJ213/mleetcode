package fxxk

import (
	"math"
	"testing"
)

//334. 递增的三元子序列
//给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
//如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。
//
//示例 1：
//输入：nums = [1,2,3,4,5] 输出：true
//解释：任何 i < j < k 的三元组都满足题意
//示例 2：
//输入：nums = [5,4,3,2,1] 输出：false
//解释：不存在满足题意的三元组
//示例 3：
//输入：nums = [2,1,5,0,4,6] 输出：true
//解释：三元组 (3, 4, 5) 满足题意，因为 nums[3] == 0 < nums[4] == 4 < nums[5] == 6
//
//提示：
//1 <= nums.length <= 5 * 105
//-231 <= nums[i] <= 231 - 1
//进阶：你能实现时间复杂度为 O(n) ，空间复杂度为 O(1) 的解决方案吗？

//贪心
//我们可以对f数组进行优化：使用有限变量进行替换（将f数组的长度压缩为2），
//数组含义不变，f[1]=x代表长度为1的上升子序列最小结尾元素为x，
//f[2]=y代表长度为2的上升子序列的最小结尾元素为y。
//从前往后扫描每个nums[i]，每次将nums[i]分别与f[1]和f[2]进行比较，
//如果能够满足nums[i]>f[2]，代表nums[i]能够接在长度为2的上升子序列的后面，直接返回True，
//否则尝试使用nums[i]来更新f[1]和f[2]。
//这样，我们只使用了有限变量，每次处理nums[i]只需要和有限变量进行比较，
//时间复杂度为O(n)，空间复杂度为O(1)

func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt32, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] <= first {
			first = nums[i]
		} else if nums[i] <= second {
			second = nums[i]
		} else {
			// first 对应small、second对应 mid
			//比如2, 3, 1, 4这个数组。
			//1的出现会替换掉2，3是在1之前，现在small和mid顺序不对了，
			//但是1替换掉2也不用紧，原因2这个曾经的small，肯定是小于mid，只有大于small才会被放到mid。
			//即有一个比 small 大比 mid 小的前·最小值出现在 mid 之前。
			//补充，比如small=1，mid=3，4能返回true，并不是因为1，3，4是满足的。
			//而是因为2，3，4满足，2肯定是小于3的，也是在4前面出现的，其实也是3之前出现的。
			//这样写的缺陷是，只能判断是否存在，没有把匹配的数组找出来。
			return true
		}
	}
	return false
}

func TestAl334(t *testing.T) {
	//3 4 1 5 2
	increasingTriplet([]int{2, 3, 1, 4})
}
