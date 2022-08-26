// 20: https://leetcode.cn/problems/valid-parentheses/
package leetcode

// 迭代删除最小括号对
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isValid(s string) bool {
	p := 0
	for p < len(s)-1 {
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

// 栈
// 时间复杂度：O(n)，其中 n 是字符串 s 的长度。
// 空间复杂度：O(n+∣Σ∣)，其中 Σ 表示字符集，本题中字符串只包含 6 种括号，∣Σ∣=6。栈中的字符数量为 O(n)，而哈希表使用的空间为 O(∣Σ∣)，相加即可得到总空间复杂度。
func isValid1(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}
	for i := 0; i < n; i++ {
		c := s[i]
		if v, ok := pairs[c]; ok {
			if len(stack) == 0 || v != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}
