package leetcode

import "fmt"

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
