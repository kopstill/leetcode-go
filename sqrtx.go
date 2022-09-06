// 69: https://leetcode.cn/problems/sqrtx/
package leetcode

import "math"

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

// 袖珍计算器算法
// 时间复杂度：O(1)
// 空间复杂度：O(1)
func mySqrt1(x int) int {
	if x == 0 {
		return 0
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans
}
