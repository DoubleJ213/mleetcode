package fxxk

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 32 位无符号

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		if (num>>i)&1 == 1 {
			ans += 1 << (31 - i)
		} else {
			ans += 0 << (31 - i)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseBits(t *testing.T) {
	//输入是一个长度为 32 的二进制字符串
	fmt.Println(reverseBits(43261596))
	//	00000010100101000001111010011100
	//	1928352384
}
