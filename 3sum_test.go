package leetcode

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{0, 1, 1}
	// nums := []int{0, 0, 0}
	fmt.Println(threeSum(nums))
}
