package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 全排列不行，有很多操作可以剪枝。
// 比如[3,1][2,1]数组肯定优先选[3,1]同样的类型，选利润大的
// 另外同样的利润，选择类型不一致的。

func findMaximumElegance(items [][]int, k int) int64 {
	// 把利润从大到小排序
	//slices.SortFunc(items, func(a, b []int) int { return b[0] - a[0] })
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] > items[j][0]
	})
	ans, totalProfit := 0, 0
	visited := make(map[int]bool)
	duplicate := make([]int, 0) // 重复类别的利润
	for i, p := range items {
		profit, category := p[0], p[1]
		if i < k {
			totalProfit += profit // 累加前 k 个项目的利润
			if !visited[category] {
				visited[category] = true
			} else { // 重复类别
				duplicate = append(duplicate, profit)
			}
		} else if len(duplicate) > 0 && !visited[category] { // 之前没有的类别
			visited[category] = true                            // len(vis) 变大
			totalProfit += profit - duplicate[len(duplicate)-1] // 选一个重复类别中的最小利润替换
			duplicate = duplicate[:len(duplicate)-1]
		} // else：比前面的利润小，而且类别还重复了，选它只会让 totalProfit 变小，len(vis) 不变，优雅度不会变大
		ans = getMax(ans, totalProfit+len(visited)*len(visited)) // 有可能被替换，也有可能虽然平方变大，还不如原来的ans大
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumEleganceOfAKLengthSubsequence(t *testing.T) {
	/*
		解释： 在这个例子中，我们需要选出长度为 3 的子序列。
		其中一种方案是 items[0] = [3,1] ，items[2] = [2,2] 和 items[3] = [5,3] 。
		子序列的总利润为 3 + 2 + 5 = 10 ，	子序列包含 3 种不同类别 [1, 2, 3] 。
		因此，优雅度为 10 + 3^2 = 19 ，可以证明 19 是可以获得的最大优雅度。
	*/
	//fmt.Println(findMaximumElegance([][]int{{5, 1}, {6, 1}, {3, 1}, {2, 2}, {4, 2}, {5, 3}}, 3))
	fmt.Println(findMaximumElegance([][]int{{2, 1}, {2, 1}, {2, 1}, {1, 2}, {1, 3}}, 3))
}

var ans2813 int64
var visited2813 []bool

/*
items[i] = [profit i, category i]
优雅度 可以用 total_profit + distinct_categories^2 计算
你的任务是从 items 所有长度为 k 的子序列中，找出 最大优雅度 。
注意：数组的子序列是经由原数组删除一些元素（可能不删除）而产生的新数组，且删除不改变其余元素相对顺序。
计算k个的所有优雅度的可能性，然后取最大值？
*/
// 1 <= k  <= items.length == n
func findMaximumElegance2(items [][]int, k int) int64 {
	//不难发现，全排列cpu、内存可能都会超限
	ans2813 = 0
	visited2813 = make([]bool, len(items))
	traverse2813(items, make([][]int, 0), 0, 0, k)
	return ans2813
}

func traverse2813(items, path [][]int, start, level, k int) {
	if k == level {
		//	找到层数为k的，处理数据
		ans2813 = getMaxInt64(ans2813, getAns(path))
	}
	for i := start; i < len(items); i++ {
		if visited2813[i] {
			continue
		}
		if level+1 > k {
			break
		}
		visited2813[i] = true
		traverse2813(items, append(path, items[i]), i+1, level+1, k)
		visited2813[i] = false
	}
}

func getMaxInt64(a int64, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func getAns(items [][]int) int64 {
	//profit  category
	//total_profit + distinct_categories^2
	m := make(map[int]bool)
	var pre int64
	var post int
	for _, item := range items {
		pre += int64(item[0])
		if _, ok := m[item[1]]; !ok {
			m[item[1]] = true
			post++
		}
	}
	return pre + int64(post*post)
}
