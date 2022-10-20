// 409: https://leetcode.cn/problems/longest-palindrome/
package leetcode

// map
// 时间复杂度：O(n)。n 为字符串 s 的长度。
// 空间复杂度：O(s)。s 为 hash 存储的大小。
func longestPalindromeI(s string) int {
	size := len(s)
	if size == 1 {
		return 1
	}

	counts := make(map[string]int)
	for i := 0; i < size; i++ {
		counts[string(s[i])] += 1
	}

	if len(counts) == 1 {
		return size
	}

	result := 0
	single := false
	for _, v := range counts {
		if v%2 == 0 {
			result += v
		} else if v%2 == 1 {
			result += v - 1
			single = true
		}
	}

	if single {
		result += 1
	}

	return result
}
