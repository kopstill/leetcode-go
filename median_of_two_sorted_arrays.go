// 4: https://leetcode.cn/problems/median-of-two-sorted-arrays/
package leetcode

import (
	"sort"
)

// 使用标准 SDK，时间复杂度参考快速排序：O(nlogn)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := sort.IntSlice(append(nums1, nums2...))
	nums.Sort()

	m := len(nums) / 2
	if len(nums)%2 == 0 {
		return float64(nums[m-1]+nums[m]) / 2
	}

	return float64(nums[m])
}

// 官方解法：二分查找，时间复杂度：O(log(m+n))，空间复杂度 O(1)
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
