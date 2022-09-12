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

// 双指针（官方写法）
// 时间复杂度：O(|s|)，其中 ∣s∣ 是字符串 s 的长度。
// 空间复杂度：O(|s|)。由于我们需要将所有的字母和数字字符存放在另一个字符串中，在最坏情况下，新的字符串 sgood 与原字符串 s 完全相同，因此需要使用 O(|s|) 的空间。
func validPalindrome1(s string) bool {
	var sgood string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			sgood += string(s[i])
		}
	}

	n := len(sgood)
	sgood = strings.ToLower(sgood)
	for i := 0; i < n/2; i++ {
		if sgood[i] != sgood[n-1-i] {
			return false
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

// 在原字符串上直接判断（官方解法）
// 时间复杂度：O(|s|)，其中 |s∣ 是字符串 s 的长度。
// 空间复杂度：O(1)
func validPalindrome2(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if left < right { // 此判断可要可不要
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}
