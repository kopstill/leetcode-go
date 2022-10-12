// 169: https://leetcode.cn/problems/majority-element/
package leetcode

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func majorityElement(nums []int) int {
	len := len(nums)
	tmp := make(map[int]int)

	for i := 0; i < len; i++ {
		if v, ok := tmp[nums[i]]; ok {
			tmp[nums[i]] = v + 1
		} else {
			tmp[nums[i]] = 1
		}
	}

	for k, v := range tmp {
		if v > len/2 {
			return k
		}
	}

	return -1
}
