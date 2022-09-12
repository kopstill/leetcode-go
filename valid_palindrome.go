// 125: https://leetcode.cn/problems/valid-palindrome/
package leetcode

import "strings"

// 双指针
// 时间复杂度：O(|s|)，其中 ∣s∣ 是字符串 s 的长度。
// 空间复杂度：O(|s|)。由于我们需要将所有的字母和数字字符存放在另一个字符串中，在最坏情况下，新的字符串 sgood 与原字符串 s 完全相同，因此需要使用 O(|s|) 的空间。
func validPalindrome(s string) bool {
	var str string
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			str += string(c)
		}
	}

	str = strings.ToLower(str)
	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}
