package leetcode

import (
	"fmt"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	// s := "  abcba   "
	// s := ""
	// s := "   "
	// s := " ab cc ba  "
	// s := " a B c d D c B A      "
	// s := " aBcbA "
	// s := "Aa"
	// s := "A man, a plan, a canal: Panama"
	s := "0P"
	fmt.Println(validPalindrome(s))

	fmt.Println('0')
	fmt.Println('9')
}
