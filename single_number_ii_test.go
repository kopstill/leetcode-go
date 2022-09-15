package leetcode

import (
	"fmt"
	"testing"
)

func TestSingleNumberII(t *testing.T) {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumberII(nums))
}
