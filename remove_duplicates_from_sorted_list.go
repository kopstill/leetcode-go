// 83: https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
package leetcode

// 一次遍历
// 时间复杂度：O(n)，其中 n 是链表的长度。
// 空间复杂度：O(1)。
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}
