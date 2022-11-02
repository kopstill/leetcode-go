// 709: https://leetcode.cn/problems/to-lower-case/
package leetcode

import (
	"strings"
)

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
