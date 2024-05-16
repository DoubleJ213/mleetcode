package fxxk

import (
	"fmt"
	"testing"
)

/*
793. 阶乘函数后 K 个零
f(x) 是 x! 末尾是 0 的数量。回想一下 x! = 1 * 2 * 3 * ... * x，且 0! = 1 。
例如， f(3) = 0 ，因为 3! = 6 的末尾没有 0 ；而 f(11) = 2 ，因为 11!= 39916800 末端有 2 个 0 。
给定 k，找出返回能满足 f(x) = k 的非负整数 x 的数量。

示例 1：
输入：k = 0
输出：5
解释：0!, 1!, 2!, 3!, 和 4! 均符合 k = 0 的条件。
示例 2：
输入：k = 5
输出：0
解释：没有匹配到这样的 x!，符合 k = 5 的条件。
示例 3:
输入: k = 3
输出: 5

提示:
0 <= k <= 10^9
*/

func preimageSizeFZF1(k int) int {
	res793 := make([]int, 0)
	//f(3) = 0   f(11) = 2
	//num 表示 数字k 某位有几个0
	//给定k，找出返回能满足 f(x) = k 的非负整数 x 的数量。
	//一旦 num 到 5 或者0，f(x)就会变成另外一个数了
	if k == 0 {
		// 0 1 2 3 4
		return 5
	}
	res793 = append(res793, 0)
	i := 1
	x := 0
	for x <= k {
		currentSum := 0 // currentSum 为当前阶乘末尾几个0
		y := i
		for y%5 == 0 {
			currentSum++
			y = y / 5
		}
		x += currentSum
		res793 = append(res793, x)
		i++
	}

	sum := 0
	for i := 0; i < len(res793); i++ {
		if res793[i] == k {
			sum++
		}
	}
	return sum
}

/*
如果k很大，会存一个很大的数组，内存占用太多
优化一下，这个没有一个很大的数组了。
*/
func preimageSizeFZF2(k int) int {
	if k == 0 {
		// 0 1 2 3 4
		return 5
	}
	i := 1
	x := 0
	sum := 0
	for x <= k {
		currentSum := 0 // currentSum 为当前阶乘末尾几个0
		y := i
		for y%5 == 0 {
			currentSum++
			y = y / 5
		}
		x += currentSum
		if x == k {
			sum++
		}
		i++
	}
	return sum
}

// 然后发现，这个数要么是0 要么是5 只要能找到x=k 就是5 找不到就是0
func preimageSizeFZF3(k int) int {
	if k == 0 {
		// 0 1 2 3 4
		return 5
	}
	i := 1
	x := 0
	for x <= k {
		currentSum := 0 // currentSum 为当前阶乘末尾几个0
		y := i
		for y%5 == 0 {
			currentSum++
			y = y / 5
		}
		x += currentSum
		if x == k {
			return 5
		}
		if currentSum != 0 {
			//进一步发现，其实不是每个数都要遍历。反正找5的正数倍。
			i = i + 5
		} else {
			i++
		}
	}
	return 0
}

// 43103094 就超时了，我测试用10^9 他不报错，真的是日了
// 时间复杂度还得优化。我这个是K等于几可能就遍历多少遍
func preimageSizeFZF(k int) int {
	if k == 0 || k == 1 {
		// 0 1 2 3 4
		//5 6 7 8 9
		return 5
	}
	return getIndex(k+1) - getIndex(k)
}

func getIndex(k int) int {
	left := 0 // i从1开始，然后i最大 math.MaxInt64,然后二分法？这样不用每5个i 就判断一次
	//right := math.MaxInt64 / 5 * 5
	right := k * 5
	//这里发现其实 left=0 最后边应该是个5的整数倍
	//找0~tmp范围内最大的5的倍数的数为right 101/5 =20 20*5 =100
	//mid := (right + left) / 2
	//防止溢出，用下面的写法
	for left <= right {
		mid := left + (right-left)/2
		count := getCount793(mid)
		if count >= k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}

func getCount793(n int) int {
	res := 0
	for n != 0 {
		//分解 n 有多少个 5 因子即为尾部零的个数 比如12 包含2个5的倍数
		res++
		n = n / 5
	}
	return res
}

func TestAl793(t *testing.T) {
	//fmt.Println(preimageSizeFZF(2))
	//fmt.Println(preimageSizeFZF(3))
	fmt.Println(preimageSizeFZF(4))
	fmt.Println(preimageSizeFZF(5))
	fmt.Println(preimageSizeFZF(100000000))
	//fmt.Println(preimageSizeFZF2(100000000))
}
