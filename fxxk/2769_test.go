package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 每次操作将 x 的值增加或减少 1 ，同时可以选择将 num 的值增加或减少 1
// 返回所有可达成数字中的最大值。可以证明至少存在一个可达成数字。
// dp题吧，x +1/-1  num +1/-1  执行下述操作不超过 t 次
func theMaximumAchievableX(num int, t int) int {
	//x +1/-1 同时 num +1/-1 对应影响几种最终值 +1 +1、 +1 -1 、-1 +1、 -1 -1 即 +2、 +0、 -2
	//x 不动，那num就是 每次就是 +2、 +0、 -2 三种操作
	//所有可达成数字中的最大值
	return num + t*2
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindTheMaximumAchievableNumber(t *testing.T) {
	fmt.Println(theMaximumAchievableX(4, 1))
}
