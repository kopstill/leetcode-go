// 387: https://leetcode.cn/problems/first-unique-character-in-a-string/
package leetcode

// 暴力轮询
func firstUniqChar(s string) int {
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

// 数组计数
func firstUniqChar1(s string) int {
	count := [26]byte{}
	for _, ch := range s {
		count[ch-'a']++
	}

	for i, ch := range s {
		if count[ch-'a'] == 1 {
			return i
		}
	}

	return -1
}

// hash
func firstUniqChar2(s string) int {
	count := make(map[rune]int)

	for _, ch := range s {
		count[ch-'a']++
	}

	for i, ch := range s {
		if count[ch-'a'] == 1 {
			return i
		}
	}

	return -1
}
