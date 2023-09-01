package review

import (
	"fmt"
	"testing"
)

func reverse(node *ListNode) *ListNode {
	var res11 *ListNode
	for node != nil {
		tmp := node.Next
		node.Next = res11
		res11 = node
		node = tmp
	}
	return res11
}

func reverse2(node *ListNode) *ListNode {
	if node.Next == nil {
		return node
	}
	newHead := reverse2(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

//类似后序遍历，出栈的时候处理
//能不能递归前处理呢,写完可以发现，这个其实和for循环遍历很像
func reverse3(node *ListNode) *ListNode {
	return traverseListNode(nil, node)
}

func traverseListNode(pre, node *ListNode) *ListNode {
	if node == nil {
		return pre
	}
	next := node.Next
	node.Next = pre
	return traverseListNode(node, next)
}

func TestF1(t *testing.T) {
	/*tmp := &ListNode{-1, nil}
	root := tmp
	for i := 1; i < 7; i++ {
		list := &ListNode{
			Val:  i,
			Next: nil,
		}
		tmp.Next = list
		tmp = tmp.Next
	}
	//reverse(root.Next)
	reverse2(root.Next)
	fmt.Printf("123")
	//a := reverse3(root.Next)
	//fmt.Printf("123 %v", a)*/

	//fmt.Printf("%v", quickSort([]int{1, 6, 7, 8, 2, 3, 4}))
	fmt.Printf("%v", mergeSort([]int{1, 6, 7, 8, 2, 3, 4}))
	//a := []int{1, 6, 7, 8, 2, 3, 4}
	//fmt.Printf("%v", a[1:2])
}

/*
归并  快排
快排相当于二叉树的前序遍历，把当前的数找到合适的位置，然后再排他的左\右子树
归并相当于二叉树的后序遍历，左\右子树都排好了，然后合并结果
*/
func quickSort(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	qsfxxk(nums, 0, len(nums)-1)
	return nums
}

func qsfxxk(nums []int, left int, right int) {
	if left >= right {
		return
	}
	l := left
	r := right
	p := nums[right]
	for l < r {
		for nums[l] <= p && l < right {
			l++
		}
		for nums[r] >= p && r > left {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		} else {
			break
		}
	}
	nums[l], nums[right] = nums[right], nums[l]
	qsfxxk(nums, left, l-1)
	qsfxxk(nums, l, right)
}

//归并 相当于二叉树的后序遍历,左、右子树排好序后，合并起来得到最终数组
func mergeSort(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nums
	}
	return msFxxk(nums)
}

func msFxxk(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2
	/*	类似于 二分法
		二分法最重要的是如何找到合适的中间点，且遍历的时候不能丢掉元素
		如果找不准，可以用具体数字带入求公式
		mid := len(nums)/2
		lNums := msFxxk(nums[:mid])
		rNums := msFxxk(nums[mid:])*/
	lNums := msFxxk(nums[:mid+1])
	rNums := msFxxk(nums[mid+1:])
	return merge(lNums, rNums)
}

func merge(left, right []int) []int {
	res := make([]int, len(left)+len(right))
	i := 0
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			res[i] = left[l]
			l++
		} else {
			res[i] = right[r]
			r++
		}
		i++
	}
	copy(res[i:], left[l:])
	copy(res[i+len(left)-l:], right[r:])
	return res
}
