package fxxk

import (
	"fmt"
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
/*
厨房里总共有 n 个橘子，你决定每一天选择如下方式之一吃这些橘子：
吃掉一个橘子。
如果剩余橘子数 n 能被 2 整除，那么你可以吃掉 n/2 个橘子。
如果剩余橘子数 n 能被 3 整除，那么你可以吃掉 2*(n/3) 个橘子。
每天你只能从以上 3 种方案中选择一种方案。
请你返回吃掉所有 n 个橘子的最少天数。

三种吃法
吃掉1个橘子
当可以被2整除,吃掉剩余的 1/2 剩余1/2
当可以被3整除,吃掉剩余的 2/3 剩余1/3
求吃完全部橘子 最少得天数
显然是个dp题，但是dp[n]等于啥呢
dp 有前面的数推导，前面3个数能吃到n 的情况 但是推导的index需要计算，有点复杂
*/

var dp1553 []int

//func minDays2(n int) int {
//	//让dp[n]为最终的结果 0 表示没有橘子需要吃，那就是0天
//	//当1个橘子要吃，那就是1天
//	dp1553 = make([]int, n+1) //n=2000000000 这得多大内存，这个得想办法优化
//	//其实只需要知道前面一个数，前一个/2 /3 这两个数分别是多少就行
//	dp1553[0] = 0
//	dp1553[1] = 1
//	for i := 2; i < n+1; i++ {
//		if i == 2 || i == 3 {
//			dp1553[i] = 2
//		} else {
//			a, b := getIndex(i)
//			//fmt.Printf("i = %d , index = %d\n", i, index)
//			/*
//				三种方式求最低值
//				吃掉1个橘子
//				当可以被2整除,吃掉剩余的 1/2 剩余1/2
//				当可以被3整除,吃掉剩余的 2/3 剩余1/3
//			*/
//			dp1553[i] = dp1553[i-1] + 1
//			if a != -1 {
//				dp1553[i] = getMin(dp1553[i], dp1553[a]+1)
//			}
//			if b != -1 {
//				dp1553[i] = getMin(dp1553[i], dp1553[b]+1)
//			}
//		}
//	}
//	return dp1553[n]
//}

var res1553 map[int]int

func minDays(n int) int {
	res1553 = make(map[int]int)
	// 这么长的数组有点懵改成按需分配的吧
	res1553[0] = 0
	res1553[1] = 1
	dfs1553(n)
	return res1553[n]
}

// dfs(i) 表示把 i 变成 0 的最少操作次数。
func dfs1553(n int) int {
	//先执行 i mod 2 次减 1 操作，把 i 变成 2 的倍数，然后再执行一次除以 2，
	//问题变成把  i/2 变成 0 的最小操作次数，即
	//dfs(i)=dfs(⌊i/2⌋)+imod2+1。
	if num, ok := res1553[n]; ok {
		return num
	}
	res1553[n] = getMin(dfs1553(n/3)+n%3, dfs1553(n/2)+n%2) + 1
	return res1553[n]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumNumberOfDaysToEatNOranges(t *testing.T) {
	//fmt.Println(minDays(5))
	//1 <= n <= 2*10^9
	fmt.Println(minDays(100))
}
