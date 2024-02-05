package fxxk

import (
	"fmt"
	"strconv"
	"testing"
)

/*
241. 为运算表达式设计优先级
给你一个由数字和运算符组成的字符串 expression ，按不同优先级组合数字和运算符，
计算并返回所有可能组合的结果。你可以 按任意顺序 返回答案。
生成的测试用例满足其对应输出值符合 32 位整数范围，不同结果的数量不超过 10^4 。

示例 1：
输入：expression = "2-1-1"
输出：[0,2]
解释：
((2-1)-1) = 0
(2-(1-1)) = 2
示例 2：
输入：expression = "2*3-4*5"
输出：[-34,-14,-10,-10,10]
解释：
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10

((2*3)-(4*5)) = -14 (已经算过)

提示：
1 <= expression.length <= 20
expression 由数字和算符 '+'、'-' 和 '*' 组成。
输入表达式中的所有整数值在范围 [0, 99]
*/

//https://labuladong.gitee.io/algo/di-san-zha-24031/jing-dian--a94a0/fen-zhi-su-65a39/#%E6%B7%BB%E5%8A%A0%E6%8B%AC%E5%8F%B7%E7%9A%84%E6%89%80%E6%9C%89%E6%96%B9%E5%BC%8F
func diffWaysToCompute(expression string) []int {
	//抽象成数、二叉树结构是个难点
	// 以每个运算符分割左右两边的算式，然后就类似二叉树了。
	// 要求当前算式有哪些可能性，那就算上左子树所有可能性，再加上右子树所有可能性
	// 外层for循环，用于分割算式的字符串。当 当前字符是+ - * 三个字符的时候说明分割正确的，需要递归当前形状的左右子树的所有可能性
	// 回溯也好，分治也好，二叉树后序处理也好，反正就是在统计出以当前运算符分割的 左右子树的所有可能性分别做运算，能得到需要的结果
	// 终止条件是啥，一个是循环走完了。那递归终止条件是啥，一直找子树，找到啥时候结束。当res为0说明左右子树都是0，自己已经是孤零零的一个数了
	// 直接把这个数加进res。当成别人的左或者右子树[]int结果

	return dfs241(expression)
}

func dfs241(expression string) []int {
	res := make([]int, 0)
	for i := 0; i < len(expression); i++ {
		current := string(expression[i])
		if current == "-" || current == "+" || current == "*" {
			left := dfs241(expression[0:i])
			right := dfs241(expression[i+1:])

			for x := 0; x < len(left); x++ {
				for y := 0; y < len(right); y++ {
					ans := getRes241(left[x], right[y], current)
					res = append(res, ans)
				}
			}
		}
	}

	// 如果 res 为空，说明算式是一个数字，没有运算符
	if len(res) == 0 {
		num, _ := strconv.Atoi(expression)
		res = append(res, num)
	}
	return res
}

func getRes241(x int, y int, s string) int {
	if s == "-" {
		return x - y
	} else if s == "+" {
		return x + y
	}
	return x * y
}

func TestAl241(t *testing.T) {
	//fmt.Println(diffWaysToCompute("2-1+1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}
