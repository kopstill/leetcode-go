package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	// strs := []string{"hello", "hiya", "halo", "hey"}
	// strs := []string{}
	// strs := []string{"   "}
	strs := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(strs))
}
