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

func Test1SingleNumber(t *testing.T) {
	// nums := []int{4, 1, 2, 1, 2}
	// nums := []int{2, 2, 1}
	// nums := []int{1, 1, 2}
	// nums := []int{1, 2, 1}
	nums := []int{12, 25, 17, 12, 17, 25, 19, 100, 100}
	fmt.Println(singleNumber1(nums))
}
