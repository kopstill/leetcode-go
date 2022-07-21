package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	// s := "kxacorrocbzx"
	// s := "cbbd"
	// s := "x"
	// s := "ac"
	s := "bacabab"
	r := longestPalindrome(s)
	fmt.Println(r)
}
