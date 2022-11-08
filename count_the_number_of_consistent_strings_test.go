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

func Test1CountConsistentStrings(t *testing.T) {
	allowed := "test"
	words := []string{"testestest", "tteesstt", "leetcode", "hello", "world", "set"}
	fmt.Println(countConsistentStrings1(allowed, words))
}
