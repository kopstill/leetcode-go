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

// 直接合并排序
// 时间复杂度：O((m+n)log(m+n))。排序序列长度为 m+n，套用快速排序的时间复杂度即可，平均情况为 O((m+n)log(m+n))。
// 空间复杂度：O(log(m+n))。排序序列长度为 m+n，套用快速排序的空间复杂度即可，平均情况为 O(log(m+n))。
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 双指针
// 时间复杂度：O(m+n)。指针移动单调递增，最多移动 m+n 次，因此时间复杂度为 O(m+n)。
// 空间复杂度：O(m+n)。需要建立长度为 m+n 的中间数组 sorted。
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
// 时间复杂度：O(m+n)。指针移动单调递增，最多移动 m+n 次，因此时间复杂度为 O(m+n)。
// 空间复杂度：O(m+n)。需要建立长度为 m+n 的中间数组 sorted。
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

// 逆向双指针：官方
// 时间复杂度：O(m+n)。指针移动单调递减，最多移动 m+n 次，因此时间复杂度为 O(m+n)。
// 空间复杂度：O(1)。直接对数组 nums1 原地修改，不需要额外空间。
func merge3(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}
