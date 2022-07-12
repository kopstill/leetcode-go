package leetcode

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := sort.IntSlice(append(nums1, nums2...))
	nums.Sort()

	m := len(nums) / 2
	if len(nums)%2 == 0 {
		return float64(nums[m-1]+nums[m]) / 2
	}

	return float64(nums[m])
}
