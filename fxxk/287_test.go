package fxxk

import (
	"fmt"
	"testing"
)

/*
287. 寻找重复数
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。

示例 1：
输入：nums = [1,3,4,2,2]
输出：2
示例 2：
输入：nums = [3,1,3,4,2]
输出：3

提示：
1 <= n <= 10^5
nums.length == n + 1
1 <= nums[i] <= n
nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次

进阶：
如何证明 nums 中至少存在一个重复的数字?
你可以设计一个线性级时间复杂度 O(n) 的解决方案吗？
*/

func findDuplicate1(nums []int) int {
	//nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次
	//把具体数字放到对应位置，如果对应位置已经有了这个数，那就是有冲突了。那就找出来了

	//数组  [1, n]  index 本来从0 ~ len(nums)-1

	for i := 0; i < len(nums); i++ {
		//比如current=1 应该放在 index=0 的位置 即 index=current-1  被交换的数nums[current-1]
		//比如current-1 就是当前 index 就是说这个数位置没问题，不要交换
		//如果 需要交换，但是需要交换的数和当前数值一样，那就找到了重复的数字，后面就不要找了
		current := nums[i]
		if i != current-1 {
			toSwap := nums[current-1]
			//	需要交换
			if toSwap == current {
				return current
			}
			nums[i], nums[current-1] = nums[current-1], nums[i]
			i--
		}
		//swap287
	}

	return -1
}

func findDuplicate(nums []int) int {
	var pre int
	for {
		//index为0的数一直交换，比如index为0对应的数为1 将 数1 放到nums[1]去，之前nums[1]的数放到 index=0来
		//一直交换，直到交换过来的 新的数和 刚刚的nums[0] 即 pre 相等

		//或者反过来，如果现在nums[0] 和要放到目标位置的数一样，就说明不能取代人家，重复了
		pre = nums[0]
		if nums[nums[0]] == pre {
			return pre
		}
		nums[0], nums[nums[0]] = nums[nums[0]], nums[0]

	}
}

func TestAl287(t *testing.T) {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 3, 2}))
}
