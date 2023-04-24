package fxxk

import (
	"fmt"
	"testing"
)

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//注意：给定 n 是一个正整数。
func climbStairs(n int) int {
	//排列组合解法 先简单搞清楚前几步，找规律
	//当n=3时：
	//可以0个2，3个1 可以1个2，1个1， 不可以2个2
	//当n=4时：
	//可以0个2，4个1 C(4,4)
	//可以1个2，2个1 C(3,2)
	//可以2个2，0个1 C(2,2)
	//不可以3个2
	//但是排列组合有没有库能直接计算，没有得自己写，先放弃

	//递推
	//假设f(n)为第n层楼梯的所有走法，那f(n)=f(n-1)+f(n-2)
	if n <= 0 {
		// 协商好
		//return 0
		return 1
	}
	tmp := make([]int, n+1)
	tmp[0] = 1
	tmp[1] = 1
	tmp[2] = 2
	for i := 3; i < n+1; i++ {
		tmp[i] = tmp[i-1] + tmp[i-2]
	}
	return tmp[n]
}

func TestAl70(t *testing.T) {
	fmt.Println(climbStairs(7))
}
