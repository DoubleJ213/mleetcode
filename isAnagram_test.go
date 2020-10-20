package mleetcode

import (
	"testing"
)

func isAnagram(s string, t string) bool {
	total := make(map[rune]int)
	for _, c := range s {
		total[c]++
	}
	for _, c := range t {
		total[c]--
	}
	for _, n := range total {
		if n != 0 {
			return false
		}
	}
	return true
}

func isAnagram1(s string, t string) bool {
	a := [26]int{}
	b := [26]int{}

	for _, v := range s {
		a[v-'a'] += 1
	}

	for _, v := range t {
		b[v-'a'] += 1
	}

	return a == b
}

func TestIsAnagram(t *testing.T) {
	strA := "cat"
	strB := "tac"
	isAnagram(strA, strB)
	isAnagram1(strA, strB)

}
