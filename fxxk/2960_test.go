package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
/*
返回一个整数，表示按顺序执行测试操作后 已测试设备 的数量。
*/
func countTestedDevices(batteryPercentages []int) int {
	var ans int
	if batteryPercentages == nil || len(batteryPercentages) == 0 {
		return ans
	}
	//var toggle int //表示当前数已经被减过几了
	//结果发现 toggle 是和ans一个值的。那我直接一个变量完事
	for i := 0; i < len(batteryPercentages); i++ {
		if batteryPercentages[i] > ans {
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCountTestedDevicesAfterTestOperations(t *testing.T) {
	fmt.Println(countTestedDevices([]int{1, 1, 2, 1, 3}))
}
