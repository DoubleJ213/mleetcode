package review

import (
	"fmt"
	"math"
	"testing"
)

/*
322. 零钱兑换
给你一个整数数组coins，表示不同面额的硬币；以及一个整数amount，表示总金额。
计算并返回可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回-1。
你可以认为每种硬币的数量是无限的。

示例 1：
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：
输入：coins = [2], amount = 3
输出：-1
示例 3：
输入：coins = [1], amount = 0
输出：0

提示：
1 <= coins.length <= 12
1 <= coins[i] <= 2^31 - 1
0 <= amount <= 10^4
*/

var dpr322 []int

/*
coins 每种硬币的数量是无限的 求可以凑成总金额 amount 所需的最少的硬币个数
比如输入为 ([]int{1, 2, 5}, 11))
凑amount=11
定义一个dp[i] 其表示什么东西呢，第一种要求什么值，我就定义数组存什么值
dp[i]表示 凑零钱i 最少多少个硬币他和前面dp[i-1] 数的关系是啥
比如用一个5面值的硬币 dp[i] = dp[i-5] +1
比如用一个2面值的硬币 dp[i] = dp[i-2] +1
dp[i] 又在一次次尝试中取最小值
*/
func coinChange4(coins []int, amount int) int {
	dpr322 = make([]int, amount+1)
	//dpr322[0] 表示凑到零钱和为0的时候，那至少0个硬币
	dpr322[0] = 0

	for i := 1; i < amount+1; i++ {
		dpr322[i] = int(math.Pow(10, 4)) + 1
		for _, coin := range coins {
			if i-coin >= 0 {
				dpr322[i] = getMin(dpr322[i], dpr322[i-coin]+1)
			}
		}
	}
	//因为一直是在整最小的，如果初始化出来的数组每个值不修改，那岂不全是0最小。得把数组的每个值变成最大的

	//n := len(coins)
	if dpr322[amount] == 10001 {
		return -1
	}
	return dpr322[amount]
}

func TestAl322_1(t *testing.T) {
	fmt.Printf("3 --- %v\n", coinChange4([]int{1, 2, 5}, 11))
	fmt.Printf("-1 --- %v\n", coinChange4([]int{2}, 3))
	fmt.Printf("暴力算法超时了 --- %v\n", coinChange4([]int{1, 2, 5}, 100))
}
