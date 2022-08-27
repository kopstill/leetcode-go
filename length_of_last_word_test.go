package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	// s := "  nice to meet you   "
	// s := "a "
	s := "a"
	fmt.Println(lengthOfLastWord(s))
}

func Test1LengthOfLastWord(t *testing.T) {
	s := "  nice to meet you   "
	// s := "a "
	// s := "a"
	fmt.Println(lengthOfLastWord1(s))
}
