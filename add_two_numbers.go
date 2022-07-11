package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

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
