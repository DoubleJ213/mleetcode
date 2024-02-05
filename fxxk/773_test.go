package fxxk

import (
	"fmt"
	"testing"
)

/*
773. 滑动谜题
在一个 2 x 3 的板上（board）有 5 块砖瓦，用数字 1~5 来表示, 以及一块空缺用 0 来表示。
一次 移动 定义为选择 0 与一个相邻的数字（上下左右）进行交换.
最终当板 board 的结果是 [[1,2,3],[4,5,0]] 谜板被解开。
给出一个谜板的初始状态 board ，返回最少可以通过多少次移动解开谜板，如果不能解开谜板，则返回 -1 。

示例 1：
输入：board = [[1,2,3],[4,0,5]]
输出：1
解释：交换 0 和 5 ，1 步完成
示例 2:
输入：board = [[1,2,3],[5,4,0]]
输出：-1
解释：没有办法完成谜板
示例 3:
输入：board = [[4,1,2],[5,0,3]]
输出：5
解释：
最少完成谜板的最少移动次数是 5 ，
一种移动路径:
尚未移动: [[4,1,2],[5,0,3]]
移动 1 次: [[4,1,2],[0,5,3]]
移动 2 次: [[0,1,2],[4,5,3]]
移动 3 次: [[1,0,2],[4,5,3]]
移动 4 次: [[1,2,0],[4,5,3]]
移动 5 次: [[1,2,3],[4,5,0]]

提示：
board.length == 2
board[i].length == 3
0 <= board[i][j] <= 5
board[i][j] 中每个值都 不同
*/

/*
最少需要多少次能解卡，这题很适合bfs
其次怎么判断是否移到最终状态，如果每次都循环判断是不是慢了点，有没有更快的？
先最直接的方法
board[0][0]=1
board[0][1]=2
board[0][2]=3
board[1][0]=4
board[1][1]=5
board[1][2]=0
1 - 0 0
2 - 0 1
3 - 0 2
4 - 1 0
5 - 1 1
0 - 1 2
*/

var endString773 string
var visited773 map[string]int

func slidingPuzzle(board [][]int) int {
	/*
		前提
		board.length == 2
		board[i].length == 3
		0 <= board[i][j] <= 5
		board[i][j] 中每个值都 不同
		几乎不用判断
	*/
	endString773 = "123450"
	visited773 = make(map[string]int)
	//0 1 2 这不就是3进制嘛, num为key key-1对应索引的三进制数，索引即对应value
	//这样算之后1 -> 0,2 -> 1,3 -> 2,4 -> 3,5 -> 4,0 -> 5
	//但是不太好表示，要不直接处理成字符串算了
	step773 := 0
	//存需要继续遍历的board [][]int
	queue773 := make([][][]int, 0)
	queue773 = append(queue773, board)
	// visited773存在 值为1就是之前判断过，本次递归层数高于之前的，后面遇到直接return
	first := arrayToString(board)
	if endString773 == first {
		return step773
	}
	//TODO 不是终点
	visited773[first] = 1
	step773++

	for len(queue773) > 0 {
		toDoList := make([][][]int, 0)
		for i := 0; i < len(queue773); i++ {
			curArr := queue773[i]
			curStr := arrayToString(curArr)
			if visited773[curStr] != 1 {
				visited773[curStr] = 1
			}
			//每个curArr 每次 移动 定义为选择 0 与一个相邻的数字（上下左右）进行交换.
			up, down, left, right := getNext773(curArr)
			//_, _, _, _ = getNext773(curArr)
			//0和某个位置的交换
			//如果为nil 证明不能交换，直接下一个
			var status bool
			if up != nil {
				status, toDoList = doNext773(up, toDoList)
				if status {
					return step773
				}
			}
			if down != nil {
				status, toDoList = doNext773(down, toDoList)
				if status {
					return step773
				}
			}
			if left != nil {
				status, toDoList = doNext773(left, toDoList)
				if status {
					return step773
				}
			}
			if right != nil {
				status, toDoList = doNext773(right, toDoList)
				if status {
					return step773
				}
			}

		}
		queue773 = toDoList
		step773++
	}
	return -1
}

