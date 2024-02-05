package fxxk

import (
	"fmt"
	"testing"
)

/*
752. 打开转盘锁
你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。

示例 1:
输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
输出：6
解释：
可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
因为当拨动到 "0102" 时这个锁就会被锁定。
示例 2:
输入: deadends = ["8888"], target = "0009"
输出：1
解释：把最后一位反向旋转一次即可 "0000" -> "0009"。
示例 3:
输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
输出：-1
解释：初始数字为 '0000' 无法旋转到目标数字且不被锁定。

提示：
1 <= deadends.length <= 500
deadends[i].length == 4
target.length == 4
target 不在 deadends 之中
target 和 deadends[i] 仅由若干位数字组成
*/

/*
锁有4位数，且每位数可以往大和往小调整一位。所以相当于每一种情况后面有8种情况
从0开始。第一次递归就是8种，如果8种中的情况种，有一种满足条件，这一种继续递归，又是8种
类似岛屿，走迷宫的题目，这之间判断下不要拨重复了。
比如 0000 往大 0001 往小0009 0009又可以往大0000拨。因为0000已经判断过，所以跳过判断，跳过这种case的递归
*/

var queue752 []string
var visited752 map[string]int

/*
这题不和111题，dfs简单。不太好计算 0000 扭到 9000的最小值。往高扭往低扭需要判断最小值
BFS 找到的路径一定是最短的，因为是一层一层往下遍历的。但代价就是空间复杂度可能比 DFS 大很多
*/
func openLock(deadends []string, target string) int {
	queue752 = make([]string, 0)
	step752 := 0
	first := "0000"
	visited752 = make(map[string]int)

	//不能扭到的设置为-1
	for i := 0; i < len(deadends); i++ {
		if deadends[i] == first {
			return -1
		}
		visited752[deadends[i]] = -1
	}

	// 从起点开始bfs
	visited752[first] = 1
	queue752 = append(queue752, first)

	for len(queue752) > 0 {
		toDoQueue := make([]string, 0)
		for i := 0; i < len(queue752); i++ {
			current := queue752[i]
			if visited752[current] == -1 {
				//死亡节点直接死亡
				continue
			}
			if current == target {
				return step752
			}
			// 当前节点需要继续遍历
			//	把子节点构造好，等待下一层遍历
			for j := 0; j < len(target); j++ {
				up := getUp752(current, j)
				if !isVisited752(up) {
					toDoQueue = append(toDoQueue, up)
					if visited752[up] != -1 {
						visited752[up] = 1
					}
				}

				down := getDown752(current, j)
				if !isVisited752(down) {
					toDoQueue = append(toDoQueue, down)
					if visited752[down] != -1 {
						visited752[down] = 1
					}
				}
			}

		}
		//TODO: queue752 赋新的值
		queue752 = toDoQueue
		step752++
	}
	return -1
}

func isVisited752(next string) bool {
	//是否访问过
	if _, ok := visited752[next]; ok {
		//存在这个key，那就判断是否为0
		if visited752[next] == 1 {
			return true
		}
	}
	//如果不存在肯定没遍历过
	return false
}

func getUp752(current string, index int) string {
	//0-48 1-49 2-50 3-51 4-52 5-53 6-54 7-55 8-56 9-57
	nc := current[index]
	if current[index] == 57 {
		nc = byte(48)
	} else {
		nc = current[index] + 1
	}
	return current[0:index] + string(nc) + current[index+1:]
}

func getDown752(current string, index int) string {
	nc := current[index]
	if current[index] == 48 {
		//0变小是9
		nc = byte(57)
	} else {
		nc = current[index] - 1
	}
	return current[0:index] + string(nc) + current[index+1:]
}

func TestAl752(t *testing.T) {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	fmt.Println(openLock([]string{"0000"}, "8888"))
}
