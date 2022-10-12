package leetcode

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}

func Test1MajorityElement(t *testing.T) {
	// nums := []int{2, 2, 1, 1, 1, 2, 2}
	nums := []int{1}
	fmt.Println(majorityElement1(nums))
}
