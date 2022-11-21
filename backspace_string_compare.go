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
