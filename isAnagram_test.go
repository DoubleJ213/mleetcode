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

func TestIsAnagram(t *testing.T) {
	strA := "cat"
	strB := "tac"
	isAnagram(strA, strB)

}
