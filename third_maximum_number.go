// 414: https://leetcode.cn/problems/third-maximum-number/
package leetcode

import (
	"math"
	"sort"

	"github.com/emirpasic/gods/trees/redblacktree"
)

// 排序
// 时间复杂度：O(nlogn)，其中 n 是数组 nums 的长度。排序需要 O(nlogn) 的时间。
// 空间复杂度：O(logn)。排序需要的栈空间为 O(logn)。
func thirdMax(nums []int) int {
	sort.Ints(nums)

	len := len(nums)
	idx := 1
	for i := len - 1; i > 0; i-- {
		if nums[i] != nums[i-1] {
			idx++
		}
		if idx == 3 {
			return nums[i-1]
		}
	}

	return nums[len-1]
}

// 排序
// 时间复杂度：O(nlogn)，其中 n 是数组 nums 的长度。排序需要 O(nlogn) 的时间。
// 空间复杂度：O(logn)。排序需要的栈空间为 O(logn)。
func thirdMax1(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, diff := 1, 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			diff++
			if diff == 3 {
				return nums[i]
			}
		}
	}
	return nums[0]
}

// 有序集合
// 时间复杂度：O(n)，其中 n 是数组 nums 的长度。由于有序集合的大小至多为 3，插入和删除的时间复杂度可以视作是 O(1) 的，因此时间复杂度为 O(n)。
// 空间复杂度：O(1)。
func thirdMax2(nums []int) int {
	t := redblacktree.NewWithIntComparator()
	for _, num := range nums {
		t.Put(num, nil)
		if t.Size() > 3 {
			t.Remove(t.Left().Key)
		}
	}
	if t.Size() == 3 {
		return t.Left().Key.(int)
	}
	return t.Right().Key.(int)
}

// 一次遍历
// 时间复杂度：O(n)，其中 n 是数组 nums 的长度。
// 空间复杂度：O(1)。
func thirdMax3(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > a {
			a, b, c = num, a, b
		} else if a > num && num > b {
			b, c = num, b
		} else if b > num && num > c {
			c = num
		}
	}
	if c == math.MinInt64 {
		return a
	}
	return c
}

// 一次遍历
// 时间复杂度：O(n)，其中 n 是数组 nums 的长度。
// 空间复杂度：O(1)。
func thirdMax4(nums []int) int {
	var a, b, c *int
	for _, num := range nums {
		num := num
		if a == nil || num > *a {
			a, b, c = &num, a, b
		} else if *a > num && (b == nil || num > *b) {
			b, c = &num, b
		} else if b != nil && *b > num && (c == nil || num > *c) {
			c = &num
		}
	}
	if c == nil {
		return *a
	}
	return *c
}
