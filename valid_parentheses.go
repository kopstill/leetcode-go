// 20: https://leetcode.cn/problems/valid-parentheses/
package leetcode

// 删除最小括号对
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isValid(s string) bool {
	p := 0
	for p < len(s)-1 && len(s) > 0 {
		if s[p] == '(' && s[p+1] == ')' ||
			s[p] == '[' && s[p+1] == ']' ||
			s[p] == '{' && s[p+1] == '}' {
			s = s[0:p] + s[p+2:]
			p = 0
		} else {
			p++
		}
	}

	return len(s) == 0
}
