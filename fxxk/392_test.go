package fxxk

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	x, y := 0, 0
	for x < len(s) {
		for y < len(t) {
			if x < len(s) && s[x] == t[y] {
				x++
			}
			y++
		}
		if y >= len(t) {
			break
		}
	}
	if x == len(s) {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIsSubsequence(t *testing.T) {

}
