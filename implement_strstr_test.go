package leetcode

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	// haystack := "hello"
	// needle := "ll"
	haystack := "mississippi"
	needle := "issip"
	fmt.Println(strStr(haystack, needle))
}

func Test1StrStr(t *testing.T) {
	// haystack := "hello"
	// needle := "ll"
	haystack := "mississippi"
	needle := "issip"
	fmt.Println(strStr1(haystack, needle))
}
