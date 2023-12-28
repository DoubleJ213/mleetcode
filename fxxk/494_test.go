package fxxk

import (
	"fmt"
	"testing"
)

/*
494. 目标和
给你一个非负整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

示例 1：
输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
示例 2：
输入：nums = [1], target = 1
输出：1

提示：
1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000
*/

/*
每个数前面，要么是+ 要么是-
因为表达式需要 串联起所有整数,比如nums = [1,1,1,1,1], target = 3
从最后一个数开始遍历 [1,1,1,1] +1 或者 [1,1,1,1] -1 =3
其实就能推出 [1,1,1,1] =2 或者 =4

因为题目返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目
加入定义了dp,那希望dp数组直接表示这个表达式的个数
dp494[i] 表示对应nums[i] 这个数，可以合成target的表达式的个数
dp494[i] 能怎么由前面的数推导出来呢，好像没看出来
dp494[i] 应该组成 target-nums[i] 或者 组成 target+nums[i] 两种情况的个数和

dp494[i] 和 dp494[i-x]怎么有关系呢

继续例子中的 nums = [1,1,1,1,1], target = 3 找规律 进一步的找 [1,1,1,1] =2 或者 =4
[1,1,1,1] =2 就是找 [1,1,1] 1 /3  和 [1,1,1] 3 / 5  这里前三个数等于3之后 要得到target5 两种方式
要求 进一步的  [1,1,1] 1 /3  和 [1,1,1] 3 / 5
其实就是算  [1,1]  t 等于 1-1 1+1 / 3-1 3+1 和 3-1 3+1 /5-1 5+1 的情况，要求这个其实就是求
[1]  等于 0-1 0+1 2-1 2+1 / 2-1 2+1 4-1 4+1 和 [1,1] 3-1 3+1 /5-1 5+1
好像dp一下写不出来，先递归写一下？
*/
func findTargetSumWays1(nums []int, target int) int {
	if len(nums) == 0 {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}

	ans := findTargetSumWays1(nums[:len(nums)-1], target-nums[len(nums)-1]) + findTargetSumWays1(nums[:len(nums)-1], target+nums[len(nums)-1])
	return ans
}

/*
这个写完，这个应该会超时吧，怎么优化下 有多少重复计算
 [1,1,1] 1 /3  和 [1,1,1] 3 / 5 数组一样，target 一样 这种应该直接给结果，不用重复计算了那就优化下
因为是从后往前依次遍历的，所以数组长度代替数组 然后 target值 两个元素一样的结果存起来
两个元素一样，然后才把这个值存起来
*/
var positive494 [][]int
var negative494 [][]int

func findTargetSumWays2(nums []int, target int) int {
	//res494[数组长度][target值]
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	positive494 = make([][]int, n+1)
	for x := 0; x < n+1; x++ {
		// 因为target有可能为负数 全正数最大值为sum，那全负数最大值就是-sum
		//那应该有2*sum的长度表示每种求和的可能性，but数组最小的index才0 负数对应的target存的坐标得转换一下
		//再或者2个数组，一个正数的数组，一个负数的数组
		positive494[x] = make([]int, sum+1)
		//	可能性不可能为负数吧，所以将其值全部初始化成-1
		for i := 0; i < sum+1; i++ {
			positive494[x][i] = -1
		}
	}

	negative494 = make([][]int, n+1)
	for x := 0; x < n+1; x++ {
		negative494[x] = make([]int, sum+1)
		//	可能性不可能为负数吧，所以将其值全部初始化成-1
		for i := 0; i < sum+1; i++ {
			negative494[x][i] = -1
		}
	}

	return dfs494(nums, target)
}

