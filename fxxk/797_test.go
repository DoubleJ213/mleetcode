package fxxk

import (
	"fmt"
	"testing"
)

/*
797. 所有可能的路径
在图论中，如果一个有向图无法从某个顶点出发经过若干条边回到该点，则这个图是一个有向无环图（DAG图）

给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）
graph[i] 是一个从节点 i 可以访问的所有节点的列表（即从节点 i 到节点 graph[i][j]存在一条有向边）。

示例 1：
输入：graph = [[1,2],[3],[3],[]]
输出：[[0,1,3],[0,2,3]]
解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
示例 2：
输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]

提示：
n == graph.length
2 <= n <= 15
0 <= graph[i][j] < n
graph[i][j] != i（即不存在自环）
graph[i] 中的所有元素 互不相同
保证输入为 有向无环图（DAG）
*/
var result797 [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	path := []int{}
	result797 = [][]int{}
	traverse797(graph, 0, path)
	return result797
}

func traverse797(graph [][]int, start int, path []int) {
	path = append(path, start)
	if len(graph)-1 == start {
		//判断走到最后，这里有个疑问，虽然说不重复，但是没说会跳着来，图里面各个值顺序增加的貌似。
		//graph如果把4改成5  [[5,3,1],[3,2,5],[3],[5],[]]
		//报错'graph' must consist of values from 0 to 4 only
		//如果没有这个前提，得把全部已经遍历过的节点存起来判断
		//不能把引用类型 path 传递给result
		//否则后续result全是中的全是一样的值
		p := make([]int, len(path))
		copy(p, path)
		//这里只能值传递
		result797 = append(result797, p)
	}

	for _, v := range graph[start] {
		traverse797(graph, v, path)
	}
	path = path[:len(path)-1]

}

func TestAl797(t *testing.T) {
	fmt.Printf("%v", allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
}
