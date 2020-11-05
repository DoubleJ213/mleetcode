package mleetcode

import "testing"

// 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		// 尾数为0，同样有问题 if x < 0 {
		return false
	}
	y := 0
	for x > y {
		y = y*10 + x%10
		x /= 10
	}
	return y == x || y/10 == x

}

func TestPalindrome(t *testing.T) {
	isPalindrome(12321)
}
