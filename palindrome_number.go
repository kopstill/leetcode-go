// 9: https://leetcode.cn/problems/palindrome-number/
package leetcode

// 时间复杂度：O(log(n))，每次迭代将目标数除以 10
// 空间复杂度：O(1)，无需额外空间存储额外变量
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	new := 0
	for {
		x = x / 10
		new = new*10 + x%10
		if new == x {
			return true
		}
		if new > x {
			if new/10 == x {
				return true
			}
			break
		}
	}

	return false
}

// 官方解法
// 时间复杂度：O(log(n))
// 空间复杂度：O(1)
func isPalindrome1(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for revertedNumber < x {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	return revertedNumber == x || revertedNumber/10 == x
}
