package fxxk

import (
	"fmt"
	"testing"
)

/*
1011. 在 D 天内送达包裹的能力

传送带上的包裹必须在 days 天内从一个港口运送到另一个港口。
传送带上的第 i个包裹的重量为weights[i]。每一天，我们都会按给出重量（weights）的顺序往传送带上装载包裹。
我们装载的重量不会超过船的最大运载重量。
返回能在 days 天内将传送带上的所有包裹送达的船的最低运载能力。

示例 1：
输入：weights = [1,2,3,4,5,6,7,8,9,10], days = 5
输出：15
解释：
船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
第 1 天：1, 2, 3, 4, 5
第 2 天：6, 7
第 3 天：8
第 4 天：9
第 5 天：10

请注意，货物必须按照给定的顺序装运，
因此使用载重能力为 14 的船舶并将包装分成 (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) 是不允许的。
示例 2：
输入：weights = [3,2,2,4,1,4], days = 3
输出：6
解释：
船舶最低载重 6 就能够在 3 天内送达所有包裹，如下所示：
第 1 天：3, 2
第 2 天：2, 4
第 3 天：1, 4
示例 3：
输入：weights = [1,2,3,1,1], days = 4
输出：3
解释：
第 1 天：1
第 2 天：2
第 3 天：3
第 4 天：1, 1

提示：
1 <= days <= weights.length <= 5 * 104
1 <= weights[i] <= 500
*/

/*
至少有一天承载了一天最重的包裹，所以最小值是max(weights)
left, right = max(weights), sum(weights) + 1

然后运载能力就在这个区间内，通过二分法判断
要么运力不够，天数超过了，需要增加运力
要么运力嫌多，天数太小了，需要减运力

再结合找这个满足要求的最小运力，找左边界
*/
func shipWithinDays(weights []int, days int) int {
	ship := func(cap int) int {
		needDays := 0
		remainCap := 0
		for _, weight := range weights {
			if remainCap < weight {
				needDays++
				remainCap = cap
			}
			remainCap -= weight
		}
		return needDays
	}
	var left, right int
	for _, weight := range weights {
		if left < weight {
			left = weight
		}
		right += weight
	}
	//	都是闭区间
	for left <= right {
		mid := left + (right-left)/2
		//fmt.Printf("ship(mid) %d ", ship(mid))
		if ship(mid) == days {
			right = mid - 1
		} else if ship(mid) < days {
			right = mid - 1
		} else if ship(mid) > days {
			left = mid + 1
		}
	}
	return left
}

func TestAl1011(t *testing.T) {
	//fmt.Printf("%d \n", shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
	fmt.Printf("%d \n", shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
}
