// 387: https://leetcode.cn/problems/first-unique-character-in-a-string/
package leetcode

func firstUniqChar(s string) int { // leetcodel
	for i := 0; i < len(s); i++ {
		isUnique := true
		for j := 0; j < len(s); j++ {
			if i != j {
				if s[i] == s[j] {
					isUnique = false
					break
				}
			}
		}
		if isUnique {
			return i
		}
	}

	return -1
}
