// 28: https://leetcode.cn/problems/implement-strstr/
package leetcode

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	c := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[c] {
			c++
			if c == len(needle) {
				return i - c + 1
			}
		} else {
			i = i - c
			c = 0
		}
	}
	return -1
}

// 暴力匹配
// 时间复杂度：O(n*m)，其中 n 是字符串 haystack 的长度，m 是字符串 needle 的长度。最坏情况下我们需要将字符串 needle 与字符串 haystack 的所有长度为 m 的子串均匹配一次。
// 空间复杂度：O(1)。我们只需要常数的空间保存若干变量。
func strStr1(haystack, needle string) int {
	n, m := len(haystack), len(needle)
outer:
	for i := 0; i+m <= n; i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}
