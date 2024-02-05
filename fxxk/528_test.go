package fxxk

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
528. 按权重随机选择
给你一个下标从0开始的正整数数组w ，其中w[i] 代表第 i 个下标的权重。
请你实现一个函数pickIndex，它可以随机地从范围[0, w.length-1]内（含0和w.length-1）选出并返回一个下标。
选取下标i的概率为w[i]/sum(w)。
例如，对于w=[1,3]，挑选下标0的概率为1/(1+3)=0.25（即，25%），而选取下标1的概率为3/(1+3)=0.75（即，75%）。

示例 1：
输入：
["Solution","pickIndex"]
[[[1]],[]]
输出：
[null,0]
解释：
Solution solution = new Solution([1]);
solution.pickIndex(); // 返回 0，因为数组中只有一个元素，所以唯一的选择是返回下标 0。
示例 2：
输入：
["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
[[[1,3]],[],[],[],[],[]]
输出：
[null,1,1,1,1,0]
解释：
Solution solution = new Solution([1, 3]);
solution.pickIndex(); // 返回 1，返回下标 1，返回该下标概率为 3/4 。
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 0，返回下标 0，返回该下标概率为 1/4 。

由于这是一个随机问题，允许多个答案，因此下列输出都可以被认为是正确的:
[null,1,1,1,1,0]
[null,1,1,1,1,1]
[null,1,1,1,0,0]
[null,1,1,1,0,1]
[null,1,0,1,0,0]
......
诸若此类。

提示：
1 <= w.length <= 104
1 <= w[i] <= 105
pickIndex将被调用不超过 104次
*/

/*
假设输入的权重数组是 w = [1,3,2,1]
首先是前缀和思路
加权  0_1___4__6_7
获得前缀和数组
然后生成随机数 比如是3
二分法 找大于等于这个随机数的最小元素索引
最后对这个索引减一（因为前缀和数组有一位索引偏移）
*/
type Solution528 struct {
	preSum []int
}

func Constructor528(w []int) Solution528 {
	preSum := make([]int, len(w)+1)
	preSum[0] = 0
	for i := 1; i < len(w)+1; i++ {
		preSum[i] = preSum[i-1] + w[i-1]
	}
	return Solution528{preSum: preSum}
}

func (this *Solution528) PickIndex() int {
	size := len(this.preSum)
	max := this.preSum[size-1]

	//时间戳做随机数种子,不然伪随机
	rand.Seed(time.Now().UnixNano())
	//左开右闭 +1 就变成 【1，max】闭区间
	target := rand.Intn(max) + 1
	//二分法 找大于等于这个随机数的最小元素索引
	//最后对这个索引减一（因为前缀和数组有一位索引偏移）
	if this.preSum[size-1] == target {
		return size - 2
	}
	left, right := 0, size-1
	//704题是元素存在的情况
	//元素不存在的情况。考虑使用左侧边界二分法或右侧边界二分法
	//[1,3,2,1]  0_1___4__6_7
	for left <= right {
		mid := (right-left)/2 + left
		//不要写成(right+left)/2防止溢出
		if this.preSum[mid] == target {
			right = mid
		} else if this.preSum[mid] < target {
			left = mid + 1
		} else if this.preSum[mid] > target {
			//	右指针不减一，因为要找最小的比target大的数
			right = mid
		}
	}
	return left - 1
}

func TestAl528(t *testing.T) {
	obj := Constructor528([]int{1, 3, 2, 1})
	fmt.Printf("PickIndex1: %d\n", obj.PickIndex())
	fmt.Printf("PickIndex2: %d\n", obj.PickIndex())
	fmt.Printf("PickIndex3: %v\n", obj.PickIndex())
	fmt.Printf("PickIndex4: %v\n", obj.PickIndex())
	fmt.Printf("PickIndex5: %v\n", obj.PickIndex())
	fmt.Printf("PickIndex6: %v\n", obj.PickIndex())
	fmt.Printf("PickIndex7: %v\n", obj.PickIndex())
}

/*
当目标元素 target 不存在数组 nums 中时，搜索左侧边界的二分搜索的返回值可以做以下几种解读：
1、返回的这个值是 nums 中大于等于 target 的最小元素索引。
2、返回的这个值是 target 应该插入在 nums 中的索引位置。
3、返回的这个值是 nums 中小于 target 的元素个数。
*/
