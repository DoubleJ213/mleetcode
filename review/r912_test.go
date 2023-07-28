package review

import (
	"fmt"
	"testing"
)

/*912. 排序数组
给你一个整数数组nums，请你将该数组升序排列。

示例 1：
输入：nums = [5,2,3,1]
输出：[1,2,3,5]
示例 2：
输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]

提示：
1 <= nums.length <= 5 * 104
-5 * 104 <= nums[i] <= 5 * 104
*/
func sortArray(nums []int) []int {
	qs(nums, 0, len(nums)-1)
	return nums
}

func qs(nums []int, left int, right int) {
	if left >= right {
		return
	}
	p := partition(nums, left, right)
	//for left+1 < right && nums[left] == nums[left+1] {
	//	left++
	//}
	//for right-1 > left && nums[right] == nums[right-1] {
	//	right--
	//}
	//缩小的是p-1 和p+1到两边的距离，不能写成上面的逻辑
	// 本身 p两边排的分别是小的和大的
	//比如左子树 ，如果从左边开始把一样的去掉，相当于后面有个更小的永远拍不到前面了
	//左子树，只能从最右边 p开始，把和p一样的值省略了，p左移
	//同样右子树，p开始 往右移动进行剪枝处理
	rightIndex := p - 1
	for rightIndex > left && nums[p] == nums[rightIndex] {
		rightIndex--
	}
	leftIndex := p + 1
	for leftIndex < right && nums[p] == nums[leftIndex] {
		leftIndex++
	}
	//不剪枝，超时，过不了
	qs(nums, left, rightIndex)
	qs(nums, leftIndex, right)
}

func partition(nums []int, left, right int) int {
	p := nums[right]
	i, j := left, right
	for i < j {
		for nums[i] <= p && i < right {
			i++
		}
		for nums[j] >= p && j > left {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	//	返回p的新的index
	return i
}

func TestAl912(t *testing.T) {
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}

//5, 1, 1, 2, 0, 0
//1,1,2,0,5
//1,1,2,0 -> 0,1,2,1
