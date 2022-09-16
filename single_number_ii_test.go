package leetcode

import (
	"fmt"
	"testing"
)

func TestSingleNumberII(t *testing.T) {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumberII(nums))
}

func Test1SingleNumberII(t *testing.T) {
	nums := []int{0, 1, 0, 1, 0, 1, 99, 0, 1}
	fmt.Println(singleNumberII1(nums))
}
