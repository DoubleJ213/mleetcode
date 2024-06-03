package fxxk

import (
	"fmt"
	"testing"
)

/*
输入：n = 3, limit = 3
输出：10
解释：总共有 10 种方法分配 3 颗糖果，
且每位小朋友的糖果数不超过 3 ：
(0, 0, 3) ，(0, 1, 2) ，(0, 2, 1) ，(0, 3, 0) ，
(1, 0, 2) ，(1, 1, 1) ，(1, 2, 0) ，
(2, 0, 1) ，(2, 1, 0) 和 (3, 0, 0)

*/
// leetcode submit region begin(Prohibit modification and deletion)
// n 糖果总数 limit最多几个 一共3个小孩
//容斥原理 所有可能的组合-不可能的

//leetcode submit region end(Prohibit modification and deletion)

func TestDistributeCandiesAmongChildrenI(t *testing.T) {
	fmt.Println(distributeCandies2928(3, 3))
}

var res2928 int

// n 糖果总数 limit最多几个 一共3个小孩  dfs这玩意儿就是容易超时，别的都挺好
func distributeCandies2928(n int, limit int) int {
	res2928 = 0
	dfs2928(n, limit, 0)
	return res2928
}

func dfs2928(target int, limit int, kids int) {
	if kids == 3 {
		if target == 0 {
			res2928++
		}
		return
	}
	for i := 0; i < limit+1; i++ {
		if target-i < 0 {
			break
		}
		dfs2928(target-i, limit, kids+1)
	}
	return
}
