// 844: https://leetcode.cn/problems/backspace-string-compare/
package leetcode

import "strings"

// 原地修改字符串
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func backspaceCompare(s string, t string) bool {
	i := 0
	for i < len(s)-1 {
		if s[i+1] == '#' {
			if i+2 < len(s) {
				s = s[:i] + s[i+2:]
			} else {
				s = s[:i]
			}
			if i != 0 {
				i--
			}
		} else {
			i++
		}
	}

	j := 0
	for j < len(t)-1 {
		if t[j+1] == '#' {
			if j+2 < len(t) {
				t = t[:j] + t[j+2:]
			} else {
				t = t[:j]
			}
			if j != 0 {
				j--
			}
		} else {
			j++
		}
	}
	return strings.ReplaceAll(s, "#", "") == strings.ReplaceAll(t, "#", "")
}

// 重构字符串
// 时间复杂度：O(N+M)，其中 N 和 M 分别为字符串 S 和 T 的长度。我们需要遍历两字符串各一次。
// 空间复杂度：O(N+M)，其中 N 和 M 分别为字符串 S 和 T 的长度。主要为还原出的字符串的开销。
func backspaceCompare1(s string, t string) bool {
	return build(s) == build(t)
}

func build(str string) string {
	s := []byte{}
	for i := range str {
		if str[i] == '#' {
			s = append(s, str[i])
		} else if len(s) > 0 {
			s = s[:len(s)-1]
		}
	}
	return string(s)
}
