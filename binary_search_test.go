package leetcode

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	// nums := []int{-1, 0, 3, 5, 9, 12}
	// target := 9
	nums := []int{2, 5}
	target := 2
	fmt.Println(search(nums, target))
}
