// 374: https://leetcode.cn/problems/guess-number-higher-or-lower/
package leetcode

import "sort"

var pick int = 42

// 二分
func guessNumber(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == -1 {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func guess(num int) int {
	if pick > num {
		return 1
	} else if pick < num {
		return -1
	} else {
		return 0
	}
}

// 二分，语言自带二分查找 API
func guessNumber1(n int) int {
	return sort.Search(n, func(x int) bool { return guess(x) <= 0 })
}
