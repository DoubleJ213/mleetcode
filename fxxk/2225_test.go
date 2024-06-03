package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

/*
输入：matches = [[2,3],[1,3],[5,4],[6,4]]
输出：[[1,2,5,6],[]]
解释： 玩家 1、2、5 和 6 都没有输掉任何比赛。玩家 3 和 4 每个都输掉两场比赛。
因此，answer[0] = [1,2,5,6] 和 answer[1] = [] 。

返回一个长度为 2 的列表 answer ：
answer[0] 是所有没有输掉任何比赛的玩家列表。answer[1] 是所有恰好输掉 一场 比赛的玩家列表。
两个列表中的值都应该按 递增 顺序返回。
*/
// leetcode submit region begin(Prohibit modification and deletion)
func findWinners(matches [][]int) [][]int {
	loser := make(map[int]int)
	for _, match := range matches {
		lose := match[1]
		if count, ok := loser[lose]; ok {
			loser[lose] = count + 1
		} else {
			loser[lose] = 1
		}
	}
	winnerList := make([]int, 0)
	loserList := make([]int, 0)
	sort.Slice(matches, func(i, j int) bool { return matches[i][0] < matches[j][0] })
	for _, ma := range matches {
		win := ma[0]
		if _, ok := loser[win]; !ok {
			if len(winnerList) <= 0 || winnerList[len(winnerList)-1] != win {
				winnerList = append(winnerList, win)
			}
		}
		lose := ma[1]
		if count, ok := loser[lose]; ok {
			if count == 1 {
				loserList = append(loserList, lose)
			}
		}
	}
	if loserList != nil {
		sort.Ints(loserList)
	}
	return [][]int{winnerList, loserList}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindPlayersWithZeroOrOneLosses(t *testing.T) {
	fmt.Println(findWinners([][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}))
}
