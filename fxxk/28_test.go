package fxxk

import (
	"fmt"
	"math"
	"testing"
)

/*
   28.找出字符串中第一个匹配项的下标
   给你两个字符串haystack和needle，请你在haystack字符串中找出needle字符串的第一个匹配项的下标（下标从0开始）。如果needle不是haystack的一部分，则返回-1。

   示例1：
   输入：haystack="sadbutsad",needle="sad"
   输出：0
   解释："sad"在下标0和6处匹配。
   第一个匹配项的下标是0，所以返回0。
   示例2：
   输入：haystack="leetcode",needle="leeto"
   输出：-1
   解释："leeto"没有在"leetcode"中出现，所以返回-1。

   提示：
   1<=haystack.length,needle.length<=104
   haystack和needle仅由小写英文字符组成
*/

/*
   strStr  haystack和needle仅由小写英文字符组成
第187偷懒用的10进制表示4个不同的字符，这题字符变多了，尝试下26进制表示

// number 的进制
int R = 4;
// 想在 number 的最低位添加的数字
int appendVal = 0~3 中的任意数字;
// 运算，在最低位添加一位
number = R * number + appendVal;

// number 的进制
int R = 4;
// number 最高位的数字
int removeVal = 0~3 中的任意数字;
// 此时 number 的位数
int L = ?;
// 运算，删除最高位数字
number = number - removeVal * R^(L-1);
*/
func strStr(haystack string, needle string) int {
	//a-0
	//b-1 ... 26
	if len(haystack) < len(needle) {
		return -1
	}
	var window int64
	for i := 0; i < len(needle); i++ {
		window = window*26 + int64(needle[i]-'a')
	}

	var tmp int64
	for j := 0; j < len(needle); j++ {
		tmp = tmp*26 + int64(haystack[j]-'a')
	}
	if tmp == window {
		return 0
	}

	left := 0
	for k := len(needle); k < len(haystack); k++ {
		char := haystack[left]
		left++
		tmp = (tmp-int64(math.Pow(26, float64(len(needle)-1)))*int64(char-'a'))*26 + int64(haystack[k]-'a')
		if tmp == window {
			return left
		}
	}

	return -1
}

func TestAl28(t *testing.T) {
	fmt.Printf("%v", strStr("aabcleetoc", "leeto"))
}
