// 217: https://leetcode.cn/problems/contains-duplicate/
package leetcode

// Hash
// Time  complexity: O(n)
// Space complexity: O(n)
func containsDuplicate(nums []int) bool {
	hash := map[int]int{}
	for i, v := range nums {
		if _, ok := hash[v]; ok {
			return true
		}
		hash[v] = i
	}
	return false
}

// For loop(exceeding the time limit)
// Time  complexity: O(n * n)
// Space complexity: O(n * n)
func containsDuplicate1(nums []int) bool {
	len := len(nums)
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
