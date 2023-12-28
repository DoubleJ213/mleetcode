package fxxk

import (
	"fmt"
	"strings"
	"testing"
)

/*
648. 单词替换
在英语中，我们有一个叫做 词根(root) 的概念，
可以词根后面添加其他一些词组成另一个较长的单词——我们称这个词为 继承词(successor)。
例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。
现在，给定一个由许多词根组成的词典 dictionary 和一个用空格分隔单词形成的句子 sentence。
你需要将句子中的所有继承词用词根替换掉。如果继承词有许多可以形成它的词根，则用最短的词根替换它。
你需要输出替换之后的句子。

示例 1：
输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"
示例 2：
输入：dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
输出："a a b c"

提示：
1 <= dictionary.length <= 1000
1 <= dictionary[i].length <= 100
dictionary[i] 仅由小写字母组成。
1 <= sentence.length <= 10^6
sentence 仅由小写字母和空格组成。
sentence 中单词的总量在范围 [1, 1000] 内。
sentence 中每个单词的长度在范围 [1, 1000] 内。
sentence 中单词之间由一个空格隔开。
sentence 没有前导或尾随空格。
*/

/*
词根后面添加其他一些词组成另一个较长的单词
sentence 空格分隔
*/
func replaceWords(dictionary []string, sentence string) string {
	words := strings.Split(sentence, " ")
	ans := ""
	for _, word := range words {
		find := false
		tmp := ""
		for i := 0; i < len(dictionary); i++ {
			if strings.HasPrefix(word, dictionary[i]) {
				find = true
				current := dictionary[i]
				if tmp == "" || len(current) < len(tmp) {
					tmp = current
				}
			}
		}
		if ans != "" {
			tmp = " " + tmp
			word = " " + word
		}
		if find {
			ans += tmp
		} else {
			ans += word
		}
	}
	return ans
}

type Trie648 struct {
	Val byte
	// 表示是否插入过,true表示插入过这个单词
	// 比如插入 apple 搜索app，其实是没有插入的
	//startsWith 是否存在这样的前缀
	//search 判断是否插入过
	Status    bool
	SearchRes bool
	Prefix    string
	Children  []*Trie648
}

func Constructor648() Trie648 {
	//65～90为26个大写英文字母
	//97～122号为26个小写英文字母
	//本题只有小写英文字母
	return Trie648{
		Val:       90,
		Status:    false,
		SearchRes: false,
		Children:  nil,
	}
}

func (this *Trie648) Insert(word string) {
	//word 剩余的字符串
	dfs648(this, word, 0)
	return
}

func (this *Trie648) GetPrefix(word string) string {
	this.Prefix = ""
	this.subPrefix(this, word, 0)
	return this.Prefix
}

func (this *Trie648) subPrefix(root *Trie648, word string, deep int) {
	if deep == len(word) {
		return
	}
	w := word[deep]
	var isIn648 bool
	for j := 0; j < len(root.Children); j++ {
		if root.Children[j].Val == w {
			isIn648 = true
			this.Prefix += string(w)
			if root.Children[j].Status {
				return
			}
			this.subPrefix(root.Children[j], word, deep+1)
		}
	}
	if !isIn648 {
		this.Prefix = word
		return
	}
}

func dfs648(root *Trie648, word string, deep int) {
	isEnd := false
	if len(word)-1 == deep {
		isEnd = true
	}
	if deep == len(word) {
		return
	}
	if root == nil {
		return
	}
	if root.Children == nil {
		me := &Trie648{word[deep], isEnd, false, "", nil}
		root.Children = []*Trie648{me}
		dfs648(root.Children[0], word, deep+1)
	} else {
		var isIn208 bool
		var level *Trie648
		for i := 0; i < len(root.Children); i++ {
			level = root.Children[i]
			if level.Val == word[deep] {
				isIn208 = true
				if isEnd {
					level.Status = isEnd
				}
				break
			}
		}
		if !isIn208 {
			me := &Trie648{word[deep], isEnd, false, "", nil}
			level = me
			root.Children = append(root.Children, me)
		}
		// 处理字符串存在，但是没有插入过
		dfs648(level, word, deep+1)
	}
}

func replaceWords2(dictionary []string, sentence string) string {
	ans648 := ""
	mt := Constructor648()
	for i := 0; i < len(dictionary); i++ {
		mt.Insert(dictionary[i])
	}

	words := strings.Split(sentence, " ")
	for j := 0; j < len(words); j++ {
		prefix := mt.GetPrefix(words[j])
		ans648 += prefix + " "
	}
	return ans648[:len(ans648)-1]
}

func TestAl648(t *testing.T) {
	//fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
	fmt.Println(replaceWords2([]string{"aa"},
		"ab aa aaa aaaaa abaa a ba"))
}

/*
最短的词根

输入
dictionary =
[ "aa", "aaa", "a","aaaa"]
sentence =
"a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"

添加到测试用例
输出
"a aa a aaaa aaa aaa aaa aaaa bbb baba a"
预期结果
"a a a a a a a a bbb baba a"
*/
