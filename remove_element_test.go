package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{1, 23, 42, 6, 23, 56, 32, 65, 23, 89}
	val := 23
	// nums := []int{3, 2, 2, 3}
	// val := 3
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
}

func Test1RemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3, 1, 4, 8, 3, 2, 5, 7, 6}
	val := 3
	fmt.Println(removeElement1(nums, val))
	fmt.Println(nums)
}

func Test2RemoveElement(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 1, 1, 9, 8, 7, 1}
	val := 1
	fmt.Println(removeElement2(nums, val))
	fmt.Println(nums)
}
