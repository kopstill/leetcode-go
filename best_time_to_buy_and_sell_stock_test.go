package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{10, 2, 30, 4, 2, 1, 100}
	fmt.Println(maxProfit(prices))
}

func Test1MaxProfit(t *testing.T) {
	// prices := []int{10, 5, 6, 1, 20, 8, 30}
	prices := []int{2, 1, 2, 0, 1}
	fmt.Println(maxProfit1(prices))
}
