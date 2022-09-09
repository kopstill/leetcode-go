// 80: https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/
package leetcode

// 单指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeDuplicatesII(nums []int) int {
	// n := len(nums)
	// if n < 3 {
	// 	return n
	// }

	count := 1
	index := 0
	for index < len(nums)-1 {
		if nums[index] == nums[index+1] {
			count++
		} else {
			count = 1
		}
		if count > 2 {
			nums = append(nums[:index], nums[index+1:]...)
			count = 1
			index--
			continue
		}

		index++
	}

	return len(nums)
}

// 双指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeDuplicatesII1(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
