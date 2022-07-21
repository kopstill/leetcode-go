// 1: https://leetcode.cn/problems/two-sum/
package leetcode

// 时间复杂度：O(n²)
// 空间复杂度：O(1)
func twoSum(nums []int, target int) []int {
	length := len(nums)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// 官方解法
// 时间复杂度：O(n²)
// 空间复杂度：O(1)
func twoSum1(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
