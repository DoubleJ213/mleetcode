package mleetcode

import "testing"

//120. 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	return 0
}

func TestMinimumTotal(t *testing.T) {
	//[
	//	 [2],
	//	[3,4],
	// [6,5,7],
	//[4,1,8,3]
	//]

	a := [][]int{
		{2},    /*  第一行索引为 0 */
		{3, 4}, /*  第二行索引为 1 */
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	minimumTotal(a)
}
