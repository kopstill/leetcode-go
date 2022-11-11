package leetcode

import (
	"fmt"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	// nums := []int{1, 1}
	fmt.Println(findDisappearedNumbers(nums))
}

func Test1FindDisappearedNumbers(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	// nums := []int{1, 1}
	fmt.Println(findDisappearedNumbers1(nums))
}

func Test2FindDisappearedNumbers(t *testing.T) {
	// nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	nums := []int{1, 1}
	fmt.Println(findDisappearedNumbers2(nums))
}

func Test3FindDisappearedNumbers(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	// nums := []int{8, 8, 8, 8, 8, 8, 8, 8}
	// nums := []int{1, 1}
	fmt.Println(findDisappearedNumbers3(nums))
}
