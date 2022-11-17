package leetcode

import (
	"fmt"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	// nums := []int{3, 0, 1}
	nums := []int{0}
	fmt.Println(missingNumber(nums))
}

func Test1MissingNumber(t *testing.T) {
	// nums := []int{3, 0, 1}
	// nums := []int{0}
	nums := []int{2, 0, 1}
	fmt.Println(missingNumber1(nums))
}

func Test2MissingNumber(t *testing.T) {
	nums := []int{3, 0, 1}
	// nums := []int{0}
	// nums := []int{2, 0, 1}
	fmt.Println(missingNumber2(nums))
}

func Test3MissingNumber(t *testing.T) {
	nums := []int{3, 0, 1}
	// nums := []int{0}
	// nums := []int{2, 0, 1}
	fmt.Println(missingNumber3(nums))
}

func Test4MissingNumber(t *testing.T) {
	// nums := []int{3, 0, 1}
	// nums := []int{0}
	nums := []int{2, 0, 1}
	fmt.Println(missingNumber4(nums))
}
