// 258: https://leetcode.cn/problems/add-digits/
package leetcode

// 递归
// 时间复杂度：O(logN)
// 空间复杂度：O(1)
func addDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}

	if sum < 10 {
		return sum
	}

	return addDigits(sum)
}

// 模拟
// 时间复杂度：O(log num)，其中 num 是给定的整数。
//
//	对于 num 计算一次各位相加需要 O(log num) 的时间，由于 num ≤ 2(31) − 1，
//	因此对于 num 计算一次各位相加的最大可能结果是 82，对于任意两位数最多只需要计算两次各位相加的结果即可得到一位数。
//
// 空间复杂度：O(1)。
func addDigits1(num int) int {
	for num >= 10 {
		sum := 0
		for ; num > 0; num /= 10 {
			sum += num % 10
		}
		num = sum
	}
	return num
}
