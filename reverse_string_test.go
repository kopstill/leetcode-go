package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := []byte{'H', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(string(s))
}

func Test1ReverseString(t *testing.T) {
	s := []byte{'L', 'e', 'e', 't', 'c', 'o', 'd', 'e'}
	reverseString1(s)
	fmt.Println(string(s))
}
