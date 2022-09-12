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
}

func Test1ValidPalindrome(t *testing.T) {
	// s := "A man, a plan, a canal: Panama"
	s := "0P"
	fmt.Println(validPalindrome1(s))
}

func Test2ValidPalindrome(t *testing.T) {
	// s := "A man, a plan, a canal: Panama"
	s := "0P"
	fmt.Println(validPalindrome2(s))
}
