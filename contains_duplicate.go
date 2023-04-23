// 217: https://leetcode.cn/problems/contains-duplicate/
package leetcode

// hash
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
