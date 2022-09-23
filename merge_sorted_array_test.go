package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeSortedArrayWithoutMandN(t *testing.T) {
	// nums1 := []int{1, 3}
	// nums2 := []int{2}

	// nums1 := []int{100, 200, 300}
	// nums2 := []int{50, 250, 1000}

	nums1 := []int{3, 6, 10, 23, 34, 99, 128}
	nums2 := []int{56, 71, 103, 124, 255}

	mergeWithoutMandN(nums1, len(nums1), nums2, len(nums2))
}

func Test0MergeSortedArray(t *testing.T) {
	// nums1 := []int{1, 3}
	// nums2 := []int{2}

	// nums1 := []int{100, 200, 300}
	// nums2 := []int{50, 250, 1000}

	// nums1 := []int{3, 6, 10, 23, 34, 99, 128}
	// nums2 := []int{56, 71, 103, 124, 255}

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func Test1MergeSortedArray(t *testing.T) {
	// nums1 := []int{1, 3}
	// nums2 := []int{2}

	// nums1 := []int{100, 200, 300}
	// nums2 := []int{50, 250, 1000}

	// nums1 := []int{3, 6, 10, 23, 34, 99, 128}
	// nums2 := []int{56, 71, 103, 124, 255}

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge1(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func Test2MergeSortedArray(t *testing.T) {
	// nums1 := []int{1, 3}
	// nums2 := []int{2}

	// nums1 := []int{100, 200, 300}
	// nums2 := []int{50, 250, 1000}

	// nums1 := []int{3, 6, 10, 23, 34, 99, 128}
	// nums2 := []int{56, 71, 103, 124, 255}

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge2(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func Test3MergeSortedArray(t *testing.T) {
	// nums1 := []int{3, 6, 10, 23, 34, 99, 128, 0, 0, 0, 0, 0}
	// nums2 := []int{56, 71, 103, 124, 255}
	// merge3(nums1, 7, nums2, 5)

	nums1 := []int{0}
	nums2 := []int{1}
	merge3(nums1, 0, nums2, 1)

	// nums1 := []int{2, 0}
	// nums2 := []int{1}
	// merge3(nums1, 1, nums2, 1)

	fmt.Println(nums1)
}
