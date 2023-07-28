package fxxk

import (
	"fmt"
	"testing"
)

//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// DFS 前序遍历、中序遍历、后续遍历
func TestF(t *testing.T) {
	node7 := &TreeNode{Val: 7}
	node6 := &TreeNode{Val: 6}
	node5 := &TreeNode{Val: 5}
	//node8 := &TreeNode{Val: 8}
	//node4 := &TreeNode{Val: 4, Right: node8}
	node4 := &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 3, Left: node6, Right: node7}
	node2 := &TreeNode{Val: 2, Left: node4, Right: node5}
	root := &TreeNode{Val: 1, Left: node2, Right: node3}
	fmt.Println("递归前序遍历: ", preorderTraversal(root))
	fmt.Println("递归前序遍历: ", preorderTraversal1(root))
	fmt.Println("迭代前序遍历: ", preorderTraversal2(root))
	fmt.Println("迭代前序遍历: ", preorderTraversal5(root))
	fmt.Println("递归中序遍历: ", inorderTraversal(root))
	fmt.Println("迭代中序遍历: ", inorderTraversal1(root))
	fmt.Println("迭代中序遍历: ", inorderTraversal2(root))
	fmt.Println("递归后序遍历: ", postorderTraversal(root))
	fmt.Println("迭代后序遍历: ", postorderTraversal2(root))
	fmt.Println("迭代后序遍历: ", postorderTraversal4(root))
	fmt.Println("迭代后序遍历-破坏原的树: ", postorderTraversal1(root))
}

func postorderTraversal4(root *TreeNode) []int {
	resp := make([]int, 0)
	treeList := make([]*TreeNode, 0)
	var tmp *TreeNode
	for root != nil || len(treeList) > 0 {
		for root != nil {
			treeList = append(treeList, root)
			root = root.Left
		}
		root = treeList[len(treeList)-1]
		treeList = treeList[:len(treeList)-1]
		if root.Right == nil || root.Right == tmp {
			resp = append(resp, root.Val)
			tmp = root
			root = nil
			//  设置 root = nil 主要是为了处理 root.right不等 nil的情况
			//	root.right 不为空，但是 root.right = tmp 表示右子树已经遍历过
			//	加完根的值root置空，处理下一个
		} else {
			// 为了确保先遍历一遍右子树，再遍历根。得把root重新加回去，等下判断root是否遍历过就行
			treeList = append(treeList, root)
			root = root.Right
		}
	}
	return resp
}

func inorderTraversal2(root *TreeNode) *[]int {
	resp := make([]int, 0)
	tmp := make([]*TreeNode, 0)
	for root != nil || len(tmp) != 0 {
		for root != nil {
			tmp = append(tmp, root)
			root = root.Left
		}
		root = tmp[len(tmp)-1]
		tmp = tmp[:len(tmp)-1]
		resp = append(resp, root.Val)
		root = root.Right
	}
	return &resp
}

func preorderTraversal5(root *TreeNode) *[]int {
	resp := make([]int, 0)
	tmp := make([]*TreeNode, 0)
	for root != nil || len(tmp) != 0 {
		for root != nil {
			resp = append(resp, root.Val)
			tmp = append(tmp, root.Right)
			//tmp = append(tmp, root)
			root = root.Left
		}
		root = tmp[len(tmp)-1]
		//root = tmp[len(tmp)-1].right
		tmp = tmp[:len(tmp)-1]
	}
	return &resp
}

// 后序遍历 145
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	dfsPostorderTraversal(root, &res)
	return res
}

func dfsPostorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	dfsPostorderTraversal(root.Left, res)
	dfsPostorderTraversal(root.Right, res)
	*res = append(*res, root.Val)
	return
}

//前序遍历 144
//func preorderTraversal(root *TreeNode) {

func preorderTraversal(root *TreeNode) (data []int) {
	res := []int{}
	node := root
	if node == nil {
		return nil
	}
	res = append(res, node.Val)
	//preorderTraversal(node.Left)的元素被打散一个个append进res
	res = append(res, preorderTraversal(node.Left)...)
	res = append(res, preorderTraversal(node.Right)...)
	return res
}

func preorderTraversal1(root *TreeNode) (res []int) {
	dfsPreorderTraversal(root, &res)
	return res
}

func dfsPreorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	dfsPreorderTraversal(root.Left, res)
	dfsPreorderTraversal(root.Right, res)
}

//中序遍历 94
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	dfsInorderTraversal(root, &res)
	return res
}

func dfsInorderTraversal(root *TreeNode, i *[]int) {
	if root == nil {
		return
	}
	dfsInorderTraversal(root.Left, i)
	*i = append(*i, root.Val)
	dfsInorderTraversal(root.Right, i)
}

// 迭代法前序遍历
// 栈S;
// p= root;
// while(p || S不空){
// while(p){
// 访问p节点；
// p的右子树入S;
// p = p的左子树;
// }
// p = S栈顶弹出;
// }

type TreeStack struct {
	roots []*TreeNode
}

func (this *TreeStack) Pop() *TreeNode {
	if len(this.roots) == 0 {
		return nil
	}
	res := this.roots[len(this.roots)-1]
	this.roots = this.roots[:len(this.roots)-1]
	return res
}

func (this *TreeStack) Top() *TreeNode {
	return this.roots[len(this.roots)-1]
}

func (this *TreeStack) Push(root *TreeNode) {
	this.roots = append(this.roots, root)
}

func postorderTraversal1(root *TreeNode) []int {
	//后序遍历
	//利用栈的后进先出，分别对每一个节点 压入中，右，左，弹出栈顶获取其值，重复上述
	//所以每个节点遍历三次,先是被压入,弹出后 又压入 （压入中，右，左节点）
	//第三次是 当左右节点遍历完成后，再弹出 获取其值
	//例如root,首先被压入；然后弹出 并压入root，root.Right和root.Left；处理完root.Left后，再次弹出root节点获取其值
	//注意：因为弹出是左，右，中的顺序，导致中节点 弹出栈时，有可能再次压入中，右，左节点，所以需要标记已处理防止陷入死循环

	if root == nil {
		return nil
	}
	mTStack := &TreeStack{}
	//mTStack.roots = append(mTStack.roots, root)
	mTStack.Push(root)
	res := make([]int, 0)
	for len(mTStack.roots) > 0 {
		getStack := mTStack.Top()
		// 判断栈首是否为叶子节点,若是,则出栈,取出叶子节点值,加入到ret等等返回
		if getStack.Left == nil && getStack.Right == nil {
			res = append(res, getStack.Val)
			mTStack.Pop()
			continue
		}

		// 由于前一个判断已经确定非叶子节点,则将其左右子节点加入到栈中,并将其左右子节点指针赋空,避免进入循环加入
		if getStack.Right != nil {
			mTStack.Push(getStack.Right)
			getStack.Right = nil
		}
		if getStack.Left != nil {
			mTStack.Push(getStack.Left)
			getStack.Left = nil
		}
	}
	return res
}

func preorderTraversal2(root *TreeNode) []int {
	// 迭代前序遍历
	mTStack := &TreeStack{}
	res := make([]int, 0)
	for len(mTStack.roots) > 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			mTStack.Push(root.Right)
			root = root.Left
		}
		root = mTStack.Pop()
	}
	return res
}

func inorderTraversal1(root *TreeNode) []int {
	//中序遍历
	res := make([]int, 0)
	mTStack := &TreeStack{}
	for len(mTStack.roots) > 0 || root != nil {
		//root != nil 只为了第一次root判断，放后面
		for root != nil {
			mTStack.Push(root)
			//入栈
			root = root.Left //移至最左
		}
		node := mTStack.Pop()
		res = append(res, node.Val) //中序输出
		root = node.Right           //右节点会进入下次循环，如果 =nil，继续出栈
	}
	return res
}

func postorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	mTStack := &TreeStack{}
	var outNode *TreeNode
	for len(mTStack.roots) > 0 || root != nil {
		for root != nil {
			mTStack.Push(root)
			root = root.Left
			//	此处已确保没有左子树了，且都压入栈中
		}

		node := mTStack.Pop()
		if node.Right == nil || node.Right == outNode {
			//通过判断右子节点是否为空，或者右子节点是否被打印，来输出
			res = append(res, node.Val)
			outNode = node
		} else {
			mTStack.Push(node)
			root = node.Right
		}

	}
	return res
}
