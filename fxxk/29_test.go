package fxxk

import (
	"fmt"
	"testing"
)

/*
29. 两数相除
给你两个整数，被除数 dividend 和除数 divisor。将两数相除，要求 不使用 乘法、除法和取余运算。
整数除法应该向零截断，也就是截去（truncate）其小数部分。例如，8.345 将被截断为 8 ，-2.7335 将被截断至 -2 。
返回被除数 dividend 除以除数 divisor 得到的 商 。

注意：假设我们的环境只能存储 32 位 有符号整数，其数值范围是 [−2^31,  2^31 − 1] 。
本题中，如果商 严格大于 2^31 − 1 ，则返回 2^31 − 1 ；如果商 严格小于 -2^31 ，则返回 -2^31 。

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = 3.33333.. ，向零截断后得到 3 。
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = -2.33333.. ，向零截断后得到 -2 。

提示：
-2^31 <= dividend, divisor <= 2^31 - 1
divisor != 0
*/

// 2^31 -1 = 2147483647
//-2^31  = -2147483648
/*var deep29 int

func divide(dividend int, divisor int) int {
	//divisor != 0 我就不判断了
	status29 := true
	//正数用true表示、负数用false
	if dividend < 0 {
		dividend = - dividend
		status29 = false
	}
	if divisor < 0 {
		divisor = - divisor
		if status29 == true {
			status29 = false
		} else {
			status29 = true
		}
	}
	//已经是正数除以正数了
	deep29 = performDivision(dividend, divisor, 0, status29)
	//如果商 严格大于 2^31 − 1 ，则返回 2^31 − 1 ；如果商 严格小于 -2^31 ，则返回 -2^31

	if ! status29 {
		deep29 = -deep29
	}

	return deep29
}

func performDivision(dividend int, divisor int, deep int, status bool) int {
	if status && deep29 > 2147483647 {
		return 2147483647
	} else if !status && deep29 < -2147483648 {
		return -2147483648
	}

	if dividend < divisor {
		return deep
	}
	return performDivision(dividend-divisor, divisor, deep+1, status)
}*/

var deep29 int

//想了想可以把递归去掉，直接for循环累加就完事了
func divide2(dividend int, divisor int) int {
	//divisor != 0 我就不判断了
	max29 := 2147483647
	min29 := -2147483648
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == min29 {
			return max29
		}
		return -dividend
	}

	status29 := true
	//正数用true表示、负数用false
	if dividend < 0 {
		dividend = -dividend
		status29 = false
	}
	if divisor < 0 {
		divisor = -divisor
		if status29 == true {
			status29 = false
		} else {
			status29 = true
		}
	}
	//已经是正数除以正数了
	for dividend > divisor {
		dividend -= divisor
		deep29++
	}
	//如果商 严格大于 2^31 − 1 ，则返回 2^31 − 1 ；如果商 严格小于 -2^31 ，则返回 -2^31

	if !status29 {
		deep29 = -deep29
	}

	if deep29 > max29 {
		return max29
	}
	if deep29 < min29 {
		return min29
	}

	return deep29
}

/*
上面版本递归太多次了，可以用这样的思路：
60/8 = (60-32)/8 + 4 = (60-32-16)/8 + 2 + 4 = 1 + 2 + 4 = 7

比如要算60/8  先将8 翻一倍，发现60 比这个数大，起码可以 60/16
继续将16翻一倍32，60比32还大。那就尝试再翻一倍64。发现不够除以64了
在刚刚递归的过程中，8其实翻了4倍，不够翻第8倍了
60-32 记录一下当前相当于已经除了4个8  相当于继续计算 4 + 28/8
这个逻辑可以先写一下，但是发现其实可以优化
如果继续递归判断28/8够不够，应该有重复计算
*/

func divide3(dividend int, divisor int) int {
	//前面这一大段逻辑，我继续保留
	max29 := 2147483647
	min29 := -2147483648
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == min29 {
			return max29
		}
		return -dividend
	}

	status29 := true
	//正数用true表示、负数用false
	if dividend < 0 {
		dividend = -dividend
		status29 = false
	}
	if divisor < 0 {
		divisor = -divisor
		if status29 == true {
			status29 = false
		} else {
			status29 = true
		}
	}

	ans := preDiv(dividend, divisor, 0)

	if !status29 {
		ans = -ans
	}

	if ans > max29 {
		return max29
	}
	if ans < min29 {
		return min29
	}

	return ans
}

func preDiv(dividend, divisor, ans int) int {
	count := 1
	preDivisor := divisor
	for dividend > divisor+divisor {
		count += count
		divisor += divisor
	}
	ans += count
	//divisor -= divisor 减几次？
	newDivisor := dividend - divisor

	if newDivisor == preDivisor {
		return ans + 1
	} else if newDivisor < preDivisor {
		return ans
	}
	return preDiv(newDivisor, preDivisor, ans)
}

func TestAl29(t *testing.T) {
	//fmt.Printf("5 --- %v\n", divide(10, 2))
	//fmt.Printf("-5 --- %v\n", divide(-10, 2))
	//fmt.Printf("**** --- %v\n", divide(9, 8))
	fmt.Printf("1073741823**** --- %v\n", divide3(2147483647, 2))
	//fmt.Printf("**** --- %v\n", divide3(2147483647, 2))
	fmt.Printf("**** --- %v\n", divide3(14, 3))
}
