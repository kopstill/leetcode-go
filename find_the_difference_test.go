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

func Test1FindTheDifference(test *testing.T) {
	// s := ""
	// t := "y"
	s := "asdfghj"
	t := "fdghjasd"
	fmt.Println(string(findTheDifference1(s, t)))
}

func Test2FindTheDifference(test *testing.T) {
	s := "aaabbbcccxxxddd"
	t := "dddxxxibbbaaaccc"
	fmt.Println(string(findTheDifference2(s, t)))
}
