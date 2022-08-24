package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	// x := 123
	// x := 0
	x := -12365472345234523
	fmt.Println(reverse(x))
}

func Test1ReverseInteger(t *testing.T) {
	// x := 123
	// x := 0
	x := -123456789
	fmt.Println(reverse1(x))
}
