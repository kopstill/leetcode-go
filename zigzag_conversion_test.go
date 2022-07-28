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

func Test1ZigZagConversion(t *testing.T) {
	s := "THISISANFIELD"
	numRows := 3
	fmt.Println(convert1(s, numRows))
}

func Test2ZigZagConversion(t *testing.T) {
	s := "THISMEANSMORE"
	numRows := 3
	fmt.Println(convert2(s, numRows))
}

func Test3ZigZagConversion(t *testing.T) {
	s := "FORMANISMANANDMASTEROFHISFATE"
	numRows := 7
	fmt.Println(convert3(s, numRows))
}

func Test4ZigZagConversion(t *testing.T) {
	s := "SOMEBODYTHATIUSEDTOKNOW"
	numRows := 4
	fmt.Println(convert4(s, numRows))
}
