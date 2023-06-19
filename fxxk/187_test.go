package fxxk

import (
	"fmt"
	"testing"
)

/*
187. 重复的DNA序列

DNA序列由一系列核苷酸组成，缩写为'A','C','G'和'T'。
例如，"ACGAATTCCG"是一个 DNA序列 。
在研究 DNA 时，识别 DNA 中的重复序列非常有用。
给定一个表示 DNA序列 的字符串 s ，返回所有在 DNA 分子中出现不止一次的长度为10的序列(子字符串)。你可以按 任意顺序 返回答案。

示例 1：
输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
输出：["AAAAACCCCC","CCCCCAAAAA"]
示例 2：
输入：s = "AAAAAAAAAAAAA"
输出：["AAAAAAAAAA"]

提示：
0 <= s.length <= 105
s[i]=='A'、'C'、'G' or 'T'
*/

/*
findRepeatedDnaSequences
长度为10的序列
字符只有4种
*/
func findRepeatedDnaSequences(s string) []string {
	//本题用10进制存窗口的数字。可以优化使用4进制就可以了
	if len(s) <= 10 {
		return nil
	}
	left, right := 0, 0
	var window int64
	windowMap := make(map[int64]int)
	result := make([]string, 0)

	for right < 10 {
		window = window*10 + getValue(s[right])
		right++
	}
	windowMap[window] = 1

	for right = 10; right < len(s); right++ {
		newWindow := (window-getValue(s[left])*1000000000)*10 + getValue(s[right])
		if value, ok := windowMap[newWindow]; ok {
			//存在一样的字符串
			if value == 1 {
				result = append(result, s[left+1:right+1])
			}
			windowMap[newWindow]++
		} else {
			windowMap[newWindow] = 1
		}
		left++
		window = newWindow
	}
	return result

	//Rabin - Karp 算法
}

/*
getValue
一共4个字符，'A'、'C'、'G' or 'T' 来一次转换
A--1 C--2 G--3 T--4
*/
func getValue(char byte) int64 {
	var value int64
	switch char {
	case 'A':
		value = 1
	case 'C':
		value = 2
	case 'G':
		value = 3
	case 'T':
		value = 4
	}
	return value
}

func TestAl187(t *testing.T) {
	/*
		uint8	无符号 8位整型 (0 到 255)
		uint16	无符号 16位整型 (0 到 65535)
		uint32	无符号 32位整型 (0 到 4294967295)
		uint64	无符号 64位整型 (0 到 18446744073709551615)
		int8	有符号 8位整型 (-128 到 127)
		int16	有符号 16位整型 (-32768 到 32767)
		int32	有符号 32位整型 (-2147483648 到 2147483647)
		int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)
		uint	32位操作系统上就是uint32，64位操作系统上就是uint64
		int	32位操作系统上就是int32，64位操作系统上就是int64
		uintptr	无符号整型，用于存放一个指针
	*/
	fmt.Printf("%v", findRepeatedDnaSequences("AAAAAAAAAAAAA"))
}
