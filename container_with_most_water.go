// 11: https://leetcode.cn/problems/container-with-most-water/
package leetcode

// 暴力循环
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
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

// 双指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxArea1(height []int) int {
	l, r, maxArea := 0, len(height)-1, 0
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		maxArea = max(area, maxArea)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}
