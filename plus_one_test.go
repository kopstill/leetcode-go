package leetcode

import (
	"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {
	// digits := []int{1, 2, 9}
	digits := []int{9, 9, 9}
	fmt.Println(plusOne(digits))
}

func Test1PlusOne(t *testing.T) {
	// digits := []int{1, 2, 9}
	// digits := []int{9, 9, 9}
	digits := []int{1, 0, 0, 9, 9, 9}
	fmt.Println(plusOne1(digits))
}
