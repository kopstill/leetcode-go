package leetcode

import (
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a, b))
}

func Test1AddBinary(t *testing.T) {
	a := "11111"
	b := "11111"
	fmt.Println(addBinary1(a, b))
}
