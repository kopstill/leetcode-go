// 27: https://leetcode.cn/problems/remove-element/
package leetcode

// 单指针
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

// 双指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeElement1(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}

	return left
}

// 双指针优化：至多遍历一次 nums
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeElement2(nums []int, val int) int {
	left, right := 0, len(nums)

	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}

	return left
}
