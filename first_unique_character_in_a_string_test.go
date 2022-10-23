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
