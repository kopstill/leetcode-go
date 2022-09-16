// 137: https://leetcode.cn/problems/single-number-ii/
package leetcode

// 使用 hash 缓存
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func singleNumberII(nums []int) int {
	cache := make(map[int]int)
	for _, num := range nums {
		cache[num]++
	}
	for k, v := range cache {
		if v == 1 {
			return k
		}
	}
	return -1
}

// 依次确定每一个二进制位（官方）
// 时间复杂度：O(nlogC)，其中 n 是数组的长度，C 是元素的数据范围，在本题中 log C = log 2^32 = 32，也就是我们需要遍历第 0∼31 个二进制位。
// 空间复杂度：O(1)。
func singleNumberII1(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
