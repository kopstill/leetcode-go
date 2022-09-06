package leetcode

import (
	"fmt"
	"testing"
)

func TestSearchInsertPosition(t *testing.T) {
	nums := []int{1, 2, 3, 6, 8, 19, 32, 54, 79, 101, 132, 256}
	target := 100
	fmt.Println(searchInsert(nums, target))
}
