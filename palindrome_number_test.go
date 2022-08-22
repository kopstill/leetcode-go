package leetcode

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	// x := 101
	// x := 12321
	// x := 8
	x := 12345054321
	fmt.Println(isPalindrome(x))
}

func Test1IsPalindrome(t *testing.T) {
	x := 6
	fmt.Println(isPalindrome1(x))
}
