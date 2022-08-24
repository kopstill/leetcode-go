// 14: https://leetcode.cn/problems/longest-common-prefix/
package leetcode

// 时间复杂度：O(m*n)
// 空间复杂度：O(1)
func longestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) < 1 {
		return prefix
	}

	for i, s := range strs[0] {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || rune(strs[j][i]) != s {
				return prefix
			}
		}
		prefix += string(s)
	}

	return prefix
}
