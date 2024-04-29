package fxxk

import (
	"fmt"
	"github.com/mlcPractice/leetcode/editor/cn"
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
var map138 map[*test.Node]*test.Node

func copyRandomList(head *test.Node) *test.Node {
	//random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为 null 。
	//0 <= n <= 1000
	//第一遍先把next都连好。第二遍再补充random？
	getNext := head
	getRandom := head
	ans := &test.Node{Val: -1, Next: nil, Random: nil}
	tmp := ans
	map138 = make(map[*test.Node]*test.Node)
	for getNext != nil {
		cur := &test.Node{Val: getNext.Val, Next: nil, Random: nil}
		map138[getNext] = cur
		tmp.Next = cur
		getNext = getNext.Next
		tmp = tmp.Next
	}
	fillRandom := ans.Next
	for getRandom != nil {
		if node, ok := map138[getRandom.Random]; ok {
			fillRandom.Random = node
		} else {
			fillRandom.Random = nil
		}
		getRandom = getRandom.Next
		fillRandom = fillRandom.Next
	}
	return ans.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCopyListWithRandomPointer(t *testing.T) {
	//head = [[1,1],[2,1]]
	//输出：[[1,1],[2,1]]
	var root2 *test.Node
	root2 = &test.Node{
		Val:    2,
		Next:   nil,
		Random: nil,
	}
	root2.Random = root2
	root := &test.Node{
		Val:    1,
		Next:   root2,
		Random: root2,
	}
	fmt.Println(copyRandomList(root))
}
