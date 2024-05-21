package fxxk

import (
	"fmt"
	"testing"
)

/*
每回合游戏都在数组的前两个元素（即 arr[0] 和 arr[1] ）之间进行。比较 arr[0] 与 arr[1] 的大小，
较大的整数将会取得这一回合的胜利并保留在位置 0 ，较小的整数移至数组的末尾。
当一个整数赢得 k 个连续回合时，游戏结束，该整数就是比赛的 赢家
题目数据 保证 游戏存在赢家。
2 <= arr.length <= 10^5
*/

//leetcode submit region begin(Prohibit modification and deletion)

func getWinner(arr []int, k int) int {
	max1535 := arr[0]
	win := 0 //表示对应max 胜利了几轮
	for i := 1; i < len(arr) && win < k; i++ {
		if arr[i] > max1535 {
			//前一个最大值，已经不是最大了
			max1535 = arr[i]
			win = 0
		}
		win++
	}
	return max1535

}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindTheWinnerOfAnArrayGame(t *testing.T) {
	//fmt.Println(getWinner([]int{3, 2, 4, 1}, 5))
	//fmt.Println(getWinner([]int{3, 2, 1}, 10))
	//fmt.Println(getWinner([]int{3, 2, 1}, 10))
	//fmt.Println(getWinner([]int{1, 9, 8, 2, 3, 7, 6, 4, 5}, 7))
	fmt.Println(getWinner([]int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99}, 1000000000))
}
