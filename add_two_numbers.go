package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

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
	} else {
		return &ListNode{num, addTwoNumbers(nil, l2.Next)}
	}
}
