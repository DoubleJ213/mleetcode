package fxxk

import (
	"fmt"
	"testing"
)

/*
2140. 解决智力问题

给你一个下标从 0 开始的二维整数数组 questions ，其中 questions[i] = [pointsi, brainpoweri] 。
这个数组表示一场考试里的一系列题目，你需要 按顺序 （也就是从问题 0 开始依次解决），针对每个问题选择 解决 或者 跳过 操作。解决问题 i 将让你 获得  pointsi 的分数，
但是你将 无法 解决接下来的 brainpoweri 个问题（即只能跳过接下来的 brainpoweri 个问题）。
如果你跳过问题 i ，你可以对下一个问题决定使用哪种操作。
比方说，给你 questions = [[3, 2], [4, 3], [4, 4], [2, 5]] ：
如果问题 0 被解决了， 那么你可以获得 3 分，但你不能解决问题 1 和 2 。
如果你跳过问题 0 ，且解决问题 1 ，你将获得 4 分但是不能解决问题 2 和 3 。
请你返回这场考试里你能获得的 最高 分数。

示例 1：
输入：questions = [[3,2],[4,3],[4,4],[2,5]]
输出：5
解释：解决问题 0 和 3 得到最高分。
- 解决问题 0 ：获得 3 分，但接下来 2 个问题都不能解决。
- 不能解决问题 1 和 2
- 解决问题 3 ：获得 2 分
总得分为：3 + 2 = 5 。没有别的办法获得 5 分或者多于 5 分。
示例 2：
输入：questions = [[1,1],[2,2],[3,3],[4,4],[5,5]]
输出：7
解释：解决问题 1 和 4 得到最高分。
- 跳过问题 0
- 解决问题 1 ：获得 2 分，但接下来 2 个问题都不能解决。
- 不能解决问题 2 和 3
- 解决问题 4 ：获得 5 分
总得分为：2 + 5 = 7 。没有别的办法获得 7 分或者多于 7 分。

提示：
1 <= questions.length <= 10^5
questions[i].length == 2
1 <= pointsi, brainpoweri <= 10^5
*/

//var res
/*理解完题目，我猜这种求最值的dp来解决了，没有更好的办法了，那dp数组怎么定义呢
一般题目都是要求什么，然后让dp数组表示什么，假如 dp[n] 表示做第n题然后得到的最高分？
然后dp[n] = dp[n-nums[x2]] + nums[x1]
*/
var dp2140 []int
var max2140 int

func mostPoints(questions [][]int) int64 {
	max2140 = 0
	dp2140 = make([]int, len(questions))
	//因为是最大值
	dp2140[0] = questions[0][0]
	if len(questions) == 1 {
		return int64(questions[0][0])
	}
	//dp 表示 选择做本题 然后合计前面的数据的最高分
	for i := 1; i < len(questions); i++ {
		//dp2140[n] = getMax(0~n-1) 加上本题分数
		cs := questions[i][0]
		for j := 0; j < i; j++ {
			wide := questions[j][1]
			if j+wide < i {
				dp2140[i] = getMax(dp2140[i], dp2140[j])
			}
		}
		dp2140[i] += cs
		max2140 = getMax(max2140, dp2140[i])
	}
	return int64(max2140)
}

/*
发现嵌套了个循环，超时了 O(n^2)，dp还得优化
dp 表示 截止到本数据的最高分。本数据可能是选择了，也可能是不被选择
本题 选择跳过，不加此题得分，后续的题继承得分
本题 选择做题，加上此题得分，跳过若干题后的下一题继承得分
*/
func mostPoints2(questions [][]int) int64 {
	size := len(questions)
	dp2140 = make([]int, size+1)
	if len(questions) == 1 {
		return int64(questions[0][0])
	}
	//第一个数选择不选择对后续一个数的影响都得提现出来
	for i := size - 1; i >= 0; i-- {
		//从小往大整不太好弄，那就反过来
		//解决本题与否会影响到后面一定数量题目的结果，但不会影响到前面题目的解决。
		// 因此我们可以考虑从反方向定义「状态」，即考虑解决每道题本身及以后的题目可以获得的最高分数。
		if i+questions[i][1]+1 < size {
			//我们用 dp[i] 来表示解决第 i 道题目及以后的题目可以获得的最高分数
			dp2140[i] = getMax(dp2140[i+1], dp2140[i+questions[i][1]+1]+questions[i][0])
		} else {
			dp2140[i] = getMax(dp2140[i+1], questions[i][0])
		}
	}
	return int64(dp2140[0])
}

//var dp2140 []int

func mostPoints3(questions [][]int) int64 {
	size := len(questions)
	dp2140 = make([]int, size+1)
	if len(questions) == 1 {
		return int64(questions[0][0])
	}
	//第一个数选择不选择对后续一个数的影响都得提现出来
	for i := 0; i < size; i++ {
		score := questions[i][0]
		wide := questions[i][1]
		//i 不被选择 ，那dp[i+1] 最小也是 dp[i]
		dp2140[i+1] = getMax(dp2140[i+1], dp2140[i])
		//	i选择时。
		next := getMin(i+wide+1, size)
		dp2140[next] = getMax(dp2140[next], dp2140[i]+score)
	}
	return int64(dp2140[size])
}

func TestAl2140(t *testing.T) {
	fmt.Println(mostPoints3([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
	//fmt.Println(mostPoints2([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}))
	//fmt.Println(mostPoints2([][]int{{12, 46}, {78, 19}, {63, 15}, {79, 62}, {13, 10}}))
	//	[[12,46],[78,19],[63,15],[79,62],[13,10]]
}
