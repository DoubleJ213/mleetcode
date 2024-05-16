package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)

//有个特别说明的地方，装水瞬间完成，然后继续浇水。不用当前轮次浇水，等下一轮浇水

/*
minimumRefill
题目规定a 走左边 b 走右边
如果 Alice 和 Bob 到达同一株植物，那么当前水罐中水更多的人会给这株植物浇水。
如果他俩水量相同，那么 Alice 会给这株植物浇水。
其实谁浇水也不重要 返回两人浇灌所有植物过程中重新灌满水罐的 次数 。
灌水操作没有说得很具体，看测试用例，不需要回到哪来，也不需要记录走的路程，反正加一次灌水的操作次数就行
*/
func minimumRefill2(plants []int, capacityA int, capacityB int) int {
	var ans int
	if len(plants) == 1 {
		return ans
	}
	left := 0
	right := len(plants) - 1
	remainA := capacityA
	remainB := capacityB
	for left <= right {
		//写完发现，其实是不需要立即去重新灌水的，a不够，有可能b能浇完
		//那准确的，需要重新灌水的条件是啥？
		//fmt.Printf("在浇 left %d 需要水 %d A 剩余 %d\n", left, plants[left], remainA)
		if remainA >= plants[left] {
			remainA -= plants[left]
			if left == right {
				break
			} else {
				left++
			}
		} else {
			//fmt.Println("注意此轮A不够了")
			ans++
			remainA = capacityA
			if left == right {
				break
			} else {
				left++
			}
		}
		//fmt.Printf("在浇 right %d 需要谁为 %d B 剩余 %d\n", right, plants[right], remainB)
		if remainB >= plants[right] {
			remainB -= plants[right]
			if left == right {
				break
			} else {
				right--
			}
		} else {
			//fmt.Println("注意此轮B不够了")
			ans++
			remainB = capacityB
		}
	}

	if remainA == capacityA {
		ans--
	}
	if remainB == capacityB {
		ans--
	}
	return ans

}

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	var ans int
	n := len(plants)
	if n == 1 {
		return ans
	}
	left := 0
	right := n - 1
	remainA := capacityA
	remainB := capacityB
	//isOdd := n&1 == 1   //奇数 true
	mid := (n - 1) >> 1 //mid := (n-1-0)>>1 + 0
	var leftEnd, rightEnd int
	rightEnd = mid + 1
	if n&1 == 1 {
		leftEnd = mid - 1
	} else {
		leftEnd = mid
	}
	for left <= leftEnd {
		if remainA >= plants[left] {
			remainA -= plants[left]
		} else {
			remainA = capacityA - plants[left]
			ans++
		}
		left++
	}
	for right >= rightEnd {
		if remainB >= plants[right] {
			remainB -= plants[right]
		} else {
			remainB = capacityB - plants[right]
			ans++
		}
		right--
	}
	if n&1 == 1 {
		if remainA >= plants[mid] || remainB >= plants[mid] {
			return ans
		} else {
			return ans + 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestWateringPlantsIi(t *testing.T) {
	//fmt.Println(minimumRefill([]int{2, 2, 3, 3}, 3, 4))
	//fmt.Println(minimumRefill([]int{2, 2, 2, 3, 3}, 5, 5))
	//fmt.Println(minimumRefill([]int{2, 2, 3, 3}, 5, 5))
	//fmt.Println(minimumRefill([]int{2, 2, 5, 2, 2}, 5, 5))
	fmt.Println(minimumRefill([]int{722, 2, 241, 465, 910, 302, 472, 941, 469, 165, 913, 497, 983, 448, 607, 626, 672, 58, 338, 889, 387}, 1518, 1121))
}
