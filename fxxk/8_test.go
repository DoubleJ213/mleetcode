package fxxk

import (
	"fmt"
	"math"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)

var isNe bool
var getSymbol bool
var getNum bool

func myAtoi(s string) int {
	ans := 0
	getSymbol = false
	getNum = false
	isNe = false
	for i := 0; i < len(s); i++ {
		//空格: 32  -: 45  +: 43
		if (!getNum && !getSymbol) && s[i] == 32 {
			//去前导空格
			continue
		}
		if !getSymbol && s[i] == 45 {
			//判断正还是负数
			isNe = true
			getSymbol = true
			continue
		} else if !getSymbol && s[i] == 43 {
			//判断正还是负数
			isNe = false
			getSymbol = true
			continue
		}

		if ok, num := isNum(s[i]); ok {
			//如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，
			//需要截断这个整数，使其保持在这个范围内。具体来说，
			//小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
			getNum = true
			if !getSymbol {
				//都已经出数字了，就默认正数，不要继续判断正负号了
				isNe = false
				getSymbol = true
			}
			ans = ans*10 + num
			if !isNe && ans >= math.MaxInt32 {
				return math.MaxInt32
			} else if isNe && -ans <= math.MinInt32 {
				return math.MinInt32
			}
		} else {
			break
		}
	}
	if isNe {
		return -ans
	} else {
		return ans
	}
}

func isNum(b byte) (bool, int) {
	//0:48 1:49 2:50 3:51 4:52 5:53 6:54 7:55 8:56 9:57
	switch b {
	case 48:
		return true, 0
	case 49:
		return true, 1
	case 50:
		return true, 2
	case 51:
		return true, 3
	case 52:
		return true, 4
	case 53:
		return true, 5
	case 54:
		return true, 6
	case 55:
		return true, 7
	case 56:
		return true, 8
	case 57:
		return true, 9

	}
	return false, -1
}

//leetcode submit region end(Prohibit modification and deletion)

// 0 <= s.length <= 200
// s 由英文字母（大写和小写）、数字（0-9）、' '、'+'、'-' 和 '.' 组成
func TestStringToIntegerAtoi(t *testing.T) {
	//fmt.Println(myAtoi("00000-42a1234"))
	//fmt.Println(myAtoi("   +0 123"))
	//fmt.Println(myAtoi("-9128347231"))
	//fmt.Println(myAtoi(" -41 1"))
	//fmt.Println(myAtoi(" -4a1 1"))
	//fmt.Println(myAtoi(" -a4a1 1"))
	//fmt.Println(myAtoi("  +  413"))
	//fmt.Println(myAtoi("  + 0 413"))
	fmt.Println(myAtoi(" + 413"))
	//-2147483648
	//-9128347231
}

//[-2^31, 2^31 - 1]
