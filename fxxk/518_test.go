package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

/*
518. 零钱兑换 II
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
假设每一种面额的硬币有无限个。
题目数据保证结果符合 32 位带符号整数。

示例 1：
输入：amount = 5, coins = [1, 2, 5]
输出：4
解释：有四种方式可以凑成总金额：
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
示例 2：
输入：amount = 3, coins = [2]
输出：0
解释：只用面额 2 的硬币不能凑成总金额 3 。
示例 3：
输入：amount = 10, coins = [10]
输出：1

提示：
1 <= coins.length <= 300
1 <= coins[i] <= 5000
coins 中的所有值 互不相同
0 <= amount <= 5000
*/

/*
类似题目322
这是个dp问题，是可以画出递归树的，先写暴力解法，可能会超时
*/
var result518 int

func change(amount int, coins []int) int {
	result518 = 0
	dp518(amount, coins)
	return result518
}

func dp518(amount int, coins []int) {
	if amount == 0 {
		result518++
		return
	}
	if amount < 0 {
		return
	}
	for _, coin := range coins {
		dp518(amount-coin, coins)
	}
	return
}

/*
change 存在重复的情况 1112 1121 1211 2111 这四种在本题看来是一种情况其实
要去掉这种重复的情况，需要在coins中做手脚，比如使用当前coin的时候amount已经小于0了
就不要再后续的循环中出现coin了？
但是有221 的情况存在，针对1112 上面的也暴力的不让不出现2吧
先排序，从大的往小和amount减
写了试一试
*/
func change1(amount int, coins []int) int {
	result518 = 0
	sort.Ints(coins)
	dp518_1(amount, coins)
	return result518
}

func dp518_1(amount int, coins []int) {
	if amount == 0 {
		result518++
		return
	}
	if amount < 0 {
		return
	}
	for i := len(coins) - 1; i >= 0; i-- {
		curr := coins[i]
		if amount < curr {
			dp518_1(amount-curr, coins[:i])
		}
		dp518_1(amount-curr, coins[:i+1])
	}
}

/*
change1 提交发现超时了，不过也是递归的老问题了
change2 加上个备忘录。
但是与以往递归不同的是，为了避免 1121 1112 等重复的情况，按照 change1 的写法本题选择的条件一直在变化
这里其实有两个条件,需要一个二维数组 dp[][]
一个状态条件 零钱数 i 一个状态 是所使用的硬币 j   dp[i][j]
表示含义 使用前i个硬币，凑面值为j的零钱数
base case

dp[0][...] = 0   代表不使⽤任何硬币⾯值（使用coins中 前0种钱币），这种情况下显然⽆ 法凑出任何⾦额
dp[...][0] = 1  代表需要凑出的⽬标⾦额为 0，那么什么都不做就是唯⼀的⼀种凑法

我们最终想得到的答案就是 dp[N][amount]，其中 N 为 coins 数组的⼤⼩。
伪代码
int dp[N+1][amount+1]
dp[0][..] = 0
dp[..][0] = 1

for i in [1..N]:
for j in [1..amount]:

	把物品 i 装进背包, 不把物品 i 装进背包

return dp[N][amount]

根据「选择」，思考状态转移的逻辑
如果你不把这第 i 个物品装⼊背包，也就是说你不使⽤ coins[i-1] 这个⾯值的硬币，
那么凑出⾯额 j 的⽅ 法数 dp[i][j] 应该等于 dp[i-1][j]，继承之前的结果。

如果你把这第 i 个物品装⼊了背包，也就是说你使⽤ coins[i-1] 这个⾯值的硬币，
那么 dp[i][j] 应该 等于 dp[i][j-coins[i-1]]。

if (j - coins[i-1] >= 0)

	dp[i][j] = dp[i - 1][j] + dp[i][j-coins[i-1]]
*/
func change2(amount int, coins []int) int {
	/*
		int dp[N+1][amount+1]
		dp[0][..] = 0
		dp[..][0] = 1
	*/
	//int dp[N+1][amount+1]
	amountRes := make([][]int, len(coins)+1)
	//dp[0][..] = 0
	for i := 0; i < len(coins)+1; i++ {
		amountRes[i] = make([]int, amount+1)
		//dp[..][0] = 1
		amountRes[i][0] = 1
	}
	//amountRes[0][0] = 1
	for i := 1; i < len(coins)+1; i++ {
		for j := 1; j < amount+1; j++ {
			//当 零钱大于等于硬币的面值，能用零钱减去当前的硬币
			if j >= coins[i-1] {
				//amountRes[i-1][j] 其实是前一个 amountRes[i][j]的结果
				amountRes[i][j] = amountRes[i-1][j] + amountRes[i][j-coins[i-1]]
			} else {
				amountRes[i][j] = amountRes[i-1][j]
			}
		}
	}
	return amountRes[len(coins)][amount]
}

/*
仔细想了想
因为dp[i][...] 只和dp[i][...] dp[i-1][...]相关
所以自底向上的dp table的写法也能简化成一维数组
*/
func change3(amount int, coins []int) int {
	res518 := make([]int, amount+1)
	//记录全部凑面值为0的零钱的方法
	//凑面值1的方法，就是凑面值为0
	res518[0] = 1
	for i := 0; i < len(coins); i++ {
		coin := coins[i]
		for j := 1; j < amount+1; j++ {
			if j >= coin {
				//零钱大于硬币面值，能用当前的硬币去减
				res518[j] += res518[j-coin]
			}
		}
	}
	return res518[amount]
}

func TestAl518(t *testing.T) {
	/*
		fmt.Printf("4 --- %v\n", change(5, []int{1, 2, 5}))
		//change 算出来的结果是有问题的
		//2111 1211 1121 1112 122 212 221 5 11111
		//这9种全部被算进结果了。其实有5个是重复的，正确答案是4不是9
	*/
	//fmt.Printf("4 --- %v\n", change1(5, []int{5, 2, 1}))
	//fmt.Printf("0 --- %v\n", change1(3, []int{2}))
	//fmt.Printf("1 --- %v\n", change1(10, []int{10}))
	//fmt.Printf("暴力算法超时了 --- %v\n", change1(500, []int{3, 5, 7, 8, 9, 10, 11}))
	fmt.Printf("4 --- %v\n", change2(5, []int{1, 2, 5}))
	//fmt.Printf("35502874 --- %v\n", change2(500, []int{3, 5, 7, 8, 9, 10, 11}))
	fmt.Printf("4 --- %v\n", change3(5, []int{1, 2, 5}))
	//fmt.Printf("35502874 --- %v\n", change3(500, []int{3, 5, 7, 8, 9, 10, 11}))
	//fmt.Printf("-1 --- %v\n", change(3, []int{2}))
	//fmt.Printf("暴力算法超时了 --- %v\n", change(100, []int{1, 2, 5}))
}
