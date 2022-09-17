// 11: https://leetcode.cn/problems/container-with-most-water/
package leetcode

// 暴力循环
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func maxArea(height []int) int {
	n, maxArea := len(height), 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			h := min(height[i], height[j])
			w := j - i
			area := h * w
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
