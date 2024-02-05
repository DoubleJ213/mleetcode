package fxxk

import (
	"fmt"
	"strconv"
	"testing"
)

/*
282. 给表达式添加运算符
给定一个仅包含数字 0-9 的字符串 num 和一个目标值整数 target ，
在 num 的数字之间添加 二元 运算符（不是一元）+、- 或 * ，返回 所有 能够得到 target 的表达式。
注意，返回表达式中的操作数 不应该 包含前导零。

示例 1:
输入: num = "123", target = 6
输出: ["1+2+3", "1*2*3"]
解释: “1*2*3” 和 “1+2+3” 的值都是6。
示例 2:
输入: num = "232", target = 8
输出: ["2*3+2", "2+3*2"]
解释: “2*3+2” 和 “2+3*2” 的值都是8。
示例 3:
输入: num = "3456237490", target = 9191
输出: []
解释: 表达式 “3456237490” 无法得到 9191 。

提示：
1 <= num.length <= 10
num 仅含数字
-2^31 <= target <= 2^31 - 1
*/

/*	num 每个数都分割出来 num 仅含数字
	比如是2个数 12，求2个数的各种可能性
	1-2 1+2 1*2三种值 同时还得直接拿到对应 “1-2” “1+2” “1*2”三种string类型的值
	这里就返回string类型的值吧，后面再求具体等于几？如果一个map能不能存.但是一样的string容易丢掉
	那暂时可以用2个数组，数组对应位置存对应的值，省得在用string计算一遍值
	测试 "123"  发现 输出nil 。第一个数前面不能添加运算符。*/

func addOperators(num string, target int) []string {
	if len(num) == 0 {
		return nil
	}
	ans282 := make([]string, 0)
	curStr := string(num[0])
	curNum, _ := strconv.Atoi(curStr)
	strList282 := make([]string, 0)
	countList282 := make([]int, 0)
	strList282 = append(strList282, curStr)
	countList282 = append(countList282, curNum)

	if len(num) == 1 {
		if curNum == target {
			return strList282
		}
		return ans282
	}

	for i := 1; i < len(num); i++ {
		current := string(num[i])
		curNum, _ := strconv.Atoi(current)
		size := len(strList282)
		levelStr := make([]string, 0)
		levelCount := make([]int, 0)
		for x := 0; x < size; x++ {
			levelStr = append(levelStr, fmt.Sprintf("%v%v%v", strList282[x], "+", current))
			levelCount = append(levelCount, countList282[x]+curNum)
			levelStr = append(levelStr, fmt.Sprintf("%v%v%v", strList282[x], "-", current))
			levelCount = append(levelCount, countList282[x]-curNum)
			levelStr = append(levelStr, fmt.Sprintf("%v%v%v", strList282[x], "*", current))
			levelCount = append(levelCount, countList282[x]*curNum)
		}
		strList282 = levelStr
		countList282 = levelCount
	}

	for i := 0; i < len(strList282); i++ {
		if countList282[i] == target {
			ans282 = append(ans282, strList282[i])
		}
	}
	return ans282
}

func TestAl282(t *testing.T) {
	//fmt.Println(addOperators("123", 6))
	fmt.Println(addOperators("105", 5))
	//	一个string要考虑乘号的优先级 	["1*0+5","10-5"]
}
