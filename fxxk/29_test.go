package fxxk

import (
	"fmt"
	"math"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)

func divide(dividend int, divisor int) int {
	//注意题目限制  假设我们的环境只能存储 32 位 有符号整数，其数值范围是 [−231,  231 − 1] 。
	//本题中，如果商 严格大于 2^31 − 1 ，则返回 2^31 − 1 ；如果商 严格小于 -2^31 ，则返回 -2^31 。
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	//divisor != 0
	if dividend == 0 {
		return 0
	} else if dividend < 0 && divisor < 0 {
		return divide(-dividend, -divisor)
	} else if dividend < 0 && divisor > 0 {
		return -divide(-dividend, divisor)
	} else if dividend > 0 && divisor < 0 {
		return -divide(dividend, -divisor)
	}
	//dividend, divisor 都>0 我们需要找到得数x
	//满足，如果有一个数y, 如果y比x小  y*divisor < dividend
	//如果y比x大   y*divisor  > dividend
	//所以得找一个 满足条件 y*divisor<= dividend  中最大的y
	left := 0
	right := dividend
	mid := left
	//上面说找最大的ans  ans和mid关系选择，最好举个例子。
	//当myMultiple算出来于dividend相等，直接返回。
	//当myMultiple算出来小于dividend，mid可以暂时作为ans，继续找。
	//myMultiple算出来大于dividend，ans不能选择mid，继续找。
	ans := 0
	for left <= right {
		//mid := (right-left)/2 + left 写这个的时候发现这里面用了除法，题目不允许的
		mid = (right-left)>>1 + left
		if myMultiple(mid, divisor) == dividend {
			return mid
		} else if myMultiple(mid, divisor) > dividend {
			right = mid - 1
		} else if myMultiple(mid, divisor) < dividend {
			left = mid + 1
			ans = mid
		}
	}
	return ans
}

/*
myMultiple
比如 5*3 = 5*(ob 11) = 5 * (1* 2^1+1* 2^0) =
5 *  1* 2^1 + 5 *  1 * 2^0
其中 需要判断这个最低位是0 还是1 如果是1就加上乘出来的结果0 如果是0 就pass。
所以进一步的看，5*2^1 ,5*2^0 这个数列其实就是对应的 乘数5 不停的翻倍，被乘数不停的右移。
*/
func myMultiple(x int, y int) int {
	ans := 0
	for y > 0 {
		if y&1 == 1 {
			// 当前最低位为1，结果里加上x
			ans += x
		}
		// 被乘数右移1位，相当于除以2
		y = y >> 1
		// 乘数倍增，相当于乘以2
		x += x
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDivideTwoIntegers(t *testing.T) {
	fmt.Printf("5 --- %v\n", divide(11, 2))
	fmt.Printf("-2 --- %v\n", divide(-7, 3))
	fmt.Printf("-2 --- %v\n", divide(7, -3))
	fmt.Printf("1073741823**** --- %v\n", divide(2147483647, 2))
	fmt.Printf("2147483648**** --- %v\n", divide(-2147483648, -1))
	//-2147483648  -2^31
	fmt.Printf("4 --- %v\n", divide(12, 3))
	fmt.Printf("6 --- %v\n", divide(20, 3))
}
