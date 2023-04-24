package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

var combRes [][]int

//var currRes []int

func combinationSum(candidates []int, target int) [][]int {
	/*
		树，root节点为target，然后不断用candidates中的某个值生成子树，即递归
		子树的限制:
		    1.从candidates中任意取一个小于root的值生成子树
			2.一个数可以一直用，一直生成子树，一直递归，递归到减去的值为0或者为负数停止
				可以考虑数组排序，一旦变成负数可以提前结束本子树递归
			3.一个数同层在不能再出现第二次，即同一层一个数只减一次
	*/
	//用于存最终结果，最终结果是每次减去的值
	combRes = make([][]int, 0)
	//记录已经减了哪些数了，[]int类型
	sort.Ints(candidates)
	//combDfs(candidates, make([]int, 0), target)
	combDfs(candidates, []int(nil), target)
	return combRes
}

func combDfs(candidates, currRes []int, target int) {
	/*
		调试后发现，currRes全局变量会被错误的赋值
		比如:一旦前一次初始化过currRes长度足够并对其中某一个索引位置赋值
		[1,1,1,1,1,3,2,2],[1,1,1,1,1,3,2]
		第一次 currRes [1,1,1,1,1,1,1,2]
		第二次 currRes [1,1,1,1,1,n,n,n] append [3,2]后
		第二次变成      [1,1,1,1,1,3,2]
		但是会改变第一次currRes 的值
		第一次的变成    [1,1,1,1,1,3,2,2]
		TODO:单独测试下，深入理解，记住
		数组是值传递，切片是引用传递
	*/

	if target == 0 {
		tmp := append([]int(nil), currRes...)
		combRes = append(combRes, tmp)
		return
	}

	for i, candidate := range candidates {
		if target < candidate {
			return
		}
		combDfs(candidates[i:], append(currRes, candidates[i]), target-candidate)
	}
}

func TestAl39(t *testing.T) {
	//[2,7,6,3,5,1]
	//9
	fmt.Println(combinationSum([]int{2, 7, 6, 3, 5, 1}, 9))
	fmt.Println(combinationSum2([]int{2, 7, 6, 3, 5, 1}, 9))
	fmt.Println("aaa")
}

var combRes2 [][]int

func combinationSum2(candidates []int, target int) [][]int {
	combRes2 = make([][]int, 0)
	var trcak []int
	backtracking(0, 0, target, candidates, trcak)
	return combRes2
}

func backtracking(startIndex, sum, target int, candidates, trcak []int) {
	//终止条件
	if sum == target {
		tmp := make([]int, len(trcak))
		copy(tmp, trcak)                 //拷贝
		combRes2 = append(combRes2, tmp) //放入结果集
		return
	}
	if sum > target {
		return
	}
	//回溯
	for i := startIndex; i < len(candidates); i++ {
		//更新路径集合和sum
		trcak = append(trcak, candidates[i])
		sum += candidates[i]
		//递归
		backtracking(i, sum, target, candidates, trcak)
		//回溯
		trcak = trcak[:len(trcak)-1]
		sum -= candidates[i]
	}

}
