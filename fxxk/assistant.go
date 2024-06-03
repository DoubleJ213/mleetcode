package fxxk

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Node589 struct {
	Val      int
	Children []*Node589
}

func findMax(a, b, c, d int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if d > max {
		max = d
	}
	return max
}

func getMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func getMax(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
