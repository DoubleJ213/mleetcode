package fxxk

import (
	"fmt"
	"strings"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)

/*
'M' 金属 'P' 纸 'G' 玻璃
先算前缀和，方便后面判断某个字母如果需要走到当前位置，需要几分钟
然后判断garbage数组，每个里面有几个同字母，算时间？
进一步，再优化一下。从后往前找GMP的字母。这样就确定了，每个车子走路花的事件。
然后统计每个字母个数
G index=3 [2, 2+4=6, 2+4+3=9] 花了9分钟走路
4四分钟收集垃圾 13分钟？ bingoo
*/

func garbageCollection(garbage []string, travel []int) int {
	if len(travel) == 0 {
		return 0
	}
	sum := make([]int, len(travel))
	sum[0] = travel[0]
	for i := 1; i < len(travel); i++ {
		sum[i] = sum[i-1] + travel[i]
	}
	var pNum, mNum, gNum, count int
	for j := len(garbage) - 1; j >= 0; j-- {
		//	开始找垃圾走到的最终位置
		if j > 0 && pNum == 0 && strings.Contains(garbage[j], "P") {
			pNum = sum[j-1]
		}
		if j > 0 && gNum == 0 && strings.Contains(garbage[j], "G") {
			gNum = sum[j-1]
		}
		if j > 0 && mNum == 0 && strings.Contains(garbage[j], "M") {
			mNum = sum[j-1]
		}
		count += len(garbage[j])
	}
	return count + pNum + gNum + mNum
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumAmountOfTimeToCollectGarbage(t *testing.T) {
	//garbage = ["MMM","PGM","GP"], travel = [3,10]
	//fmt.Println(garbageCollection([]string{"MMM", "PGM", "GP"}, []int{3, 10}))
	fmt.Println(garbageCollection([]string{"G", "M"}, []int{1}))
	//	["G","M"]
	//	[1]
}
