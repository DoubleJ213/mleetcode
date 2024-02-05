package fxxk

import (
	"fmt"
	"testing"
)

/*
204. 计数质数
给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。

示例 1：
输入：n = 10
输出：4
解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
示例 2：
输入：n = 0
输出：0
示例 3：
输入：n = 1
输出：0

提示：
0 <= n <= 5 * 10^6
*/

/*
1.第一种。比如判断n是不是素数，从2开始一直整除，然后整除到n-1.
2.第二种，相对上面的，其实可以发现只需要从2 整除到根号n sqrt(n)就行了.前一半如果没能找到因子，后面一半也就不会有了
3.第三种。可以看204的动图。初始化一个数组，用于存bool值。然后发现值为2是素数。那标记所有2的倍数都不是素数
标记到n,3是素数，继续标记所有3的倍数。4已经被刚刚2倍数的时候标记过了，5还是false。标记其他所有5的倍数。
6被标记为true了。7false，标记所有7的倍数。8、9、10标记过。11 还是false。11 是素数。标记所有11的倍数
依次类推
*/
func countPrimes(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	res204 := make([]bool, n+1)
	//默认值全部false。如果不是素数true
	//	n+1 个元素 index 0~n  0  是素数吗？ 不清楚哎，但是题目示例有 n = 0 输出：0
	index := 2
	ans204 := 0
	for index < n {
		//if res204[index] ==false {
		x := 2 // 表示倍数
		if !res204[index] {
			ans204++
		}
		for !res204[index] && index*x < n {
			res204[index*x] = true
			x++
		}
		index++
	}
	return ans204
}

func TestAl204(t *testing.T) {
	fmt.Println(countPrimes(100))
	fmt.Println(countPrimes(2))
	fmt.Println(countPrimes(3))
	fmt.Println(countPrimes(4))
}
