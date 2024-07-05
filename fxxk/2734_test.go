package fxxk

import (
	"fmt"
	"testing"
)

/*

接着将字符串中的每一个字符替换为英文字母表中的前一个字符。例如，'b' 用 'a' 替换，'a' 用 'z' 替换。

如果a变成z字典序会变大，其余情况都会变小
返回执行上述操作 恰好一次 后可以获得的 字典序最小 的字符串。
选择 s 的任一非空子字符串，可能是整个字符串，
*/
//leetcode submit region begin(Prohibit modification and deletion)
func smallestString(s string) string {
	ans := []byte(s)
	for i, c := range s {
		if c > 'a' {
			for ; i < len(ans) && ans[i] > 'a'; i++ {
				ans[i]--
			}
			return string(ans)
		}
	}
	//走到这里，所有字母都是a,修改最后一个a为z
	ans[len(ans)-1] = 'z'
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLexicographicallySmallestStringAfterSubstringOperation(t *testing.T) {
	//fmt.Println(smallestString("aa"))
	fmt.Println(smallestString("aaaaaaaa"))
	//fmt.Println(smallestString("aba"))
	//fmt.Println(smallestString("caba"))
	//fmt.Println(smallestString("cbaba"))
	//
}

func smallestString2(s string) string {
	//必须替换，非空,也可以是整个字符串
	//a不能随意替换，其余字母连续的情况下有一个算一个，换就完事,如果中途出现了a，替换停止
	//如果字母是a，先看是否不得不换a，如果其他有字符串已经换过了，那就不换a。
	//如果必须换，不管有多少个a，换最靠右边的a就行
	var ans string
	last := -1
	var start, end bool
	var i int
	for i = 0; i < len(s); i++ {
		if s[i] == 'a' {
			last = i
			if start {
				end = true
				ans = fmt.Sprintf("%s%s", ans, s[i:])
				break
			}
			ans = fmt.Sprintf("%s%s", ans, string(s[i]))
		} else {
			if !end {
				ans = fmt.Sprintf("%s%s", ans, string(s[i]-1))
				start = true
			}
		}
	}
	if !start {
		ans = fmt.Sprintf("%s%s%s", s[:last], "z", s[last+1:])
	}
	return ans
	//	写完发现 aaa..aaa会超时
}
