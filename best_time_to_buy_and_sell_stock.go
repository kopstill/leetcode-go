// 121: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
package leetcode

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
