package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 记录每个字母组成的最长连续的线段
func maximumLength3(s string) int {
	groups := [26][]int{}
	count := 0
	for i := range s {
		count++
		if i+1 == len(s) || s[i] != s[i+1] {
			groups[s[i]-'a'] = append(groups[s[i]-'a'], count) // 统计连续字符长度
			count = 0
		}
	}

	ans := 0
	for _, group := range groups {
		if len(group) == 0 {
			continue
		}
		sort.Slice(group, func(i, j int) bool {
			return group[i] > group[j]
		})
		group = append(group, 0, 0) // 假设还有两个空串
		ans = findMax(ans, group[0]-2, getMin(group[0]-1, group[1]), group[2])
	}

	if ans == 0 {
		return -1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func maximumLength2(s string) int {
	ans := -1
	listMap := make(map[rune][]int) //key 字符串 value 该字符出现字符串位置的索引的列表
	for index, ss := range s {
		if _, ok := listMap[ss]; ok {
			listMap[ss] = append(listMap[ss], index)
		} else {
			listMap[ss] = []int{index}
		}
	}
	for _, list := range listMap {
		if len(list) < 2 {
			continue
		}
		winMax := len(list) - 2 //长度为5 最大为2，然后看是不是满足2，不满足只能是1了1不用判断
		if winMax < ans {
			//最大可能性都小于当前的记录，直接别判断了
			continue
		}
		var current int
		if ans == -1 {
			current = getMaxLen2981(list, 0)
		} else {
			current = getMaxLen2981(list, ans)
		}

		if ans < current {
			ans = current
		}
	}
	return ans
}

// 寻找 至少出现三次的 连续索引长度 wide 当前已有的最长的长度
func getMaxLen2981(list []int, wide int) int {
	max := -1
	//winMax := len(list) - 2    //因为之前已经判断 数组长度3 或者3以上
	for win := wide + 1; win < len(list)-2; win++ { // win 是假设的字符串长度。 比如4位的看有没有3串，连续的索引
		need := 3 //一旦找到三个就可以返回当前窗口长度了
		find := false
		for left := 0; left+win-1 < len(list) && need > 0; left++ {
			for right := left + win - 1; right > left; right-- {
				if list[right-1]+1 != list[right] {
					find = false
					break
				}
				find = true
			}
			if find {
				need--
			}
		}
		if need == 0 {
			max = win
			continue
		} else {
			break
		}
	}
	return max
}

func TestFindLongestSpecialSubstringThatOccursThriceI(t *testing.T) {
	//fmt.Println(maximumLength("abcaba"))
	fmt.Println(maximumLength3("abcabaabcabaabcabaabcabaabcaba"))
	//fmt.Println(maximumLength("aaaa"))
	//fmt.Println(maximumLength("abcdef"))

}
