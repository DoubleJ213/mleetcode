package fxxk

import (
	"fmt"
	"testing"
)

/*
上升 \ 平稳 \ 下降 三种趋势
返回两地气温变化趋势相同的最大连续天数。
*/
//leetcode submit region begin(Prohibit modification and deletion)
func temperatureTrend(temperatureA []int, temperatureB []int) int {
	ans, i, j, size := 0, 1, 1, 0
	for i < len(temperatureA) && j < len(temperatureB) {
		preI, postI, preJ, postJ := temperatureA[i-1], temperatureA[i], temperatureB[j-1], temperatureB[j]
		if preI > postI && preJ > postJ || preI == postI && preJ == postJ || preI < postI && preJ < postJ {
			size++
		} else {
			size = 0
		}
		i++
		j++
		ans = getMax(size, ans)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSixCE719(t *testing.T) {
	temperatureA := []int{21, 18, 18, 18, 31}
	temperatureB := []int{34, 32, 16, 16, 17}
	fmt.Println(temperatureTrend(temperatureA, temperatureB))
}
