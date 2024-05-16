package fxxk

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
序列化和反序列化 二叉搜索树
只需确保二叉搜索树可以序列化为字符串
并且可以将该字符串反序列化为最初的二叉搜索树。
编码的字符串应尽可能紧凑。
*/
type Codec449 struct {
	TextM string
	TextP string
}

func Constructor449() Codec449 {
	return Codec449{"", ""}
}

// Serializes a tree to a single string.
func (this *Codec449) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	this.TextM = ""
	this.submOrder(root)
	this.subpOrder(root)
	return fmt.Sprintf("%s'%s", this.TextM, this.TextP)
}

// 中序遍历
func (this *Codec449) submOrder(root *TreeNode) {
	if root == nil {
		return
	}
	this.submOrder(root.Left)
	if this.TextM == "" {
		this.TextM = fmt.Sprintf("%d", root.Val)
	} else {
		this.TextM = fmt.Sprintf("%s,%d", this.TextM, root.Val)
	}

	this.submOrder(root.Right)
	return
}

// 前序遍历
func (this *Codec449) subpOrder(root *TreeNode) {
	if root == nil {
		return
	}
	if this.TextP == "" {
		this.TextP = fmt.Sprintf("%d", root.Val)
	} else {
		this.TextP = fmt.Sprintf("%s,%d", this.TextP, root.Val)
	}
	this.subpOrder(root.Left)
	this.subpOrder(root.Right)
	return
}

// Deserializes your encoded data to tree.
func (this *Codec449) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	trees := strings.Split(data, "'")
	if len(trees) != 2 {
		return nil
	}
	mArr := strings.Split(trees[0], ",")
	pArr := strings.Split(trees[1], ",")
	return this.buildTree(mArr, pArr)
}

func getVal(arr []string, index int) int {
	val, _ := strconv.Atoi(arr[index])
	return val
}

func (this *Codec449) buildTree(mArr, pArr []string) *TreeNode {
	if len(mArr) == 0 && len(pArr) == 0 {
		return nil
	}
	var index int
	for i := 0; i < len(mArr); i++ {
		if getVal(pArr, 0) == getVal(mArr, i) {
			index = i
			break
		}
	}
	left := make([]string, index)
	right := make([]string, len(mArr)-1-index)
	copy(left, mArr[:index])
	copy(right, mArr[index+1:])

	leftTree := this.buildTree(left, pArr[1:len(left)+1])
	rightTree := this.buildTree(right, pArr[len(left)+1:])

	return &TreeNode{Val: getVal(pArr, 0), Left: leftTree, Right: rightTree}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestSerializeAndDeserializeBst(t *testing.T) {
	//root6 := &TreeNode{1, nil, nil}
	root5 := &TreeNode{4, nil, nil}
	root4 := &TreeNode{2, nil, nil}
	root3 := &TreeNode{6, nil, nil}
	root2 := &TreeNode{3, root4, root5}
	root := &TreeNode{5, root2, root3}

	ser := Constructor449()
	deser := Constructor449()
	str := ser.serialize(root)
	ans := deser.deserialize(str)
	fmt.Println(ans)
}