func dfs494(nums []int, target int) int {
	if len(nums) == 0 {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}
	n := len(nums)
	if target >= 0 {
		if positive494[n][target] != -1 {
			return positive494[n][target]
		}
	} else {
		if negative494[n][-target] != -1 {
			return negative494[n][-target]
		}
	}

	last := nums[n-1]
	newLeft := make([]int, n-1)
	copy(newLeft, nums[:n-1])
	newRight := make([]int, n-1)
	copy(newRight, nums[:n-1])
	ans := dfs494(newLeft, target-last) + dfs494(newRight, target+last)
	if target >= 0 {
		if positive494[n][target] == -1 {
			positive494[n][target] = ans
		}
	} else {
		if negative494[n][-target] == -1 {
			negative494[n][-target] = ans
		}
	}
	return ans
}

//实在是不好优化 结果发现可以这样  两个元素一样，然后才把这个值存起来  String key = u + "_" + cur;
//因为两个元素是都要一样才取值，那就索性把两个元素组合起来，变成一个元素

//下面想dp的方法，之前想了一半
//其实写完上面那些，不难发现其实什么变量，条件乱七八糟的不少
//dp[i]如果不够表示，dp[i][j]来表示
//https://leetcode.cn/problems/target-sum/solutions/816711/gong-shui-san-xie-yi-ti-si-jie-dfs-ji-yi-et5b/

var dp494 [][]int

/*
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
dp494[i][i] 表示 到nums[i] 这个数 计算结果没target的 目标为target时 有多少个。
*/

func findTargetSumWays3(nums []int, target int) int {
	if len(nums) == 1 {
		ans := 0
		if nums[0] == target {
			ans++
		}
		if nums[0]+target == 0 {
			ans++
		}
		return ans
	}
	dp494 = make([][]int, len(nums))
	//sum := 0
	//for _, num := range nums {
	//	sum += num
	//}
	// 这里不应该自己计算sum的最大值， 有可能sum全加起来也不够target
	// 直接使用上限 0 <= sum(nums[i]) <= 1000

	for x := 0; x < len(nums); x++ {
		//这里得确定数组长度，同时当前的和有可能是负数
		//假如算出来总和是 2 那个取值范围应该是 -2 ,2 这个5个数，对应长度应该是2*2+1
		//到时候存数据的时候，每个 index + sum 的值做偏移 比如是-2 这个应该存在数组第一个才对，那他对应index 0
		// -2 + sum 然后存进去就哦了
		dp494[x] = make([]int, 2*1000+1)
		//然后对应的可能性，最小是0不可能存在负数的可能性，如果到时候getmax 0 当默认值也没有问题
	}
	if len(nums) > 0 {
		first := nums[0]
		dp494[0][1000+first] += 1
		dp494[0][1000-first] += 1
	}
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		//因为 我假设 dp494[i][j] 表示 到nums[i] 这个数 计算结果没target的 目标为target时 有多少个。
		//怎么由前面的数递推出来呢，先找dp494[i][j] 中不为0的数，
		// 不为0的数就证明当前 这两个坐标对应的数是可以算得到的，而且为几就表示有多少种算式凑出来
		// 然后在依次加上nums[i]或者减去nums[i] 然后在dp[i][xx] 上分别加1
		for j := 0; j < 2*1000+1; j++ {
			for dp494[i-1][j] != 0 {
				//currentTarget := j - sum
				dp494[i][j+num] += 1
				dp494[i][j-num] += 1
				dp494[i-1][j] -= 1
			}
		}
	}
	return dp494[len(nums)-1][target+1000]
}

//dp这个还能继续优化

func TestAl494(t *testing.T) {
	//fmt.Println(findTargetSumWays1([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
	//fmt.Println(findTargetSumWays2([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))  这个代码有问题，可以考虑 map优化一下
	//fmt.Println(findTargetSumWays3([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
	//b := time.Now()
	//fmt.Println(findTargetSumWays3([]int{0}, 0))
	fmt.Println(findTargetSumWays3([]int{2, 107, 109, 113, 127, 131, 137, 3, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 47, 53}, 1000))
	//fmt.Printf("%v\n", time.Since(b))
	//fmt.Println(findTargetSumWays3([]int{1, 2, 3, 4, 5}, 2))
}
