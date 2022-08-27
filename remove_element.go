// 27: https://leetcode.cn/problems/remove-element/
package leetcode

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeElement(nums []int, val int) int {
	p := 0
	for p < len(nums) {
		if nums[p] == val {
			nums = append(nums[:p], nums[p+1:]...)
		} else {
			p++
		}
	}

	return len(nums)
}
