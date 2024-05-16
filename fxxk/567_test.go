package fxxk

import (
	"fmt"
	"testing"
)

/*
567. 字符串的排列
给你两个字符串s1和s2 ，写一个函数来判断 s2 是否包含 s1的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。

示例 1：
输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
示例 2：
输入：s1= "ab" s2 = "eidboaoo"
输出：false

提示：
1 <= s1.length, s2.length <= 104
s1 和 s2 仅包含小写字母
*/

/*
	checkInclusion

先搞清楚题目意思
s2包含一个s1的全排列，就是s2的一个子串刚好和s1中的各字母数量一致。和顺序无关

由于排列不会改变字符串中每个字符的个数，
所以只有当两个字符串每个字符的个数均相等时，一个字符串才是另一个字符串的排列。
根据这一性质，记s1的长度为n，我们可以遍历s2中的每个长度为n的子串，
判断子串和s1中每个字符的个数是否相等，若相等则说明该子串是s1的一个排列。

又说都是小写字母，那就是可以认为数组长度26位。
0表示a 1表示b这样以此类推，这样就可以一个数组，能全部记录所有字符串的个数。
*/
func checkInclusion(s1 string, s2 string) bool {
	//s2包含s1的排列之一.
	if len(s2) < len(s1) {
		return false
	}
	countS1 := make([]int, 26)
	//这里可以整两个数组进行直接比较
	//当然算法复杂度一样的
	for i := 0; i < len(s1); i++ {
		countS1[s1[i]-'a']++
	}
	for j := 0; j < len(s1); j++ {
		countS1[s2[j]-'a']--
	}
	if isTrue(countS1) {
		return true
	}
	//	窗口左边的索引
	left := 0
	for k := len(s1); k < len(s2); k++ {
		//新元素加入
		countS1[s2[k]-'a']--
		//旧元素移除
		countS1[s2[left]-'a']++
		if isTrue(countS1) {
			return true
		}
		left++
	}

	return false
}

// 在golang中我们可以轻松地通过==来判断两个数组(array)是否相等
// 但遗憾的是slice并没有相关的运算符
// 当需要判断两个slice是否相等时我们只能另寻捷径了
func isTrue(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			return false
		} else {
			continue
		}
	}
	return true
}

func TestAl567(t *testing.T) {
	fmt.Printf("%v", checkInclusion("ab", "eidbaooo"))
}
