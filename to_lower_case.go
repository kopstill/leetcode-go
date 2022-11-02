// 709: https://leetcode.cn/problems/to-lower-case/
package leetcode

import (
	"strings"
)

// 通过大小写之间的距离实现转换
// 时间复杂度：O(n)，其中 n 是字符串 s 的长度。
// 空间复杂度：O(1)，不考虑返回值的空间占用。
func toLowerCase(s string) string {
	sb := strings.Builder{}
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			sb.WriteRune(c + 32)
		} else {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

// 使用语言内置 API
// 时间复杂度：O(n)，其中 n 是字符串 s 的长度。
// 空间复杂度：O(1)，不考虑返回值的空间占用。
func toLowerCase1(s string) string {
	return strings.ToLower(s)
}

// 时间复杂度：O(n)，其中 n 是字符串 s 的长度。
// 空间复杂度：O(1)，不考虑返回值的空间占用。
func toLowerCase2(s string) string {
	lower := &strings.Builder{}
	lower.Grow(len(s))
	for _, ch := range s {
		if 65 <= ch && ch <= 90 {
			ch |= 32
		}
		lower.WriteRune(ch)
	}
	return lower.String()
}
