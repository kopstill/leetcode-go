package leetcode

import (
	"fmt"
	"testing"
)

func TestZigZagConversion(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3
	// numRows := 4

	// s := "A"
	// numRows := 1
	fmt.Println(convert(s, numRows))
}
