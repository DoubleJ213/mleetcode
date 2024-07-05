package fxxk

import (
	"fmt"
	"testing"
)

/*
如果下标对 i、j 满足 0 ≤ i < j < nums.length ，
如果 nums[i] 的 第一个数字 和 nums[j] 的 最后一个数字 互质 ，
则认为 nums[i] 和 nums[j] 是一组 美丽下标对 。
i 在 j 左边 i第一个数字，j最后一个数字
nums[i] % 10 != 0 尾数不为0，首位也不能为0呀
所以本题很特殊，只是2个一位数比较。1~9 所以质数对很有限，枚举就好
除了1 1 互质，其余相同的数都不互质
*/
// leetcode submit region begin(Prohibit modification and deletion)
func countBeautifulPairs(nums []int) int {
	var ans int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			left := getLeft2748(nums[i])
			right := nums[j] % 10
			if isGcd2748(left, right) {
				fmt.Println(left, right)
				ans++
			}
		}
	}
	return ans
}

// 判断a、b两个数是否互质, 确保计算时 a大  b小 入参反过来时，多走一次循环
func isGcd2748(a int, b int) bool {
	for b != 0 {
		a, b = b, a%b
	}
	if a == 1 {
		return true
	}
	return false
}

func getLeft2748(i int) int {
	for i >= 10 {
		i = i / 10
	}
	return i
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNumberOfBeautifulPairs(t *testing.T) {
	//fmt.Println(countBeautifulPairs([]int{9, 3}))
	//isGcd2748(64, 35)
	isGcd2748(35, 64)
	//isGcd2748(37*2, 37)
}