func doNext773(nextArray [][]int, toDoList [][][]int) (bool, [][][]int) {
	upStr := arrayToString(nextArray)
	//当新的数组和目标数组一样，直接return
	if upStr == endString773 {
		return true, nil
	}
	//下面是新的数组不一致，如果新的数组之前没有遍历过，则加入待遍历列表
	if visited773[upStr] != 1 {
		//之前出现过，这里不继续加入到toDoList中了
		toDoList = append(toDoList, nextArray)
	} else {
		visited773[upStr] = 1
	}
	return false, toDoList
}

func getNext773(board [][]int) (up, down, left, right [][]int) {
	x, y := getXY773(board)
	//{1, 2, 3}, {4, 0, 5}
	if x > 0 {
		up = getArray773(board, x, y, x-1, y)
	}
	if x < 1 {
		down = getArray773(board, x, y, x+1, y)
	}
	if y-1 >= 0 {
		left = getArray773(board, x, y, x, y-1)
	}
	if y+1 <= 2 {
		right = getArray773(board, x, y, x, y+1)
	}
	//up, down, left, right 分别是0和上下左右的元素交换的新数组
	return up, down, left, right
}

//x0 y0 为0的坐标
func getArray773(board [][]int, x0, y0, xn, yn int) [][]int {
	ans := make([][]int, len(board))
	for i := 0; i < len(ans); i++ {
		ans[i] = make([]int, len(board[i]))
		for j := 0; j < len(ans[i]); j++ {
			ans[i][j] = board[i][j]
		}
	}
	ans[x0][y0], ans[xn][yn] = ans[xn][yn], ans[x0][y0]

	return ans
}

func getXY773(board [][]int) (int, int) {
	//这里就不判断board的合法性了，题目已经约束了
	//先找数字0在哪里
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	return 0, 0
}

func arrayToString(board [][]int) string {
	vKey := ""
	if len(board) > 0 {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				vKey = fmt.Sprintf("%s%d", vKey, board[i][j])
			}
		}
	}
	return vKey
}

func TestAl773(t *testing.T) {
	//[[1,2,3],[4,0,5]]
	//fmt.Println(slidingPuzzle([][]int{{1, 2, 3}, {4, 0, 5}}))
	//fmt.Println(slidingPuzzle([][]int{{4, 1, 2}, {5, 0, 3}}))
	fmt.Println(slidingPuzzle([][]int{{1, 2, 3}, {5, 4, 0}}))
	//	[4,1,2],[5,0,3]
}

/*
下面是别人的代码，他是每次把需要遍历的都加进数组，然后在遍历的时候再去判断是不是到了终止条件，比我写的简洁点
就不需要像我那样，在每个邻居节点里面判断是不是到达了，如果没到达加入待遍历的数组

同时在添加数组的时候遍历是不是存在过，如果存在过就不加入，如果不存在就加入，且设置遍历过
这样的结构清晰一些

		int step = 0;
        while (!q.isEmpty()) {

            int sz = q.size();
            for (int i = 0; i < sz; i++) {
                String cur = q.poll();
                // 判断是否达到目标局面
                if (target.equals(cur)) {
                    return step;
                }
                // 找到数字 0 的索引
                int idx = 0;
                for (; cur.charAt(idx) != '0'; idx++) ;
                // 将数字 0 和相邻的数字交换位置
                for (int adj : neighbor[idx]) {
                    String new_board = swap(cur.toCharArray(), adj, idx);
                    // 防止走回头路
                    if (!visited.contains(new_board)) {
                        q.offer(new_board);
                        visited.add(new_board);
                    }
                }
            }
            step++;
*/
