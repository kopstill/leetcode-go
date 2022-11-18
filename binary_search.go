// 704: https://leetcode.cn/problems/binary-search/
package leetcode

// 二分查找
// 时间复杂度：O(logn)，其中 n 是数组的长度。
// 空间复杂度：O(1)。
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
