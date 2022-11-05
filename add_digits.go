// 258: https://leetcode.cn/problems/add-digits/
package leetcode

// 递归
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
