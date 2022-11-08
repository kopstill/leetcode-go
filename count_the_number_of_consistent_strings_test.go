package leetcode

import (
	"fmt"
	"testing"
)

func TestCountConsistentStrings(t *testing.T) {
	allowed := "abc"
	words := []string{"cba", "bac", "cab", "abcd", "dbac", "ccc", "aaa", "bbb"}
	fmt.Println(countConsistentStrings(allowed, words))
}
