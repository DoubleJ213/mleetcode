package fxxk

import (
	"fmt"
	"testing"
)

/*
合法路径 指的是图中任意一条从节点 0 开始，最终回到节点 0 ，
且花费的总时间 不超过 maxTime 秒的一条路径。你可以访问一个节点任意次。
一条合法路径的 价值 定义为路径中 不同节点 的价值 之和
每个节点 至多 有 四条 边与之相连

和3067 很像
*/
// leetcode submit region begin(Prohibit modification and deletion)

type tree2065 struct {
	next int // 下一个数
	time int // 耗时
}

var ans2065 int
var visited2065 []bool

// 输入：values = [1,2,3,4], edges = [[0,1,10],[1,2,11],[2,3,12],[1,3,13]], maxTime = 50
func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	// 将edges 展开成二维数组
	ans2065 = 0
	g := make([][]tree2065, len(values)) //节点从0开始
	for _, e := range edges {
		//x, y, t := , , e[2]
		g[e[0]] = append(g[e[0]], tree2065{e[1], e[2]})
		g[e[1]] = append(g[e[1]], tree2065{e[0], e[2]})
	}

	//图中任意一条从节点 0 开始，最终回到节点 0
	//for i := 0; i < len(g[0]); i++ { //遍历所有0的下一跳
	//从 当前 0 到每一个下一跳 g[0][i].next  花费时间 g[0][i].time
	visited2065 = make([]bool, len(values))
	visited2065[0] = true
	dfs2065(g, values, maxTime, 0, 0, values[0]) //values[i] 起点，肯定没被访问过，直接加上
	//哎，写作一个变量，values[0] 起点，肯定没被访问过，直接加上
	//}
	return ans2065
}

// edge 当前走到哪里了
func dfs2065(g [][]tree2065, values []int, maxTime int, cur, sumTime int, sumValue int) {
	if cur == 0 { //当前路过0了，算一下总价值
		ans2065 = getMax(ans2065, sumValue)
	}

	for _, gg := range g[cur] { //遍历当前edge的所有下一跳
		next, t := gg.next, gg.time
		//sumTime += t  第一遍这里这样写的，然后进入递归，回溯的时候忘记 -t了
		if sumTime+t > maxTime {
			continue //这里不能return，当前路径不合适，还得遍历后面的路径呢
		}
		if visited2065[next] {
			dfs2065(g, values, maxTime, next, sumTime+t, sumValue)
		} else {
			visited2065[next] = true
			dfs2065(g, values, maxTime, next, sumTime+t, sumValue+values[next])
			visited2065[next] = false
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumPathQualityOfAGraph(t *testing.T) {
	//[0,32,10,43]
	//	[[0,1,10],[1,2,15],[0,3,10]]
	//	49
	//[[0,1,10],[1,2,11],[2,3,12],[1,3,13]]
	fmt.Println(maximalPathQuality([]int{1, 2, 3, 4}, [][]int{{0, 1, 10}, {1, 2, 11}, {2, 3, 12}, {1, 3, 13}}, 50))
	//fmt.Println(maximalPathQuality([]int{0, 32, 10, 43}, [][]int{{0, 1, 10}, {1, 2, 15}, {0, 3, 10}}, 49))
}
