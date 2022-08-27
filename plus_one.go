// 66: https://leetcode.cn/problems/plus-one/
package leetcode

// 逆向遍历
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

// 逆向找到第一个不为 9 的数
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func plusOne1(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}

	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

// 直接将末尾的 9 改成 0
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func plusOne2(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	return append([]int{1}, digits...)
}

// 判断 +1 之后是否为 0
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func plusOne3(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i] = (digits[i] + 1) % 10
		if digits[i] != 0 {
			return digits
		}
	}
	return append([]int{1}, digits...)
}
