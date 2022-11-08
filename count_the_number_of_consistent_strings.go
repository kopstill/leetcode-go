// 1684: https://leetcode.cn/problems/count-the-number-of-consistent-strings/
package leetcode

import "github.com/emirpasic/gods/sets/hashset"

// set + 遍历
func countConsistentStrings(allowed string, words []string) int {
	count := 0
	set := hashset.New()
	for _, c := range allowed {
		set.Add(c)
	}
	for _, w := range words {
		count++
		for _, c := range w {
			if !set.Contains(c) {
				count--
				break
			}
		}
	}
	return count
}

// 位运算 + 遍历
// 时间复杂度：O(n+∑m_i)，其中 n 是字符串 allowed 的长度，m_i 是字符串 words[i] 的长度。
// 空间复杂度：O(1)。
func countConsistentStrings1(allowed string, words []string) (res int) {
	mask := 0
	for _, c := range allowed {
		mask |= 1 << (c - 'a')
	}
	for _, word := range words {
		mask1 := 0
		for _, c := range word {
			mask1 |= 1 << (c - 'a')
		}
		if mask1|mask == mask {
			res++
		}
	}
	return
}
