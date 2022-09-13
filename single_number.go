// 136: https://leetcode.cn/problems/single-number/
package leetcode

// 使用 hash 缓存
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func singleNumber(nums []int) int {
	cache := make(map[int]int)
	for _, num := range nums {
		cache[num] = cache[num] + 1
	}
	for k, v := range cache {
		if v == 1 {
			return k
		}
	}
	return -1
}

// 位运算：异或
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func singleNumber1(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
