package leetcode

import (
	"fmt"
	"testing"
)

func TestToLowerCase(t *testing.T) {
	s := "LeetCode"
	fmt.Println(toLowerCase(s))
}

func Test1ToLowerCase(t *testing.T) {
	s := "Hello World!!!"
	fmt.Println(toLowerCase1(s))
}

func Test2ToLowerCase(t *testing.T) {
	s := "LOVELY"
	fmt.Println(toLowerCase2(s))
}
