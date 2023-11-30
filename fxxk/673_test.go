package fxxk

import (
	"fmt"
	"testing"
)

/*
673. 最长递增子序列的个数
给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。
注意 这个数列必须是 严格 递增的。

示例 1:
输入: [1,3,5,4,7]
输出: 2
解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
示例 2:
输入: [2,2,2,2,2]
输出: 5
解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。

提示:
1 <= nums.length <= 2000
-10^6 <= nums[i] <= 10^6
*/

//dp[i] 表示以 nums[i] 这个数结尾的最长递增子序列的长度
//返回最长递增子序列的个数 数组最后再遍历一遍，看看最大值有几个
var dp673 []int
var max673 []int

//max673[i] 为考虑以 nums[i] 结尾的最长上升子序列的个数。
var maxNum673 int

func findNumberOfLIS(nums []int) int {
	//dp[i] 对应nums[i] 这个数结尾的最长递增的序列的长度
	//额外定义 max673 为考虑以 nums[i] 结尾的最长上升子序列的个数。
	//dp[i] = dp[i]
	n := len(nums)
	dp673 = make([]int, n)
	max673 = make([]int, n)
	maxNum673 = 1
	dp673[0] = 1
	max673[0] = 1
	for i := 1; i < len(nums); i++ {
		//	当前的数不大于前面的数，不能在当前的基础上加1
		// 类似dp[0] 起码是等于1的。
		dp673[i] = 1
		max673[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp673[i] < dp673[j]+1 {
					dp673[i] = dp673[j] + 1
					max673[i] = max673[j]
				} else if dp673[i] == dp673[j]+1 {
					max673[i] += max673[j]
				}
			}
		}
		maxNum673 = getMax(maxNum673, dp673[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		if dp673[i] == maxNum673 {
			ans += max673[i]
		}
	}
	return ans

	/*
		dp673[i] 对应nums[i] 这个数结尾的最长递增的序列的长度
		max673[k]表示 对应nums[i] 最长递增子序列的个数，及能达到这个长度的所有的路径的个数
		max673[k] 怎么由 前面的项推算出来呢
		1.当前数在dp中比前一个数大，不需要处理，证明是一条路径算出来的 max673[k] = max673[k-1]
		2. dp当前的数不大于前面一个比如 12 3 34 那是两个3组成的两条路径能到3，那就是 max673[k] = max673[k-1] +1
		3. 突发发现 规律1 有问题。比如 4前面有2个3， 4得判断比他小的3 是不是 两条路都能走到4
			有可能有一个长度为3的尾部是比长度为4这个数大的，所以4得判断比他小的所有的3个数，然后是不是每个3都能走到4，有几个算几个
		所以逻辑应该是 max673[k] =  dp[k-1]值的个数,且是这些值对应的nums数小于nums[k]
	*/

	//https://leetcode.cn/problems/number-of-longest-increasing-subsequence/solutions/1007341/gong-shui-san-xie-lis-de-fang-an-shu-wen-obuz/

}

func TestAl673(t *testing.T) {
	//fmt.Println(findNumberOfLIS([]int{1, 2, 5}))
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	//fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 5}))
	//fmt.Println(findNumberOfLIS([]int{1, 2, 2, 2, 2, 3, 4}))
	//[10,55,45,25,25]
}

//334
// 718、最长重复子数组（HJ75. 公共子串计算）
//LeetCode 1143、最长公共子序列
//
//LeetCode 188、买卖股票的最佳时机 IV
//LeetCode 121、买卖股票的最佳时机
//LeetCode 123、买卖股票的最佳时机 III
//LeetCode 122、买卖股票的最佳时机 II
//LeetCode 309、最佳买卖股票时机含冷冻期
//LeetCode 714、买卖股票的最佳时机含手续费

/*
第四十六天（周四）
LeetCode 120、三角形最小路径和
LeetCode 343、整数拆分
LeetCode 279、完全平方数
LeetCode 174、地下城游戏
第四十七天（周五）
LeetCode 5、最长回文子串
LeetCode 53、最大子数组和
LeetCode 516、最长回文子序列
LeetCode 718、最长重复子数组
第四十八天（周六）
LeetCode 322、零钱兑换
LeetCode 139、单词拆分
LeetCode 264、丑数II（动态规划）
LeetCode 72、编辑距离
打家劫舍问题
LeetCode 198、打家劫舍
LeetCode 213、打家劫舍II
LeetCode 337、打家劫舍III
第四十九天（周日）
动态规划（背包DP）
0-1背包
LeetCode 474、一和零
LeetCode 494、目标和
LeetCode 416、分割等和子集
完全背包
LeetCode 518、零钱兑换II
第五十天（周一）
什么是字典树
LeetCode 208、实现Trie（前缀树）
LeetCode 211、添加与搜索单词 - 数据结构设计
LeetCode 648、单词替换
LeetCode 676、实现一个魔法字典
第五十一天（周二）
什么是并查集
LeetCode 547、省份数量
LeetCode 200、岛屿数量（并查集解法）
第五十二天（周三）
LeetCode 445、两数相加II
LeetCode 1049、最后一块石头的重量II
LeetCode 322、零钱兑换（完全背包解法）
LeetCode 811、子域名访问次数
第五十三天（周四）
LeetCode 1109、航班预订统计
LeetCode 45、跳跃游戏II
LeetCode 376 、摆动序列
LeetCode 946、验证栈序列
第五十四天（周五）
LeetCode 231、2 的幂
LeetCode 268、丢失的数字
LeetCode 85、最大矩形
LeetCode 64、最小路径和
第五十五天（周六）
LeetCode 34、在排序数组中查找元素的第一个和最后一个位置
LeetCode 35、搜索插入位置
LeetCode 153、寻找旋转排序数组中的最小值
LeetCode 154、寻找旋转排序数组中的最小值 II
第五十六天（周日）
LeetCode 461、汉明距离
LeetCode 295、数据流中的中位数
LeetCode 297 、二叉树的序列化与反序列化
LeetCode 110、平衡二叉树
*/
