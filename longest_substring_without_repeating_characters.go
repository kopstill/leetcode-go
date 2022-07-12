package leetcode

// 支持中文等非等字节字符
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 || length == 1 {
		return length
	}

	max := 1

	r := []rune(s)

	t := []string{}

	for i := 0; i < len(r); i++ {
		t = append(t, string(r[i]))
		for j := i + 1; j < len(r); j++ {
			if isRepeat(t, string(r[j])) {
				t = nil
				break
			} else {
				t = append(t, string(r[j]))
			}
			if len(t) > max {
				max = len(t)
			}
		}
	}

	return max
}

func isRepeat(slc []string, r string) (flag bool) {
	for _, sv := range slc {
		if sv == r {
			return true
		}
	}

	return
}

// 仅支持单字节字符
func lengthOfLongestSubstring1(s string) int {
	length := len(s)
	if length == 0 || length == 1 {
		return length
	}

	max := 1

	t := []byte{}

	for i := 0; i < len(s); i++ {
		t = append(t, s[i])
		for j := i + 1; j < len(s); j++ {
			if isRepeat1(t, s[j]) {
				t = nil
				break
			} else {
				t = append(t, s[j])
			}
			if len(t) > max {
				max = len(t)
			}
		}
	}

	return max
}

func isRepeat1(slc []byte, r byte) (flag bool) {
	for _, sv := range slc {
		if sv == r {
			return true
		}
	}

	return
}

// 官方解法（仅支持单字节字符）
func lengthOfLongestSubstring2(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
