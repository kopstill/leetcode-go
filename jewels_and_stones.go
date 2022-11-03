// 771: https://leetcode.cn/problems/jewels-and-stones/
package leetcode

// 暴力
// 时间复杂度：O(mn)，其中 m 是字符串 jewels 的长度，n 是字符串 stones 的长度。
//
//	遍历字符串 stones 的时间复杂度是 O(n)，对于 stones 中的每个字符，
//	需要遍历字符串 jewels 判断是否是宝石，时间复杂度是 O(m)，因此总时间复杂度是 O(mn)。
//
// 空间复杂度：O(1)。只需要维护常量的额外空间。
func numJewelsInStones(jewels string, stones string) int {
	count := 0
	for _, stone := range stones {
		for _, jewel := range jewels {
			if stone == jewel {
				count++
				break
			}
		}
	}
	return count
}

// hash
// 时间复杂度：O(m+n)，其中 m 是字符串 ewels 的长度，n 是字符串 stones 的长度。
//
//	遍历字符串 jewels 将其中的字符存储到哈希集合中，时间复杂度是 O(m)，然后遍历字符串 stones，
//	对于 stones 中的每个字符在 O(1) 的时间内判断当前字符是否是宝石，时间复杂度是 O(n)，因此总时间复杂度是 O(m+n)。
//
// 空间复杂度：O(m)，其中 m 是字符串 jewels 的长度。使用哈希集合存储字符串 jewels 中的字符。
func numJewelsInStones1(jewels string, stones string) int {
	hash := make(map[rune]int, len(jewels))
	for i, jewel := range jewels {
		hash[jewel] = i
	}

	num := 0
	for _, stone := range stones {
		if _, ok := hash[stone]; ok {
			num++
		}
	}

	return num
}

// hash
// 时间复杂度：O(m+n)，其中 m 是字符串 ewels 的长度，n 是字符串 stones 的长度。
//
//	遍历字符串 jewels 将其中的字符存储到哈希集合中，时间复杂度是 O(m)，然后遍历字符串 stones，
//	对于 stones 中的每个字符在 O(1) 的时间内判断当前字符是否是宝石，时间复杂度是 O(n)，因此总时间复杂度是 O(m+n)。
//
// 空间复杂度：O(m)，其中 m 是字符串 jewels 的长度。使用哈希集合存储字符串 jewels 中的字符。
func numJewelsInStones2(jewels string, stones string) int {
	jewelsCount := 0
	jewelsSet := map[byte]bool{}
	for i := range jewels {
		jewelsSet[jewels[i]] = true
	}
	for i := range stones {
		if jewelsSet[stones[i]] {
			jewelsCount++
		}
	}
	return jewelsCount
}
