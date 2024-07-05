package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 输入保证 edges 构成一棵合法的树。
// 那就构造出来一个树出来，但是不是二叉
// pre []int next []int

type edge3067 struct {
	next   int // 下一个数
	weight int // 权重
}

var count3067 int

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	/*
		2 <= n <= 1000
		edges.length == n - 1
		edges[i].length == 3
		0 <= ai, bi < n
		edges[i] = [ai, bi, weighti]
	*/
	res := make([]int, len(edges)+1)
	grid := make([][]edge3067, len(edges)+1)
	//有限个数的数字，构造二维数组，代替构造map，这个思路得熟悉
	for _, v := range edges {
		left, right, weight := v[0], v[1], v[2]
		grid[left] = append(grid[left], edge3067{right, weight})
		grid[right] = append(grid[right], edge3067{left, weight})
	}

	for current, nextList := range grid {
		sum := 0
		for _, n := range nextList {
			count3067 = 0
			dfs3067(grid, signalSpeed, current, n.next, n.weight)
			res[current] += count3067 * sum
			sum += count3067
		}
	}
	return res
}

func dfs3067(grid [][]edge3067, signalSpeed, cur, next, sum int) {
	if sum%signalSpeed == 0 {
		count3067++
	}
	for _, g := range grid[next] {
		if g.next != cur {
			dfs3067(grid, signalSpeed, next, g.next, sum+g.weight)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCountPairsOfConnectableServersInAWeightedTreeNetwork(t *testing.T) {
	//	edges = [[0,6,3],[6,5,3],[0,3,1],[3,2,7],[3,1,6],[3,4,2]], signalSpeed = 3
	fmt.Println(countPairsOfConnectableServers([][]int{{0, 6, 3}, {6, 5, 3},
		{0, 3, 1}, {3, 2, 7}, {3, 1, 6}, {3, 4, 2}}, 3))
}
