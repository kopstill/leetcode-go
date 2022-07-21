package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 5, 1, 6, 9, 7}
	target := 15
	result := twoSum(nums, target)
	fmt.Println(result)
}

func Test1TwoSum(t *testing.T) {
	nums := []int{362, 509, 187, 604, 985, 736}
	target := 1494
	result := twoSum1(nums, target)
	fmt.Println(result)
}

func Test2TwoSum(t *testing.T) {
	nums := []int{32, 68, 51, 93, 25, 80, 19}
	target := 99
	result := twoSum2(nums, target)
	fmt.Println(result)
}
