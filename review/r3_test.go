package review

import (
	"fmt"
	"testing"
)

//leetcode 3
/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。

提示：
0 <= s.length <= 5 * 104
s由英文字母、数字、符号和空格组成
*/

/*
lengthOfLongestSubstring
遇到子数组/子串相关的问题，运用滑动窗口算法：
1、什么时候应该扩大窗口？
2、什么时候应该缩小窗口？
3、什么时候应该更新答案？
*/
func lengthOfLongestSubstring(s string) int {
	//可以考虑数组，index为ascii码
	countMap := make(map[byte]int)
	left, right, size := 0, 0, 0
	for right < len(s) {
		char := s[right]

		//窗口右边开始往右扩
		right++

		//加入窗口
		countMap[char]++
		// 当窗口内存在重复元素即个数大于1
		for isInclude(countMap, char) {
			// 删除最左边的元素，左边窗口开始收缩
			countMap[s[left]]--
			left++
		}
		if right-left > size {
			size = right - left
		}
	}
	return size
}

func isInclude(count map[byte]int, s byte) bool {
	if value, ok := count[s]; ok {
		if value > 1 {
			return true
		}
	}
	return false
}

func TestString(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abbkeb"))
	//fmt.Println(lengthOfLongestSubstring("babkeb"))
}

/*
int left = 0, right = 0;

while (left < right && right < s.size()) {
    // 增大窗口
    window.add(s[right]);
    right++;

    while (window needs shrink) {
        // 缩小窗口
        window.remove(s[left]);
        left++;
    }
}

*/
