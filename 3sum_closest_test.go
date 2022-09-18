package leetcode

import (
	"fmt"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	// nums := []int{-1, 2, 1, -4}
	// nums := []int{0, 0, 0}
	nums := []int{12, 3, -5, 7, 1, 9, -10, 32}
	target := -15
	fmt.Println(threeSumClosest(nums, target))
}
