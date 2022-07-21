// 88: https://leetcode.cn/problems/merge-sorted-array/
package leetcode

import (
	"fmt"
	"sort"
)

func mergeWithoutMandN(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < len(nums2); i++ {
		for j := 0; j < len(nums1); j++ {
			if nums2[i] <= nums1[j] {
				nums1 = append(nums1[:j], append([]int{nums2[i]}, nums1[j:]...)...)
				nums2 = nums2[i+1:]
				if len(nums2) > 0 {
					mergeWithoutMandN(nums1, len(nums1), nums2, len(nums2))
					break
				}
				break
			}
			if j == len(nums1)-1 {
				nums1 = append(nums1, nums2[i])
				break
			}
		}
	}

	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 双指针
func merge1(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted[m+n:])
}

// 官方解法：双指针
func merge2(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}
