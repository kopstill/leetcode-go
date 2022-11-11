// 448: https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/
package leetcode

import (
	"github.com/emirpasic/gods/sets/hashset"
)

// set
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func findDisappearedNumbers(nums []int) []int {
	set := hashset.New()
	for _, v := range nums {
		set.Add(v)
	}

	result := []int{}
	for i := 1; i <= len(nums); i++ {
		if !set.Contains(i) {
			result = append(result, i)
		}
	}

	return result
}

// hash
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func findDisappearedNumbers1(nums []int) (result []int) {
	hash := map[int]int{}
	for i, v := range nums {
		hash[v] = i
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := hash[i]; ok {
			continue
		}
		result = append(result, i)
	}
	return
}

// 遍历
// 时间复杂度：O(n * n)
// 空间复杂度：O(1)
func findDisappearedNumbers2(nums []int) []int {
	var result []int
	for i := 1; i <= len(nums); i++ {
		flag := true
		for _, v := range nums {
			if v == i {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, i)
		}
	}
	return result
}

// 原地修改
// 时间复杂度：O(n)。其中 n 是数组 nums 的长度。
// 空间复杂度：O(1)。返回值不计入空间复杂度。
func findDisappearedNumbers3(nums []int) (ans []int) {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return
}
