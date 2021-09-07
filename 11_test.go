package mleetcode

import (
	"fmt"
	"testing"
)

//11. 盛最多水的容器
//给你n个非负整数a1，a2，...，an，每个数代表坐标中的一个点(i,ai)。
//在坐标内画n条垂直线，垂直线i的两个端点分别为(i,ai)和(i,0)。
//找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
//说明：你不能倾斜容器。

//本题输出的是面积
//开始以为本题只是先找2个最高点，取低的值，然后乘以间距
//这个想法不对，应该还是以面积最大为准，穷举，依次求面积，配合一些剪枝逻辑，下面尝试写一下
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	//穷举，只是比较左右边界，没有过多剪枝逻辑
	var res int
	left, right := 0, 0
	for left = 0; left < len(height); left++ {
		for right = left + 1; right < len(height); right++ {
			//长度
			a := right - left
			//宽度
			var b int
			if height[left] > height[right] {
				b = height[right]
			} else {
				b = height[left]
			}
			tmp := a * b
			if res < tmp {
				res = tmp
			}
		}
	}
	//gg写完发现LeetCode不能通过，超时了，这个方法继续优化了
	return res
}

func maxArea1(height []int) int {
	if len(height) < 2 {
		return 0
	}
	var res int
	left, right := 0, 0
	for left = 0; left < len(height); left++ {
		//本层最大的数
		var levelMax int
		for right = len(height); right < left+1; right-- {
			if height[right] < levelMax {
				break
			}
			a := right - left
			var b int
			if height[left] > height[right] {
				b = height[right]
			} else {
				b = height[left]
			}
			tmp := a * b
			if res < tmp {
				res = tmp
			}
		}
	}
	//还是超时，继续优化，刚刚优化right从右边开始向左逼近，那能不能left同时向右逼近。
	//如果能，那将变成o(n)复杂度的了，铁定不会超时了
	return res
}

func maxArea2(height []int) int {
	if len(height) < 2 {
		return 0
	}
	var res int
	left, right := 0, len(height)-1
	//这个地方应该能发现，当left移到right右边了，可以不用遍历了，所以上面maxArea1其实可以优化下
	//道理明白了，那怎么改呢right一直在变化,想当然改成下面这个是有问题的，直接就0了
	//for left = 0; left < right; left++ {
	//怎么写，其实跟下面的这个写法是一样的
	//for left < right {
	//就不单独写出来，再优化了，直接写O(n)的
	//沿着刚刚maxArea1最后想的，两个指针同时像中间逼近，同时逼近对不对呢，必须对，为什么
	//TODO 补充
	for left < right {
		//left++
		//right--
		a := right - left
		var b int
		if height[left] > height[right] {
			b = height[right]
			right--
		} else {
			b = height[left]
			left++
		}
		tmp := a * b
		if res < tmp {
			res = tmp
		}
	}
	//再看看，还能不能优化
	return res
}

func TestAl11(*testing.T) {
	he := []int{1, 2, 3, 1}
	//fmt.Println(maxArea(he))
	//fmt.Println(maxArea1(he))
	fmt.Println(maxArea2(he))
	fmt.Println("done")
}
