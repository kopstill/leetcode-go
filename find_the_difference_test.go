package leetcode

import (
	"fmt"
	"testing"
)

func TestFindTheDifference(test *testing.T) {
	s := ""
	t := "y"
	// s := "asdfghj"
	// t := "fdghjasd"
	result := findTheDifference(s, t)
	fmt.Println(result)
	fmt.Println(string(result))
}
