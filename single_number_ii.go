// 137: https://leetcode.cn/problems/single-number-ii/
package leetcode

// 使用 hash 缓存
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func singleNumberII(nums []int) int {
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
