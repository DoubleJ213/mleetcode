package fxxk

import (
	"testing"
)

/*
474. 一和零
给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。

示例 1：
输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
输出：4
解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。
示例 2：
输入：strs = ["10", "0", "1"], m = 1, n = 1
输出：2
解释：最大的子集是 {"0", "1"} ，所以答案是 2 。

提示：
1 <= strs.length <= 600
1 <= strs[i].length <= 100
strs[i] 仅由 '0' 和 '1' 组成
1 <= m, n <= 100
*/

//之前背包问题只有一个条件限制的时候是个二维数组，现在有2个条件限制，同时还得遍历数组strs的元素
//所以这里初始一个三维数组  题目问啥，dp就表示啥
//dp[i][j][k] 表示输入字符串在子区间 [0, i] 能够使用 j 个 0 和 k 个 1 的字符串的最大数量。
//之前背包问题是可以用一个一维数组的，这里按道理是可以用一个二维数组，先不写了
var dp474 [][][]int

func findMaxForm1(strs []string, m int, n int) int {
	length := len(strs)
	dp474 = make([][][]int, length+1)
	for x := 0; x < length+1; x++ {
		dp474[x] = make([][]int, m+1)
		for y := 0; y < m+1; y++ {
			dp474[x][y] = make([]int, n+1)
		}
	}
	for i := 1; i < length+1; i++ {
		zeros, ones := getZerosAndOnes(strs[i-1])
		for j := 0; j < m+1; j++ { //m 个0
			for k := 0; k < n+1; k++ { // n 个1
				//物品一个一个尝试，容量一点一点尝试
				if j >= zeros && k >= ones {
					dp474[i][j][k] = getMax(dp474[i-1][j][k], dp474[i-1][j-zeros][k-ones]+1)
				} else {
					dp474[i][j][k] = dp474[i-1][j][k]
				}
			}
		}

	}
	return dp474[length][m][n]
}

func getZerosAndOnes(s string) (int, int) {
	zeros, ones := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zeros++
		} else {
			ones++
		}
	}
	return zeros, ones
}

func TestAl474(t *testing.T) {
	//fmt.Println(findMaxForm1([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	//fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}

//https://leetcode.cn/problems/ones-and-zeroes/solutions/814942/gong-shui-san-xie-xiang-jie-ru-he-zhuan-174wv/
