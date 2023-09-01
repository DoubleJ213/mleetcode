package fxxk

import (
	"fmt"
	"testing"
)

/*
139. 单词拆分
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

示例 1：
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
示例 2：
输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
     注意，你可以重复使用字典中的单词。
示例 3：
输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false

提示：
1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s 和 wordDict[i] 仅由小写英文字母组成
wordDict 中的所有字符串 互不相同
*/

/*
这题实际是 遍历一棵高度为 N + 1 的满 N 叉树（N 为 wordDict 的长度）可以画出递归的n叉树

*/

var res139 bool

func wordBreak(s string, wordDict []string) bool {
	//由于输入已经限制了，这里就不各种非空乱七八糟的判断了
	res139 = false
	backtrack139(s, "", wordDict)
	return res139
}

func backtrack139(input, current string, dict []string) {
	if len(input) < len(current) {
		return
	}
	for i := 0; i < len(dict); i++ {
		if current == input {
			res139 = true
			break
		} else {
			backtrack139(input, fmt.Sprintf("%s%s", current, dict[i]), dict)
		}
	}
}

//存在侥幸心理提交了下，满树递归果然是超时了，先看看剪枝能不能优化下
func wordBreak1(s string, wordDict []string) bool {
	//由于输入已经限制了，这里就不各种非空乱七八糟的判断了
	res139 = false
	backtrack139_1(s, "", wordDict)
	return res139
}

func backtrack139_1(input, current string, dict []string) {
	if len(input) < len(current) {
		return
	}
	//除了长度相等的校验外
	//也提前校验下，当前current的长度是否和input一样
	if current != input[:len(current)] {
		//这次没有提交，但执行时间3s+，貌似还是超时，还得继续优化
		return
	}

	for i := 0; i < len(dict); i++ {
		if current == input {
			res139 = true
			break
		} else {
			backtrack139_1(input, fmt.Sprintf("%s%s", current, dict[i]), dict)
		}
	}
}

/*比如输入 wordDict = ["a", "aa"], s = "aaab"
算法无法找到一个可行的组合，所以一定会遍历整棵回溯树，这里面会存在重复的情况
因为s 是由 wordDict 组成的，那必然s的开头是有wordDict的某个单词组成的。
一旦找到s[0,i]这个前缀是由谁组成的，就不用再尝试用别的字符串再来拼前缀
接着处理s[i+1，end] 后面的长度就好. 这个就是「分解问题」的思维模式
那代码怎么写
*/
func wordBreak2(s string, wordDict []string) bool {
	return dp139(s, 0, wordDict)
}

func dp139(input string, index int, dict []string) bool {
	//	base case
	if index == len(input) {
		return true
	}

	for _, di := range dict {
		//遍历 wordDict，看看哪些单词是 s[i..] 的前缀
		//要想s[index..]能被拼出，其实是s[index+len(current..] 可以被拼出，
		//所以当 s[index+len(current..] 可以被拼出，s[index..] 就能被拼出
		if index+len(di) > len(input) {
			continue
		}
		if di == input[index:index+len(di)] {
			if dp139(input, index+len(di), dict) {
				return true
			}
		}

	}
	return false
}

/*
60ms左右，本以为能提交通过，结果又超时了。还得继续优化
因为整体上还是for循环里面有个递归
对于 wordBreak2 中的 dp139 其变化的状态是index，尝试通过添加备忘录的方法减少重复计算缩短时间复杂度
甚至是 设置合适的dp table 通过自底向上的循环遍历。
先尝试使用备忘录优化下.
比如 s = bccdbacdbdacddabbaaaadababadad 的例子
前缀时 b 可以匹配 然后 继续判断 ccdbacdbdacddabbaaaadababadad 是否能继续匹配
dp139[0:1] 是可以的 那要想知道后一个行不行，就继续算dp139[2:]
想了想 wordBreak2 可以有不同的理解。

dp问题的套路如下： 明确 base case -> 明确「状态」-> 明确「选择」 -> 定义 dp 数组/函数的含义。
拿零钱的问题类比下，字符串类似零钱总数，状态->index，选择 -> 不同的单词 类似不同面值的硬币

那外层循环就是遍历s index是变化的量，遍历到末尾的话，都能找到正确的选择那就是正确
dpRes[index] 的含义即表示 前index位置的字符串能被拼接出来，那要想知道dpRes[index]能不能被拼接出来
本单词前面某个位置是否能拼出来，然后 本单词长度是否 和对应字符串相等
dpRes[index] = dpRes[index-len(word)] && deRes[index-len(word): index] == word

想了想，原来这是自底向上遍历
下面的写法是逆着上面来的，看着正着遍历，但是写结果是反着来的
*/
var wordMap map[string]int
var dpRes []int

