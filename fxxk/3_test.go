package fxxk

import (
	"fmt"
	"strings"
	"testing"
)

// leetcode 3
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
// 滑动窗口
func lengthOfLongestSubstring2(s string) int {
	var length int
	var str string
	left := 0
	right := 0
	//当前最长的字符串
	str = s[left:right]
	for right = 0; right < len(s); right++ {
		if index := strings.IndexByte(str, s[right]); index != -1 {
			//存在
			left = left + index + 1
		}
		str = s[left : right+1]
		if len(str) > length {
			length = len(str)
		}
	}
	return length
}

func TestString(t *testing.T) {
	aaa := lengthOfLongestSubstring("asadbascs")
	fmt.Println(aaa)
}

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
