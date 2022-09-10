package leetcode

import (
	"fmt"
	"testing"
)

func Test_RemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
}

func Test_1_RemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates1(nums))
	fmt.Println(nums)
}
