/**
<p>编写一个函数来查找字符串数组中的最长公共前缀。</p>

<p>如果不存在公共前缀，返回空字符串&nbsp;<code>""</code>。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>strs = ["flower","flow","flight"]
<strong>输出：</strong>"fl"
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>strs = ["dog","racecar","car"]
<strong>输出：</strong>""
<strong>解释：</strong>输入不存在公共前缀。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
 <li><code>1 &lt;= strs.length &lt;= 200</code></li>
 <li><code>0 &lt;= strs[i].length &lt;= 200</code></li>
 <li><code>strs[i]</code> 仅由小写英文字母组成</li>
</ul>

<div><div>Related Topics</div><div><li>字典树</li><li>字符串</li></div></div><br><div><li>👍 3101</li><li>👎 0</li></div>
*/

package fxxk

import (
	"fmt"
	"strings"
	"testing"
)

var ans14 string

func longestCommonPrefix(strs []string) string {
	ans14 = ""
	for x := 0; x < len(strs[0]); x++ {
		pre := strs[0][x]
		for i := 1; i < len(strs); i++ { //pre就是从index=0 取的没必要判断
			if x >= len(strs[i]) {
				return ans14
			}
			cur := strs[i][x]
			if pre == cur {
				continue
			} else {
				return ans14
			}
		}
		ans14 = strings.Join([]string{ans14, string(pre)}, "")
	}
	return ans14
}

func TestLongestCommonPrefix(t *testing.T) {
	fmt.Println(longestCommonPrefix([]string{"rbace", "racecar", "ra"}))
	//fmt.Println(longestCommonPrefix([]string{"flower", "flower", "flower", "flower"}))
}
