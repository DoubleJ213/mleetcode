package review

import (
	"fmt"
	"sort"
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
1 <= coins[i] <= 231 - 1
0 <= amount <= 10^4
*/

var len322 int
var path322 []int

//暴力递归
func coinChange(coins []int, amount int) int {
	len322 = -1
	path322 = make([]int, 0)
	dp322(coins, path322, amount)
	//计算最短的path
	return len322
}

func dp322(coins, path []int, amount int) {
	if amount < 0 {
		return
	}
	if amount == 0 {
		//path322
		length := len(path)
		if len322 == -1 || length < len322 {
			len322 = length
		}
		return
	}
	for i := 0; i < len(coins); i++ {
		coin := coins[i]
		path = append(path, coin)
		dp322(coins, path, amount-coin)
		path = path[:len(path)-1]
	}
	//	amount为9时 比如减了两次1 和减了1次2
	//	其递归过程中存在n多的重复计算，后续可以优化
}

func getMin(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}

//dp[n]=dp[n-coin] + 1
//假如找到dp[n-coin]的最优解，dp[n] 的最优解就是 再凑一个面值为coin的硬币就完事
func coinChange2(coins []int, amount int) int {
	if len(coins) == 0 || amount == 0 {
		return 0
	}
	return dp322_a(coins, amount)
}

func dp322_a(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res322 := 10001
	for i := 0; i < len(coins); i++ {
		coin := coins[i]
		subProblem := dp322_a(coins, amount-coin)
		if subProblem == -1 {
			continue
		}
		res322 = getMin(res322, subProblem+1)
	}

	if res322 == 10001 {
		return -1
	}

	return res322
}

var res322 []int

//自底向上，最优子问题，明确dp table的含义
//res322[n] 表示凑 总数为n的需要至少 多少个硬币
func coinChange3(coins []int, amount int) int {
	if len(coins) == 0 || amount == 0 {
		return 0
	}
	sort.Ints(coins)
	res322 = make([]int, amount+1)
	res322[0] = 0
	//凑面值0 至少0个硬币
	for i := 1; i <= amount; i++ {
		res322[i] = 10001
		for j := 0; j < len(coins); j++ {
			coin := coins[j]
			if i < coin {
				break
			}
			res322[i] = getMin(res322[i], res322[i-coin]+1)
		}
	}
	if res322[amount] == 10001 {
		return -1
	}
	return res322[amount]
}

func TestAl322(t *testing.T) {
	fmt.Printf("3 --- %v\n", coinChange2([]int{1, 2, 5}, 11))
	fmt.Printf("-1 --- %v\n", coinChange2([]int{2}, 3))
	fmt.Printf("暴力算法超时了 --- %v\n", coinChange3([]int{1, 2, 5}, 100))
}
