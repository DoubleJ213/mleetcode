package fxxk

import (
	"fmt"
	"testing"
)

/*
303. 区域和检索 - 数组不可变

给定一个整数数组 nums，处理以下类型的多个查询:
计算索引left和right（包含 left 和 right）之间的 nums 元素的 和 ，其中left <= right
实现 NumArray 类：
NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int i, int j) 返回数组 nums中索引left和right之间的元素的 总和 ，包含left和right两点（也就是nums[left] + nums[left + 1] + ... + nums[right])

示例 1：
输入：
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
输出：
[null, 1, -1, -3]
解释：
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))
numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))

*/

type NumArray struct {
	length int
	nums   []int
}

func Constructor303(nums []int) *NumArray {
	return &NumArray{len(nums), nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > right || right > this.length-1 || left < 0 {
		return 0
	}
	resp := 0
	for i := left; i <= right; i++ {
		resp += this.nums[i]
	}
	return resp
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func TestAl303(t *testing.T) {
	nums := []int{3, 5, 2, -2, 4, 1}
	obj := Constructor303(nums)
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))

	obj1 := Constructor303_1(nums)
	fmt.Println(obj1.SumRange1(0, 2))
	fmt.Println(obj1.SumRange1(2, 5))
	fmt.Println(obj1.SumRange1(0, 5))
}

/*
核心思路是我们new一个新的数组preSum出来，preSum[i]记录nums[0..i-1]的累加和
这是别人的思路

我修改了，暂时不确定好坏，但是好理解
sum[i] 就是 num[0~i]的值的累加

nums   3  5  2  -2  4  1
sum    3  8  10  8  12 13

求nums 1~2 的值 7
1 的值 sum[1]-sum[1-1]
2 的值 sum[2]-sum[2-1]
1~2的值 sum[1]-sum[1-1] + sum[2]-sum[2-1]

约去sum[1],即sum[2]-sum[1-1] 即 sum[r]-sum[l-1]

这样一整，你会发现当l为0的时候第二个sum数组出来index [-1]了
当然处理好边界也行。但是还是别人思路是更合适的

让Sum[i]记录nums[0..i-1]的累加和

nums   3  5  2  -2  4  1
sum  0 3  8  10  8  12 13

比如算nums[1~2]的值 5+2=7
sum[2+1]-sum[1]
sum[r+1]-sum[l]

*/

type NumArray1 struct {
	preSums []int
}

func Constructor303_1(nums []int) NumArray1 {
	//sum[i] 为 nums[0~i-1] 累加值
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += nums[j]
		}
		preSum[i] = sum
	}

	return NumArray1{preSum}
}

func (this *NumArray1) SumRange1(i, j int) int {
	//sum[r+1]-sum[l]
	return this.preSums[j+1] - this.preSums[i]
}
