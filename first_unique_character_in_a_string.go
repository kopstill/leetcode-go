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

// 使用哈希表存储索引
func firstUniqChar3(s string) int {
	n := len(s)
	pos := [26]int{}
	for i := range pos[:] {
		pos[i] = n
	}
	for i, ch := range s {
		ch -= 'a'
		if pos[ch] == n {
			pos[ch] = i
		} else {
			pos[ch] = n + 1
		}
	}
	ans := n
	for _, p := range pos[:] {
		if p < ans {
			ans = p
		}
	}
	if ans < n {
		return ans
	}
	return -1
}

type pair struct {
	ch  byte
	pos int
}

// 队列
func firstUniqChar4(s string) int {
	n := len(s)
	pos := [26]int{}
	for i := range pos[:] {
		pos[i] = n
	}
	q := []pair{}
	for i := range s {
		ch := s[i] - 'a'
		if pos[ch] == n {
			pos[ch] = i
			q = append(q, pair{ch, i})
		} else {
			pos[ch] = n + 1
			for len(q) > 0 && pos[q[0].ch] == n+1 {
				q = q[1:]
			}
		}
	}
	if len(q) > 0 {
		return q[0].pos
	}
	return -1
}
