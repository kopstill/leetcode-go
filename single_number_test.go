package leetcode

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	// nums := []int{4, 1, 2, 1, 2}
	// nums := []int{2, 2, 1}
	nums := []int{1, 2, 1}
	fmt.Println(singleNumber(nums))
}
