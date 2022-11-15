// 383: https://leetcode.cn/problems/ransom-note/
package leetcode

import (
	"github.com/emirpasic/gods/lists/arraylist"
)

// hash
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func canConstruct(ransomNote string, magazine string) bool {
	hash1 := make(map[rune]int)
	for _, r := range ransomNote {
		hash1[r]++
	}
	hash2 := make(map[rune]int)
	for _, r := range magazine {
		hash2[r]++
	}
	for k, v := range hash1 {
		if hash2[k] < v {
			return false
		}
	}
	return true
}

// list
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func canConstruct1(ransomNote string, magazine string) bool {
	list := arraylist.New()
	for _, r := range ransomNote {
		list.Add(r)
	}
	count := 0
	it := list.Iterator()
	for it.Next() {
		value := it.Value()
		for i, v := range magazine {
			if v == value {
				count++
				magazine = magazine[0:i] + magazine[i+1:]
				break
			}
		}
	}
	return count == len(ransomNote)
}

// 字符统计
// 时间复杂度：O(m+n)，其中 m 是字符串 ransomNote 的长度，n 是字符串 magazine 的长度，我们只需要遍历两个字符一次即可。
// 空间复杂度：O(∣S∣)，S 是字符集，这道题中 S 为全部小写英语字母，因此 ∣S∣=26。
func canConstruct2(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	cnt := [26]int{}
	for _, ch := range magazine {
		cnt[ch-'a']++
	}
	for _, ch := range ransomNote {
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return false
		}
	}
	return true
}
