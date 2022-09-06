// 69: https://leetcode.cn/problems/sqrtx/
package leetcode

// 暴力循环
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func mySqrt(x int) (res int) {
	if x == 0 || x == 1 {
		return x
	}

	for res = 1; res <= x; res++ {
		if res*res == x {
			return res
		} else if res*res > x {
			return res - 1
		}
	}

	return
}
