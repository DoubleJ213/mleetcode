
303、304
前缀和主要适用的场景是原始数组不会被修改的情况下，频繁查询某个区间的累加和。

1109、1094
差分数组，主要适用场景是频繁对原始数组的某个区间的元素进行增减。


⼆叉树解题的思维模式分两类：
1、是否可以通过遍历⼀遍⼆叉树得到答案？
如果可以，⽤⼀个 traverse 函数配合外部变量来实现，这叫 「遍历」的思维模式。 
2、是否可以定义⼀个递归函数，通过⼦问题（⼦树）的答案推导出原问题的答案？
如果可以，写出这个递归 函数的定义，并充分利⽤这个函数的返回值，这叫「分解问题」的思维模式。

快速排序就是个二叉树的前序遍历
归并排序就是⼆叉树的后序遍历

再刷一次
98、95



待刷




基础
二维数组初始化和赋值 509 322 518 1143
二叉树构造
链表构造
最大、小堆
stack
fifo

常刷问题
链表翻转-循环、递归
二叉树前、中、后序遍历
快排、归并排序
二分法
前缀和、差分数组
滑动窗口

dp 509(fib) 322 300(最长递增子序列)
组合/子集问题 、排列问题

元素无重不可复选
元素可重不可复选
元素无重可复选

https://labuladong.github.io/algo/di-ling-zh-bfe1b/hui-su-sua-56e11/



    public int coinChange(int[] coins, int amount) {

        int[][] dp = new int[coins.length+1][amount+1];
        // coin = 0  没有钱 
        for(int j = 1 ; j <=amount;j++){
            dp[0][j] = Integer.MAX_VALUE;
        }
        //amount = 0  要凑的钱 为 0 
        for(int i = 0; i<=coins.length;i++){
            dp[i][0] = 0;
        }
        
        for(int i = 1; i<=coins.length;i++){
            for(int j = 1; j<=amount;j++){
                //注意这里需要对 Interger.MAX_VALUE 这个边界初始值进行判断
                if( j >= coins[i-1] && dp[i][j-coins[i-1]]!=Integer.MAX_VALUE ){

                    dp[i][j] = Math.min(dp[i-1][j],dp[i][j-coins[i-1]]+1);
                }else{//面额大于 要凑的钱
                    dp[i][j] = dp[i-1][j];
                }

            }
        }
        
        return dp[coins.length][amount] == Integer.MAX_VALUE ? -1: dp[coins.length][amount];
    }
}


class Solution {
    // dp[i][j] 表示 将i个硬币装入背包 ， 剩余 需要凑的钱(j)是多少
    public int change(int amount, int[] coins) {
        
        int[][] dp = new int[coins.length+1][amount+1];
        //base case
        for(int i = 0; i<= coins.length;i++){
            dp[i][0] = 1;//base case  就是要凑 0 元 说明直接每种硬币都能拿到一种可能性 就是啥也不装
        }


        for(int i = 1; i<=coins.length; i++){
            for(int j = 1;j<=amount; j++){
                if(j >= coins[i-1]){//要凑的钱 比 当前的硬币面值大
                    int a = dp[i-1][j];//依赖于前一种  当前不拿硬币  就剩余要凑 j 了
                    int b = dp[i][j-coins[i-1]];//当前拿了硬币 装入硬币了  就剩余要凑 j-coins[i-1]
                    dp[i][j] = a+b;//因为要求的是总共多少可能 组合数 需要将两种情况的组合数相加
                }else{//剩余要凑的 j 元 比 当前的硬币面值还要小 说明凑不出来 只能依赖前一种
                    dp[i][j] = dp[i-1][j];
                }
            }
        }
        return  dp[coins.length][amount];
    }
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
698. 划分为k个相等的子集
600
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