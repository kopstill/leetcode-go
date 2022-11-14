// 278: https://leetcode.cn/problems/first-bad-version/
package leetcode

import "sort"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// 二分查找
// 时间复杂度：O(log n)
// 空间复杂度：O(1)
func firstBadVersion(n int) int {
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

var BadVersion int = 1

func isBadVersion(version int) bool {
	return version >= BadVersion
}

// 二分查找
// 时间复杂度：O(log n)，其中 n 是给定版本的数量。
// 空间复杂度：O(1)。我们只需要常数的空间保存若干变量。
func firstBadVersion1(n int) int {
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}
