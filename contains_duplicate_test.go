package leetcode

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
}

func Test1ContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 7}
	fmt.Println(containsDuplicate1(nums))
}

func Test2ContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 2, 2, 7, 0, 7}
	fmt.Println(containsDuplicate2(nums))
}
