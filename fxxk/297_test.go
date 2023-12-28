package fxxk

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
297. 二叉树的序列化与反序列化
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，
同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，
你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅LeetCode 序列化二叉树的格式。
你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

示例 1：
输入：root = [1,2,3,null,null,4,5]
输出：[1,2,3,null,null,4,5]
示例 2：
输入：root = []
输出：[]
示例 3：
输入：root = [1]
输出：[1]
示例 4：
输入：root = [1,2]
输出：[1,2]

提示：
树中结点数在范围 [0, 104] 内
-1000 <= Node.val <= 1000
*/

type Codec struct {
	strSlice []string
}

func Constructor297() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.traverse5(root)
	var data string
	for _, str := range this.strSlice {
		if data != "" {
			data = fmt.Sprintf("%s,%s", data, str)
		} else {
			data = fmt.Sprintf("%s", str)
		}

	}
	return data
}

func (this *Codec) traverse5(root *TreeNode) {
	if root == nil {
		this.strSlice = append(this.strSlice, "null")
		return
	}
	this.strSlice = append(this.strSlice, fmt.Sprintf("%d", root.Val))
	this.traverse5(root.Left)
	this.traverse5(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.strSlice = strings.Split(data, ",")
	return this.buildTree297()
}

func (this *Codec) buildTree297() *TreeNode {
	if this.strSlice == nil {
		return nil
	}
	rootVal := this.strSlice[0]
	this.strSlice = this.strSlice[1:]
	if rootVal == "null" {
		return nil
	}
	val, err := strconv.Atoi(rootVal)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	root := &TreeNode{Val: val}
	root.Left = this.buildTree297()
	//root.Left = this.buildTree297(slice[1:])
	//如果是类似这样的写法，slice不是每次递归都变，容易出问题
	//889题，那是我前面有个辅助函数，确保每次数组都变化，所以才对的
	root.Right = this.buildTree297()

	return root

}

func TestAl297(t *testing.T) {
	node7 := &TreeNode{Val: 7}
	node6 := &TreeNode{Val: 6}
	node5 := &TreeNode{Val: 5}
	//node8 := &TreeNode{Val: 8}
	//node4 := &TreeNode{Val: 4, Right: node8}
	//node4 := &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 3, Left: node6, Right: node7}
	node2 := &TreeNode{Val: 2, Left: nil, Right: node5}
	root := &TreeNode{Val: 1, Left: node2, Right: node3}

	ser := Constructor297()
	data := ser.serialize(root)
	fmt.Println(data)

	deSer := Constructor297()
	ans := deSer.deserialize(data)
	fmt.Println(ans)
	fmt.Println("done")
}
