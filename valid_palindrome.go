// 125: https://leetcode.cn/problems/valid-palindrome/
package leetcode

import (
	"strings"
)

// 双指针
func validPalindrome(s string) bool {
	news := ""
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			news += strings.ToLower(string(c))
		}
	}

	left, right := 0, len(news)-1
	for left < right {
		if news[left] == news[right] {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}
