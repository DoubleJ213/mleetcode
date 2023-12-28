package fxxk

import (
	"fmt"
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

/*
刷之前可以看下509
第一种，暴力递归
第二种，优化暴力法，通过添加数组存子问题的结果简化部分子树的计算逻辑，简称带备忘录的递归
第三种，dp 数组的迭代解法

重叠子问题、最优子结构、状态转移方程就是动态规划三要素
明确 base case -> 明确「状态」-> 明确「选择」 -> 定义 dp 数组/函数的含义。
base case 零钱0的时候最小需要0个硬币
状态是 零钱数
选择是 每次增加了某个面值的硬币
明确 dp 函数/数组的定义 子问题的处理办法
*/

func coinChange0(coins []int, amount int) int {
	//暴力递归
	if len(coins) == 0 || amount == 0 {
		return 0
	}
	return dp322(coins, amount)
}

/*
函数签名和之前一样，可以考虑不要单独拎出来一个方法
为了方便套公式，都这个模式写
*/
func dp322(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	result322 := 10001
	//因为amount最大10^4，硬币最小1，所以这里取个10^4+1，理论上是不可能取到这个值的
	//如果每个子问题都找最小值，都和105比较，肯定小于10^4+1
	//fuck 从leetcode复制过来的数据有问题，ide里面显示为104 其实是10^4
	//害我提交错误一次，日
	for _, coin := range coins {
		subProblem322 := dp322(coins, amount-coin)
		if subProblem322 == -1 {
			continue
		}
		result322 = getMin322(result322, subProblem322+1)
	}
	if result322 == 10001 {
		return -1
	}
	return result322
}

func getMin322(x int, y int) int {
	if x >= y {
		return y
	}
	return x
}

func TestAl322(t *testing.T) {
	fmt.Printf("3 --- %v\n", coinChange0([]int{1, 2, 5}, 11))
	fmt.Printf("-1 --- %v\n", coinChange2([]int{2}, 3))
	fmt.Printf("暴力算法超时了 --- %v\n", coinChange2([]int{1, 2, 5}, 100))
}

var minRes []int

func coinChange1(coins []int, amount int) int {
	//优化暴力递归，
	//增加一个数组，存子问题的最优解，减少重复的计算
	//称呼为 带备忘录的递归
	if len(coins) == 0 || amount == 0 {
		return 0
	}
	minRes = make([]int, amount+1)
	//整一个数组---备忘录，里面存针对amount为amount-1 amount-2 等等的情况的最优解
	//下次递归到此处直接查询，不要递归
	return dp322_1(coins, amount)
}

//func dp322_1(coins []int, amount int) int {
//	if amount == 0 {
//		minRes[0] = 0
//		return 0
//	}
//	if amount < 0 {
//		return -1
//	}
//	result := 105
//	for _, coin := range coins {
//		var subProblem int
//		if amount-coin > 0 && minRes[amount-coin] != 0 {
//			subProblem = minRes[amount-coin]
//		} else {
//			subProblem = dp322_1(coins, amount-coin)
//		}
//		if subProblem == -1 {
//			continue
//		}
//		result = getMin(result, subProblem+1)
//	}
//
//	if result == 105 {
//		minRes[amount] = -1
//		return -1
//	} else {
//		minRes[amount] = result
//		return result
//	}
//
//}

func dp322_1(coins []int, amount int) int {
	if amount == 0 {
		minRes[0] = 0
		return 0
	}
	if amount < 0 {
		return -1
	}

	if minRes[amount] != 0 {
		//这里提前return比在for循环里面return应该又快一些
		//反正目的就是已知结果的情况下 少递归
		return minRes[amount]
	}

	result := 10001
	for _, coin := range coins {
		subProblem := dp322_1(coins, amount-coin)
		if subProblem == -1 {
			continue
		}
		result = getMin(result, subProblem+1)
	}

	if result == 10001 {
		minRes[amount] = -1
		return -1
	} else {
		minRes[amount] = result
		return result
	}
}

//自底向上使用 dp table 来消除重叠子问题
//类似fib的for循环那种，自底向上求每个数组的值，然后直接通过查备忘录得出结果
func coinChange2(coins []int, amount int) int {
	if len(coins) == 0 || amount == 0 {
		return 0
	}
	minRes = make([]int, amount+1)
	minRes[0] = 0
	//dp 数组的定义：当目标金额为 i 时，至少需要 dp[i] 枚硬币凑出
	//外层 for 循环在遍历所有状态的所有取值
	for i := 1; i <= amount; i++ {
		// 内层 for 循环在求所有选择的最小值
		minRes[i] = 10001
		//把结果改成一个不可能出现的值，取最小值，最小值肯定小于这个值，当值等于这个值的时候证明最终结果应该是-1
		for _, coin := range coins {
			// 子问题无解，跳过
			if minRes[i] == 1 {
				//minRes[i] == 1
				//通过1个硬币就能得到零钱数已经是最优解了，直接不用判断其余硬币凑这个零钱了
				break
			}
			if i-coin < 0 {
				continue
			}
			//每次循环，不断的更新目标金额为i 最小需要minRes[i]的数组 备忘录
			minRes[i] = getMin322(minRes[i], 1+minRes[i-coin])
			//1+minRes[i-coin]
			//当前使用零钱为coin， minRes[总零钱i-当前使用的零钱coin]
			//比如总零钱为2
			//循环 使用1时 那就是minRes[2-1]+1
			//使用2时 minRes[2-2]+1
			//使用5时 minRes[2-5] +1
			//以上三种所有可能性的最小值为 总零钱2最优解
		}
	}

	if minRes[amount] == 10001 {
		return -1
	} else {
		return minRes[amount]
	}
}
