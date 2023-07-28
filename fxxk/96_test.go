package fxxk

import "testing"

/*
96.不同的二叉搜索树
给你一个整数n，求恰由n个节点组成且节点值从1到n互不相同的二叉搜索树有多少种？
返回满足题意的二叉搜索树的种数。

示例1：
输入：n=3
输出：5
示例2：
输入：n=1
输出：1
提示：
1<=n<=19
*/

/*
哈哈，看到这个边界我就能想到，肯定有人19个枚举值直接得结果。
---------------
   2
  1

  1
    2
---------------
  3
 2
1

     3
 1
   2

1
  2
    3

    1
      3
     2

  2
1   3
---------------
f 1  =  1
f 2  =  2
f 3  =  5
f4= 5+5+2+2 = 14
f5 =42

f3 = f2 （3为根 左子树  为f2结果）+
1为根（还剩2个数，f2种树形）+
2为根（左右各1个数，f1种树形）

f4 = ？
（f3 4为根个数）+
（1为根 还剩3个数，f3种）+
（2为根，左边1个数，右边2个数。f2）+
（3为根，左边2个数 ， 右边一个数 f2 ）
f3 +f3+f2+f2

f5 =?
5为根，f4的个数
1为根，4个数全在右边，又是个f4
2为根，左边1个数，右边3个数 f1*f3
3为根，左边2个，右边2个 f2*f2
4为根，左边3个数，右边1个数。f3*f1
leftNum * rightNum
f4+f4+f1*f3+f2*f2+f3*f1
14+14+1*5+2*2+5 = 42
*/
var count []int

func numTrees(n int) int {
	//f1 = 1
	//f2 = 2
	count = make([]int, 20)
	count[0] = 1
	count[1] = 1
	count[2] = 2
	traverse96(n)
	return count[n]
}

func traverse96(n int) {
	if n == 1 || n == 2 {
		return
	}
	traverse96(n - 1)

	mid := 1
	var sum int
	left, right := 1, n
	for mid <= n {
		//需要知道左边和右边剩余几个数，然后对应count里面去找
		leftSum := count[mid-left]
		rightSum := count[right-mid]
		sum += leftSum * rightSum
		mid++
	}
	count[n] = sum
}

func TestAl96(t *testing.T) {
	numTrees(19)
}
