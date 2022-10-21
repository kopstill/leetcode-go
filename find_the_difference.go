// 389: https://leetcode.cn/problems/find-the-difference/
package leetcode

// 循环删除字符串 t 中的重合字符，剩余字符即为不重合字符
func findTheDifference(s, t string) (result byte) {
	tLength := len(t)
	tarr := make([]byte, tLength)
	for i := 0; i < tLength; i++ {
		tarr[i] = t[i]
	}

	for i := 0; i < len(s); i++ {
		if ok, idx := isContain(s[i], tarr); ok {
			tarr = append(tarr[:idx], tarr[idx+1:]...)
		}
	}

	return tarr[0]
}

func isContain(target byte, arr []byte) (bool, int) {
	for i, v := range arr {
		if v == target {
			return true, i
		}
	}
	return false, -1
}

// 计数
// 首先遍历字符串 s，对其中的每个字符都将计数值加 1；
// 然后遍历字符串 t，对其中的每个字符都将计数值减 1。
// 当发现某个字符计数值为负数时，说明该字符在字符串 t 中出现的次数大于在字符串 s 中出现的次数，因此该字符为被添加的字符。
// 时间复杂度：O(N)，其中 N 为字符串的长度。
// 空间复杂度：O(∣Σ∣)，其中 Σ 是字符集，这道题中字符串只包含小写字母，∣Σ∣=26。需要使用数组对每个字符计数。
func findTheDifference1(s, t string) byte {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

	for i := 0; ; i++ {
		ch := t[i]
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return ch
		}
	}
}
