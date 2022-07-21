/*
	Leetcode practices with go code.
*/
// 2: https://leetcode.cn/problems/add-two-numbers/
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用递归（栈）
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num := 0

	if l1 != nil {
		num += l1.Val
	}
	if l2 != nil {
		num += l2.Val
	}

	if num >= 10 {
		num = num % 10
		if l1 != nil && l1.Next != nil {
			l1.Next.Val += 1
		} else if l2 != nil && l2.Next != nil {
			l2.Next.Val += 1
		} else if l1 != nil && l1.Next == nil {
			l1.Next = &ListNode{1, nil}
		} else if l2 != nil && l2.Next == nil {
			l2.Next = &ListNode{1, nil}
		}
	}

	if l1 != nil && l2 != nil {
		return &ListNode{num, addTwoNumbers(l1.Next, l2.Next)}
	} else if l1 != nil {
		return &ListNode{num, addTwoNumbers(l1.Next, nil)}
	} else if l2 != nil {
		return &ListNode{num, addTwoNumbers(nil, l2.Next)}
	} else {
		return nil
	}
}

// 使用轮询
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) (result *ListNode) {
	result = &ListNode{0, nil}
	temp := result
	carry := 0

	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum := carry + x + y
		carry = sum / 10

		temp.Next = &ListNode{sum % 10, nil}
		temp = temp.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		temp.Next = &ListNode{carry, nil}
	}

	return result.Next
}

// 官方解法（轮询）
func addTwoNumbers2(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0

		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return
}
