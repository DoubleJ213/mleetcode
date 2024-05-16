package fxxk

import (
	"fmt"
	"testing"
)

func longestPalindrome(s string) string {
	length := len(s)
	if len(s) < 2 {
		return s
	}

	maxLen := 1
	begin := 0
	for i := 0; i < length-1; i++ {
		//左边界
		for j := i + 1; j < length; j++ {
			//右边界
			if j-i+1 > maxLen && isPalinDrome(s, i, j) {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	//字符串截取
	return s[begin : maxLen+begin]
}

func isPalinDrome(s string, i, j int) bool {
	//判断是否是回文字符串
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func TestA(t *testing.T) {
	//test1 := "abcabacd"
	//fmt.Println(longestPalindrome(test1))
	//fmt.Println(longestPalindrome("aacabdkacaa"))
	fmt.Println(longestPalindrome("abcdecba"))
	//test2 := "asca"
	//fmt.Println(longestPalindrome(test2))
}

func longestPalindrome2(s string) string {
	length := len(s)
	if len(s) < 2 {
		return s
	}
	maxLen := 0
	begin := 0
	for i := 0; i < length-1; i++ {
		//i为中心位置
		//奇数
		oddLen := getMaxLen(s, i, i)
		//偶数
		evenLen := getMaxLen(s, i, i+1)
		moreLen := 0
		if oddLen > evenLen {
			moreLen = oddLen
		} else {
			moreLen = evenLen
		}
		if moreLen > maxLen {
			maxLen = moreLen
			begin = i - (maxLen-1)/2
		}
	}

	return s[begin : maxLen+begin]
}

func getMaxLen(s string, i, j int) int {
	index := 0
	maxLen := 0
	isOdd := true
	if i != j {
		isOdd = false
	}
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			break
		}
		i--
		j++
		index++
	}
	if isOdd {
		//奇数 aba
		maxLen = 2*index - 1
	} else {
		//偶数
		maxLen = 2 * index
	}
	return maxLen
}

// 其实可以不用翻转，收尾元素都是已知的，直接比较就行
func longestPalindrome3(s string) string {
	//第一遍遍历，先翻转原来字符串
	//然后遍历第二个字符串，每次遍历索引+1，依次和第一个字符串去比较，一直找到一样的。返回一样的长度
	//循环套循环，时间复杂度高。有没有更简单的方法？
	length := len(s)
	str := make([]byte, length)
	left := 0
	right := length - 1
	for right >= 0 {
		str[left] = s[right]
		left++
		right--
	}
	//i 为翻转后的索引， j为原来数组的索引
	//index 为当前正在遍历的翻转后的位置
	resp := ""
	max := 0
	//原来 bbaca
	//翻转 acabb 97 99 97 98 98  // bbdac
	index := 0
	for i := 0; i+index < length; i++ {
		//	翻转后的字符串
		for j := 0; j < length; j++ {
			//	翻转前的字符串
			if i+index >= length {
				index = 0
				break
			} else if str[i+index] == byte(s[j]) {
				if index >= max {
					max = index + 1
					resp = string(str[i : i+max])
				}
				index++
				//} else if j < length {
				//	continue
			} else {
				if j < length && index == 0 {
					continue
				} else {
					index = 0
					break
				}
			}
		}
		index = 0
	}

	return resp
}

//fmt.Println(longestPalindrome("aacabdkacaa"))
//写完发现理解错题目鸟，回文子串得是对称的  上面按照我的思路得到aaca  aaca翻转得到acaa 本身不等于 aaca
//只是出入里面碰巧有aaca 等找到子串相等。

// 重新写
func longestPalindrome4(s string) string {
	//左右指针，之前都是从两边向中间逼近，其实可以考虑从中间向两边逼近
	//然后回文串，奇数和偶数逻辑是不一样的
	//奇数一个中间点，偶数两个中间点
	resp := ""
	for i := 0; i < len(s); i++ {
		oddRes := palindrome(i, i, s)
		if len(oddRes) > len(resp) {
			resp = oddRes
		}
		evenRes := palindrome(i, i+1, s)
		//比长短，返回更长的
		if len(evenRes) > len(resp) {
			resp = evenRes
		}
	}
	return resp
}

func palindrome(l int, r int, s string) string {
	//	给定中间点，奇数l=r 偶数 l = r+1
	//	找最长的回文串
	length := 0
	left := l
	right := r
	for l >= 0 && r <= len(s)-1 {
		//	l 一直减  r 一直加
		if s[l] == s[r] {
			length++
			if l-1 >= 0 && r+1 <= len(s)-1 {
				l--
				r++
			} else {
				break
			}
		} else {
			break
		}
	}
	return s[left-(length-1) : right+length]
}

//测试用例
//"abcdecba"
//"aacabdkacaa"
//"bbaca"
//"bbdac"
//"babad"
//"cbbd"
//"a"
