// 414: https://leetcode.cn/problems/third-maximum-number/
package leetcode

import (
	"sort"
)

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func thirdMax(nums []int) int {
	sort.Ints(nums)

	len := len(nums)
	idx := 1
	for i := len - 1; i > 0; i-- {
		if nums[i] != nums[i-1] {
			idx++
		}
		if idx == 3 {
			return nums[i-1]
		}
	}

	return nums[len-1]
}
