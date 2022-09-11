// 121: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
package leetcode

import "math"

// 暴力循环
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	maxProfit := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if prices[i] < prices[j] {
				currProfit := prices[j] - prices[i]
				if currProfit > maxProfit {
					maxProfit = currProfit
				}
			}
		}
	}

	return maxProfit
}

// 一次遍历
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxProfit1(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	minPrice := math.MaxInt
	maxProfit := 0
	for i := 0; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}
