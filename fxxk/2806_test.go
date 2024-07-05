package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 最近 的 10 的 倍数 取整 非负
func accountBalanceAfterPurchase(purchaseAmount int) int {
	b := purchaseAmount / 10
	if purchaseAmount%10 >= 5 {
		b++
	}
	return 100 - b*10
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAccountBalanceAfterRoundedPurchase(t *testing.T) {
	//fmt.Println(accountBalanceAfterPurchase(25))
	fmt.Println(accountBalanceAfterPurchase(9))
}
