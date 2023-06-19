package fxxk

import (
	"fmt"
	"math"
	"testing"
)

/*
875. 爱吃香蕉的珂珂
珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有piles[i]根香蕉。警卫已经离开了，将在 h 小时后回来。
珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。
如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。

示例 1：
输入：piles = [3,6,7,11], h = 8
输出：4
示例 2：
输入：piles = [30,11,23,4,20], h = 5
输出：30
示例 3：
输入：piles = [30,11,23,4,20], h = 6
输出：23

提示：
1 <= piles.length <= 104
piles.length <= h <= 109
1 <= piles[i] <= 109
*/

/*
先确定最小吃香蕉的速度和最大吃香蕉的速度
最小吃香蕉速度，可以每次吃1根。
最大吃香蕉速度，因为一堆吃完，不会吃第二堆，所以最大速度为香蕉数量最多的香蕉数
然后找满足要求的最小吃香蕉速度，即左边界
*/
func minEatingSpeed(piles []int, h int) int {
	if h < len(piles) {
		return -1
	}
	left, right := 1, 1
	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}
	/*每次都做减运算，太耗时了
	eat := func(cap int) int {
		index := 0
		hour := 0
		remainPile := 0
		for remainPile > 0 || index <= len(piles)-1 {
			if remainPile <= 0 {
				remainPile = piles[index]
				index++
			}
			remainPile -= cap
			hour++
		}
		return hour
	}*/

	eat := func(cap int) int {
		hour := 0
		for _, pile := range piles {
			//剩余香蕉不够一次吃的也得花一小时吃，所以向上取整
			hour += int(math.Ceil(float64(pile) / float64(cap)))
		}
		return hour
	}
	for left <= right {
		//闭区间
		mid := left + (right-left)/2
		if eat(mid) == h {
			right = mid - 1
		} else if eat(mid) > h {
			left = mid + 1
		} else if eat(mid) < h {
			right = mid - 1
		}

	}
	return left
}

func TestAl875(t *testing.T) {
	fmt.Printf("%d \n ", minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Printf("%d \n ", minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	fmt.Printf("%d \n ", minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
	//fmt.Printf("%d \n ", minEatingSpeed([]int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184}, 823855818))
}
