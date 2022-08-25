// 21: https://leetcode.cn/problems/merge-two-sorted-lists/
package leetcode

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func mergeTwoSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{0, nil}
	tail := result

	for list1 != nil || list2 != nil {
		if list1 != nil && list2 == nil {
			tail.Next = list1
			list1 = nil
		} else if list1 == nil && list2 != nil {
			tail.Next = list2
			list2 = nil
		} else if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				tail.Next = &ListNode{list1.Val, nil}
				list1 = list1.Next
			} else {
				tail.Next = &ListNode{list2.Val, nil}
				list2 = list2.Next
			}
			tail = tail.Next
		}
	}

	return result.Next
}
