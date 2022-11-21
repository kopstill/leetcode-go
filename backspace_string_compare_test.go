package leetcode

import (
	"fmt"
	"testing"
)

func TestBackspaceCompare(t *testing.T) {
	// ss := "c#b#"
	// tt := "c#b#"
	// ss := "a##c"
	// tt := "#a#c"
	ss := "y#fo##f"
	tt := "y#f#o##f"
	fmt.Println(backspaceCompare(ss, tt))
}