func wordBreak3(s string, wordDict []string) bool {
	wordMap = make(map[string]int)
	for _, word := range wordDict {
		wordMap[word] = 1
	}
	dpRes = make([]int, len(s))
	/*	默认值为0,
		当值为0认为当前的index没有计算过
		值为1 认为 true
		值为-1 认为 false*/
	return dp139_1(s, 0)
}

func dp139_1(input string, index int) bool {
	//	base case
	if index == len(input) {
		return true
	}

	fmt.Printf("%v \n", index)
	// 防止冗余计算
	if dpRes[index] != 0 {
		if dpRes[index] == 1 {
			return true
		}
		return false
	}

	for i := 0; i+index <= len(input); i++ {
		// 看看哪些前缀存在 wordDict 中
		prefix := input[index : index+i]
		if _, ok := wordMap[prefix]; ok {
			// 找到一个单词匹配 s[index..index+i)
			// 只要 s[index+index+i..] 可以被拼出，s[index..] 就能被拼出
			if dp139_1(input, index+i) {
				dpRes[index] = 1
				return true
			}
		}
	}

	// s[index..] 无法被拼出
	dpRes[index] = -1
	return false
}

func isOk(input, word string, index int) int {
	if input[index:index+len(word)] == word {
		return 1
	}
	return -1
}

func TestAL139(t *testing.T) {
	//fmt.Printf("true -- %v", wordBreak("s", []string{"a", "c", "s"}))
	//fmt.Printf("true -- %v", wordBreak("leetcode", []string{"leet", "code"}))
	//fmt.Printf("false -- %v", wordBreak2("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))

	/*fmt.Printf("true -- %v", wordBreak("bccdbacdbdacddabbaaaadababadad",
	[]string{"cbc", "bcda", "adb", "ddca", "bad", "bbb", "dad",
		"dac", "ba", "aa", "bd", "abab", "bb", "dbda", "cb",
		"caccc", "d", "dd", "aadb", "cc", "b", "bcc", "bcd",
		"cd", "cbca", "bbd", "ddd", "dabb", "ab", "acd", "a",
		"bbcc", "cdcbd", "cada", "dbca", "ac", "abacd", "cba",
		"cdb", "dbac", "aada", "cdcda", "cdc", "dbc", "dbcb",
		"bdb", "ddbdd", "cadaa", "ddbc", "babb"}))*/

	/*b := time.Now()
	fmt.Printf("true -- %v\n", wordBreak1("bccdbacdbdacddabbaaaadababadad",
		[]string{"cbc", "bcda", "adb", "ddca", "bad", "bbb", "dad",
			"dac", "ba", "aa", "bd", "abab", "bb", "dbda", "cb",
			"caccc", "d", "dd", "aadb", "cc", "b", "bcc", "bcd",
			"cd", "cbca", "bbd", "ddd", "dabb", "ab", "acd", "a",
			"bbcc", "cdcbd", "cada", "dbca", "ac", "abacd", "cba",
			"cdb", "dbac", "aada", "cdcda", "cdc", "dbc", "dbcb",
			"bdb", "ddbdd", "cadaa", "ddbc", "babb"}))
	fmt.Printf("%v\n", time.Since(b))
	*/

	//fmt.Printf("true -- %v", wordBreak3("s", []string{"a", "c", "s"}))
	//fmt.Printf("true -- %v", wordBreak3("leetcode", []string{"leet", "code"}))
	fmt.Printf("false -- %v", wordBreak3("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))

	//c := time.Now()
	//fmt.Printf("true -- %v\n", wordBreak3("bccdbacdbdacddabbaaaadababadad",
	//	[]string{"cbc", "bcda", "adb", "ddca", "bad",
	//		"bbb", "dad", "dac", "ba", "aa",
	//		"bd", "abab", "bb", "dbda", "cb",
	//		"caccc", "d", "dd", "aadb", "cc",
	//		"b", "bcc", "bcd", "cd", "cbca",
	//		"bbd", "ddd", "dabb", "ab", "acd",
	//		"a", "bbcc", "cdcbd", "cada", "dbca",
	//		"ac", "abacd", "cba", "cdb", "dbac",
	//		"aada", "cdcda", "cdc", "dbc", "dbcb",
	//		"bdb", "ddbdd", "cadaa", "ddbc", "babb"}))
	//fmt.Printf("%v\n", time.Since(c))

	//d := time.Now()
	//fmt.Printf("true -- %v\n", wordBreak2("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
	//	[]string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}))
	//fmt.Printf("%v\n", time.Since(d))

	//e := time.Now()
	//fmt.Printf("true -- %v\n", wordBreak2("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
	//	[]string{"aaaaaaaaaa","aaaaaaaaa","aaaaaaaa","aaaaaaa","aaaaaa","aaaaaa","aaaaa","aaaa","aaa","aa","a"}))
	//fmt.Printf("%v\n", time.Since(e))
}
