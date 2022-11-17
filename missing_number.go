// 268: https://leetcode.cn/problems/missing-number/
package leetcode

import (
	"sort"

	"github.com/emirpasic/gods/sets/hashset"
)

// set
// 时间复杂度：O(n)，其中 n 是数组 nums 的长度。遍历数组 nums 将元素加入 set 的时间复杂度是 O(n)，遍历从 0 到 n 的每个整数并判断是否在 set 中的时间复杂度也是 O(n)。
// 空间复杂度：O(n)，其中 n 是数组 nums 的长度。set 中需要存储 n 个整数。
func missingNumber(nums []int) int {
	set := hashset.New()
	for _, v := range nums {
		set.Add(v)
	}
	for i := 0; i <= len(nums); i++ {
		if !set.Contains(i) {
			return i
		}
	}
	return -1
}

// 排序
// 时间复杂度：O(nlogn)，其中 n 是数组 nums 的长度。排序的时间复杂度是 O(nlogn)，遍历数组寻找丢失的数字的时间复杂度是 O(n)，因此总时间复杂度是 O(nlogn)。
// 空间复杂度：O(logn)，其中 n 是数组 nums 的长度。空间复杂度主要取决于排序的递归调用栈空间。
func missingNumber1(nums []int) int {
	sort.Ints(nums)
	for i, num := range nums {
		if num != i {
			return i
		}
	}
	return len(nums)
}

// hash
// 时间复杂度：O(n)，其中 n 是数组 nums 的长度。遍历数组 nums 将元素加入哈希集合的时间复杂度是 O(n)，遍历从 0 到 n 的每个整数并判断是否在哈希集合中的时间复杂度也是 O(n)。
// 空间复杂度：O(n)，其中 n 是数组 nums 的长度。哈希集合中需要存储 n 个整数。
func missingNumber2(nums []int) int {
	has := map[int]bool{}
	for _, v := range nums {
		has[v] = true
	}
	for i := 0; ; i++ {
		if !has[i] {
			return i
		}
	}
}

// 位运算
// 时间复杂度：O(n)，其中 n 是数组 nums 的长度。需要对 2n+1 个数字计算按位异或的结果。
// 空间复杂度：O(1)。
func missingNumber3(nums []int) (xor int) {
	for i, num := range nums {
		xor ^= i ^ num
	}
	return xor ^ len(nums)
}

// 数学
// 时间复杂度：O(n)，其中 n 是数组 nums 的长度。需要 O(1) 的时间计算从 0 到 n 的全部整数之和以及需要 O(n) 的时间计算数组 nums 的元素之和。
// 空间复杂度：O(1)。
func missingNumber4(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	arrSum := 0
	for _, num := range nums {
		arrSum += num
	}
	return total - arrSum
}
