package fxxk

import (
	"fmt"
	"testing"
)

/*
127. 单词接龙
字典 wordList 中从单词 beginWord 和 endWord 的 转换序列
是一个按下述规格形成的序列 beginWord -> s1 -> s2 -> ... -> sk：
每一对相邻的单词只差一个字母。
对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
sk == endWord
给你两个单词 beginWord 和 endWord 和一个字典 wordList ，
返回 从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0 。


示例 1：
输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
输出：5
解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
示例 2：
输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
输出：0
解释：endWord "cog" 不在字典中，所以无法进行转换。


提示：
1 <= beginWord.length <= 10
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord、endWord 和 wordList[i] 由小写英文字母组成
beginWord != endWord
wordList 中的所有字符串 互不相同
*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	inWordList := make(map[string]int)
	size := len(beginWord)
	visited127 := make(map[string]int)
	for x := 0; x < len(wordList); x++ {
		cur := wordList[x]
		//如果值为-1 才能证明列表中存在，才是合理的变化，其余跳过继续遍历下一个
		inWordList[cur] = 1
	}
	if inWordList[endWord] != 1 {
		return 0
	}
	step := 1
	queue127 := make([]string, 0)
	queue127 = append(queue127, beginWord)

	for len(queue127) > 0 {
		toDo127 := make([]string, 0)
		//	queue127 新值
		for i := 0; i < len(queue127); i++ {
			current := queue127[i]
			//if visited127[current] == 1 {
			//	//等于1说明该字符串已经访问过，之前有更快的方法到这里
			//	continue
			//}
			if current == endWord {
				return step
			}
			for j := 0; j < size; j++ {
				// 比如beginWord 为 "hit" 其size个字母都能变化.
				// 小写英文字母 26个，这里剪枝处理下不然内存涨爆了
				for k := 97; k < 123; k++ {
					next := getNext127(current, j, byte(k))
					if next != current && visited127[next] != -1 {
						//不是当前字符串。且没有遍历过
						if inWordList[next] == 1 {
							//只有 在WordList中，才是正确的，其余肯定是不满足继续判断条件的
							toDo127 = append(toDo127, next)
						}
						//该字符串不满足要求，但是访问过了
						visited127[next] = -1
					}
				}
			}
		}
		queue127 = toDo127
		step++
	}
	return 0
}

func getNext127(current string, index int, word byte) string {
	return current[0:index] + string(word) + current[index+1:]
}

func TestAl127(t *testing.T) {
	//fmt.Println(ladderLength(
	//	"hit", "cog",
	//	[]string{"hot", "dot", "dog", "lot", "log", "cog"}))
	//hit -> hot -> dot/lot -> hot/lot/dog/dot/hot/log
	fmt.Println(ladderLength(
		"hit", "cog",
		[]string{"hot", "dot", "dog", "lot", "log"}))
}
