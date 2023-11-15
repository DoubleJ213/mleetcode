package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

/*
LCR 103. 零钱兑换
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
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
示例 4：
输入：coins = [1], amount = 1
输出：1
示例 5：
输入：coins = [1], amount = 2
输出：2

提示：
1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 10^4
*/

var resLrc103 []int

func coinChange(coins []int, amount int) int {
	resLrc103 = make([]int, amount+1)
	/*
		暴力的，先求出所有的组合数，然后再求最短的情况，不写了
	*/
	//resLrc103[amount] 表示 总额amount的时候需要多少
	//resLrc103[amount] = resLrc103[amount-coin] +1 和 resLrc103[amount] 的最小值
	sort.Ints(coins)

	resLrc103[0] = 0

	for i := 1; i < amount+1; i++ {
		resLrc103[i] = 10005
		for j := 0; j < len(coins); j++ {
			coin := coins[j]
			if i < coin {
				break
			}
			resLrc103[i] = getMin(resLrc103[i], resLrc103[i-coin]+1)
		}
	}
	if resLrc103[amount] == 10005 {
		return -1
	}
	return resLrc103[amount]
}

func TestLrc103(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	//[10,55,45,25,25]
}
