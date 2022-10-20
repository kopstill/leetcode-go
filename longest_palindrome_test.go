package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestPalindromeI(t *testing.T) {
	// s := "xxxbbaceddab"
	// s := "abccccdd"
	// s := "ccc"
	s := "ababababa"
	fmt.Println(longestPalindromeI(s))
}

func Test1LongestPalindromeI(t *testing.T) {
	// s := "ababababa"
	s := "xxx"
	fmt.Println(longestPalindromeI1(s))
}
