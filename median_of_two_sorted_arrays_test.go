package leetcode

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	// nums1 := []int{1, 2}
	// nums2 := []int{3, 4}

	// nums1 := []int{1, 3, 5, 10}
	// nums2 := []int{2, 4, 13}

	// nums1 := []int{1}
	// nums2 := []int{}

	nums1 := []int{1, 3}
	nums2 := []int{2}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func Test1FindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3, 4, 9}
	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(findMedianSortedArrays1(nums1, nums2))
}
