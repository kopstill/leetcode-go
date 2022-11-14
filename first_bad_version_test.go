package leetcode

import (
	"fmt"
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	version := 97
	// version := 5
	// version := 1
	fmt.Println(firstBadVersion(version))
}
func Test1FirstBadVersion(t *testing.T) {
	// version := 97
	// version := 5
	version := 1
	fmt.Println(firstBadVersion1(version))
}
