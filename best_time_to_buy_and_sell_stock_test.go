package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{10, 2, 30, 4, 2, 1, 100}
	fmt.Println(maxProfit(prices))
}
