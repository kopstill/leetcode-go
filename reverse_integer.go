// 7: https://leetcode.cn/problems/reverse-integer/
package leetcode

import "math"

// 时间复杂度：O(log(n))
// 空间复杂度：O(1)
func reverse(x int) int {
	reversedNumber := 0
	for x > 0 || x < 0 {
		reversedNumber = reversedNumber*10 + x%10
		x /= 10
	}
	if reversedNumber > math.MaxInt32 || reversedNumber < math.MinInt32 {
		return 0
	}
	return reversedNumber
}

// 官方解法
// 时间复杂度：O(log|x|)。翻转的次数即 x 十进制的位数。
// 空间复杂度：O(1)
func reverse1(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}
