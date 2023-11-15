
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

