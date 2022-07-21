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

// 哈希表
// 时间复杂度：O(N)，单层循环，哈希表查找时间复杂度为 O(1)
// 空间复杂度：O(N)，主要是哈希表存储的数据
func twoSum2(nums []int, target int) []int {
	m := map[int]int{}

	for i, v := range nums {
		if p, ok := m[target-v]; ok {
			return []int{p, i}
		}
		m[v] = i
	}

	return nil
}
