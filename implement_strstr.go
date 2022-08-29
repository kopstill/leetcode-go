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
