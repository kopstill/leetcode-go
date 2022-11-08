// 1684: https://leetcode.cn/problems/count-the-number-of-consistent-strings/
package leetcode

import "github.com/emirpasic/gods/sets/hashset"

// set
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
