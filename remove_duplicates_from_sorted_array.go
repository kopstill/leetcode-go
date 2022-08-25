// 26: https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
package leetcode

// func removeDuplicates(nums []int) int {
// 	for i, num := range nums {
// 		if i < len(nums)-1 && num == nums[i+1] {
// 			nums = append(nums[:i], nums[i+1:]...)
// 			removeDuplicates(nums)
// 		}
// 	}

// 	return len(nums)
// }

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
