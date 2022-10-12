// 169: https://leetcode.cn/problems/majority-element/
package leetcode

import "sort"

// 使用 hash
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

// 先排序，然后计算个数
// 时间复杂度：O(nlogn)，将数组排序的时间复杂度为 O(nlogn)。
// 空间复杂度：O(logn)，使用语言自带的排序算法，需要使用 O(logn) 的栈空间。
func majorityElement1(nums []int) (val int) {
	sort.Ints(nums)

	len := len(nums)
	count := 1
	val = nums[0]
	for i := 0; i < len-1; i++ {
		if nums[i] == nums[i+1] {
			count++
		} else {
			count = 1
		}
		if count > len/2 {
			val = nums[i]
			break
		}
	}
	return
}

// 先排序，nums[n/2] 的元素一定是众数
// 时间复杂度：O(nlogn)，将数组排序的时间复杂度为 O(nlogn)。
// 空间复杂度：O(logn)，使用语言自带的排序算法，需要使用 O(logn) 的栈空间。
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
