package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

/*
相比2981 本题数据范围变了，暴力会超时
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maximumLength(s string) int {
	wordMap := make(map[uint8][]int) // key 单词 value 长度 list
	var count int
	for i := range s {
		count++
		if i+1 == len(s) || s[i] != s[i+1] {
			// 和后一个字符不一样，或者已经到结尾了。统计长度
			if _, ok := wordMap[s[i]]; ok {
				wordMap[s[i]] = append(wordMap[s[i]], count)
			} else {
				wordMap[s[i]] = []int{count}
			}
			count = 0
		}
	}
	ans := 0
	for _, list := range wordMap {
		sort.Slice(list, func(i, j int) bool {
			return list[i] > list[j]
		})
		list = append(list, 0, 0)
		ans = findMax2982(ans, list[0]-2, getMin2982(list[0]-1, list[1]), list[2])
		//4 3 1
	}
	if ans == 0 {
		return -1
	}
	return ans
}

func getMin2982(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func findMax2982(a, b, c, d int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if d > max {
		max = d
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindLongestSpecialSubstringThatOccursThriceIi(t *testing.T) {
	fmt.Println(maximumLength("assssaasssdad"))
}
