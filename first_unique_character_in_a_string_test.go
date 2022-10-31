package leetcode

import (
	"fmt"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	// s := "leetcode"
	s := "aabb"
	fmt.Println(firstUniqChar(s))
}

func Test1FirstUniqChar(t *testing.T) {
	// s := "leetcode"
	// s := "aabb"
	// s := ""
	s := "xabcx"
	fmt.Println(firstUniqChar1(s))
}

func Test2FirstUniqChar(t *testing.T) {
	s := "leetcode"
	fmt.Println(firstUniqChar2(s))
}

func Test3FirstUniqChar(t *testing.T) {
	s := "leetcode"
	fmt.Println(firstUniqChar3(s))
}

func Test4FirstUniqChar(t *testing.T) {
	s := "leetcode"
	fmt.Println(firstUniqChar4(s))
}
