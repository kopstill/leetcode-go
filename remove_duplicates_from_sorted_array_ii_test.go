package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicatesFromSortedArrayII(t *testing.T) {
	// nums := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	// nums := []int{1, 1, 1, 2, 2, 3}
	nums := []int{1, 1, 1, 1}
	fmt.Println(removeDuplicatesII(nums))
	fmt.Println(nums)
}

func Test1RemoveDuplicatesFromSortedArrayII(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	fmt.Println(removeDuplicatesII1(nums))
	fmt.Println(nums)
}
