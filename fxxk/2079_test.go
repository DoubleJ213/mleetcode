package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
/*
输入：plants = [2,2,3,3], capacity = 5
输出：14
解释：
从河边开始，此时水罐是装满的：
- 走到植物 0 (1 步) ，浇水。水罐中还有 3 单位的水。
- 走到植物 1 (1 步) ，浇水。水罐中还有 1 单位的水。
- 由于不能完全浇灌植物 2 ，回到河边取水 (2 步)。
- 走到植物 2 (3 步) ，浇水。水罐中还有 2 单位的水。
- 由于不能完全浇灌植物 3 ，回到河边取水 (3 步)。
- 走到植物 3 (4 步) ，浇水。
需要的步数是 = 1 + 1 + 2 + 3 + 3 + 4 = 14 。
*/
func wateringPlants(plants []int, capacity int) int {
	//第一种方法就傻算，然后看看会不会超时
	//输入已经限制了 max(plants[i]) <= capacity <= 10^9
	var ans int
	remain := capacity
	for i := 0; i < len(plants); i++ {
		need := plants[i]
		if remain >= need {
			remain -= need
			ans++
		} else {
			remain = capacity
			ans += i << 1
			i--
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestWateringPlants(t *testing.T) {
	//plants = [2,2,3,3], capacity = 5
	fmt.Println(wateringPlants([]int{2, 2, 3, 3}, 5))

}
