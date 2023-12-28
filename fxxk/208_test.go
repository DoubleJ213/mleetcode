package fxxk

import (
	"fmt"
	"testing"
)

/*
208. 实现 Trie (前缀树)
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，
用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
请你实现 Trie 类：
Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，
返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix)
如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

示例：
输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True

提示：
1 <= word.length, prefix.length <= 2000
word 和 prefix 仅由小写英文字母组成
insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次
*/

/*
Trie树的基本性质：
根节点不包含字符，除根节点外的每一个子节点都包含一个字符。
从根节点到某一个节点，路径上经过的字符连接起来，为该节点对应的字符串。
每个节点的所有子节点包含的字符互不相同。
从第一字符开始有连续重复的字符只占用一个节点，比如上面的to，和ten，中重复的单词t只占用了一个节点。
*/
type Trie208 struct {
	Val byte
	// 表示是否插入过,true表示插入过这个单词
	// 比如插入 apple 搜索app，其实是没有插入的
	//startsWith 是否存在这样的前缀
	//search 判断是否插入过
	Status    bool
	SearchRes bool
	Children  []*Trie208
}

func Constructor208() Trie208 {
	//65～90为26个大写英文字母
	//97～122号为26个小写英文字母
	//本题只有小写英文字母
	return Trie208{
		Val:       90,
		Status:    false,
		SearchRes: false,
		Children:  nil,
	}
}

func (this *Trie208) Insert(word string) {
	//word 剩余的字符串
	dfs208(this, word, 0)
	return
}

func dfs208(root *Trie208, word string, deep int) {
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
		me := &Trie208{word[deep], isEnd, false, nil}
		root.Children = []*Trie208{me}
		dfs208(root.Children[0], word, deep+1)
	} else {
		var isIn208 bool
		var level *Trie208
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
			me := &Trie208{word[deep], isEnd, false, nil}
			level = me
			root.Children = append(root.Children, me)
		}
		// 处理字符串存在，但是没有插入过
		dfs208(level, word, deep+1)
	}
}

func (this *Trie208) Search(word string) bool {
	this.SearchRes = false
	if word == "" {
		return true
	}
	if this == nil {
		return false
	}
	this.subSearch208(this, word, 0, false)
	return this.SearchRes
}

func (this *Trie208) subSearch208(root *Trie208, word string, deep int, isSub bool) {
	if root == nil {
		return
	}
	if word == "" {
		return
	}
	isEnd := false
	if len(word)-1 == deep {
		isEnd = true
	}
	children := root.Children
	if len(children) == 0 {
		return
	}

	var isIn208 bool
	var got *Trie208
	for i := 0; i < len(children); i++ {
		current := children[i]
		if current.Val == word[deep] {
			got = current
			isIn208 = true
			break
		}
	}
	_ = isEnd
	if !isIn208 {
		return
	} else {
		if got == nil {
			return
		}
		if isEnd {
			if isSub {
				this.SearchRes = true
			} else {
				this.SearchRes = got.Status
			}
			return
		}
		this.subSearch208(got, word, deep+1, isSub)
	}
	return
}

func (this *Trie208) StartsWith(prefix string) bool {
	this.SearchRes = false
	if prefix == "" {
		return true
	}
	if this == nil {
		return false
	}
	this.subSearch208(this, prefix, 0, true)
	return this.SearchRes
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

/*
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True
*/
func TestAl208(t *testing.T) {
	/*Trie trie = new Trie();
	trie.insert("apple");
	trie.search("apple");   // 返回 True
	trie.search("app");     // 返回 False
	trie.startsWith("app"); // 返回 True
	trie.insert("app");
	trie.search("app");     // 返回 True*/
	fmt.Printf("Constructor +\n")
	myTrie := Constructor208()
	myTrie.Insert("apple")
	//myTrie.Insert("app")
	//myTrie.Insert("bpp")
	//fmt.Println(myTrie.Search("bpp"))
	fmt.Println(myTrie.Search("apple1"))
	fmt.Println(myTrie.Search("app"))
	fmt.Println(myTrie.StartsWith("app"))
	myTrie.Insert("app")
	fmt.Println(myTrie.Search("app"))
}
