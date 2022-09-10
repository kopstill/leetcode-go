// 206: https://leetcode.cn/problems/reverse-linked-list/
package leetcode

// 迭代
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 递归
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
