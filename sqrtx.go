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

// 二分查找
// 时间复杂度：O(logn)
// 空间复杂度：O(1)
func mySqrt2(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
