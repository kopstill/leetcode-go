package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	// height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	height := []int{1, 1}
	fmt.Println(maxArea(height))
}

func Test1MaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// height := []int{1, 1}
	fmt.Println(maxArea1(height))
}
