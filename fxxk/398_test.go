package fxxk

import (
	"fmt"
	"math/rand"
	"testing"
)

/*
398. 随机数索引

给你一个可能含有 重复元素 的整数数组 nums ，请你随机输出给定的目标数字 target 的索引。
你可以假设给定的数字一定存在于数组中。实现 Solution 类：
Solution(int[] nums) 用数组 nums 初始化对象。
int pick(int target) 从 nums 中选出一个满足 nums[i] == target 的随机索引 i 。
如果存在多个有效的索引，则每个索引的返回概率应当相等。

示例：
输入
["Solution", "pick", "pick", "pick"]
[[[1, 2, 3, 3, 3]], [3], [1], [3]]
输出
[null, 4, 0, 2]
解释
Solution solution = new Solution([1, 2, 3, 3, 3]);
solution.pick(3); // 随机返回索引 2, 3 或者 4 之一。每个索引的返回概率应该相等。
solution.pick(1); // 返回 0 。因为只有 nums[0] 等于 1 。
solution.pick(3); // 随机返回索引 2, 3 或者 4 之一。每个索引的返回概率应该相等。
提示：
1 <= nums.length <= 2 * 10^4
-2^31 <= nums[i] <= 2^31 - 1
target 是 nums 中的一个整数
最多调用 pick 函数 10^4 次
*/

type Solution struct {
	value map[int][]int
}

func Constructor(nums []int) Solution {
	va := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		if num, ok := va[cur]; ok {
			num = append(num, i)
			va[cur] = num
		} else {
			va[cur] = []int{i}
		}
	}
	so := Solution{value: va}
	return so
}

func (this *Solution) Pick(target int) int {
	if array, ok := this.value[target]; ok {
		//array
		return array[rand.Intn(len(array))]
	}
	return -1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
func TestAl398(t *testing.T) {
	obj := Constructor([]int{1, 2, 3, 3, 3, 2, 1, 4})
	fmt.Println(obj.Pick(3))
	fmt.Println(obj.Pick(3))
	fmt.Println(obj.Pick(3))
	fmt.Println(obj.Pick(3))
	fmt.Println(obj.Pick(3))
	fmt.Println(obj.Pick(3))
	fmt.Println(obj.Pick(3))
}
