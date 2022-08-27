// 66: https://leetcode.cn/problems/plus-one/
package leetcode

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func plusOne(digits []int) []int {
	n, carry := len(digits), 0
	result := []int{}

	for i := n - 1; i >= 0; i-- {
		d := digits[i] + carry
		if i == n-1 {
			d++
		}

		if d == 10 {
			d = 0
			carry = 1
		} else {
			carry = 0
		}

		result = append([]int{d}, result...)
	}

	if carry > 0 {
		result = append([]int{carry}, result...)
	}

	return result
}
