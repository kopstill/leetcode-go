package leetcode

import (
	"fmt"
	"testing"
)

func TestThirdMax(test *testing.T) {
	// nums := []int{1, 3, 4, 5, 7, 5, 3, 4, 1, 4}
	nums := []int{1, 2}
	fmt.Println(thirdMax(nums))
}

func Test1ThirdMax(test *testing.T) {
	nums := []int{2, 2, 2, 3, 3, 1, 1, 1, 0}
	fmt.Println(thirdMax1(nums))
}

func Test2ThirdMax(test *testing.T) {
	nums := []int{3, 5, 1}
	fmt.Println(thirdMax2(nums))
}
