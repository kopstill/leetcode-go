package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	// l1 := ListNode{4, &ListNode{9, &ListNode{5, nil}}}
	// l2 := ListNode{3, &ListNode{2, &ListNode{5, nil}}}

	// l1 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	// l2 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}

	result := addTwoNumbers(&l1, &l2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func Test1AddTwoNumbers(t *testing.T) {
	// l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	// l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	l1 := ListNode{4, &ListNode{9, &ListNode{5, nil}}}
	l2 := ListNode{3, &ListNode{2, &ListNode{5, nil}}}

	// l1 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	// l2 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}

	result := addTwoNumbers1(&l1, &l2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func Test2AddTwoNumbers(t *testing.T) {
	// l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	// l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	// l1 := ListNode{4, &ListNode{9, &ListNode{5, nil}}}
	// l2 := ListNode{3, &ListNode{2, &ListNode{5, nil}}}

	l1 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}

	result := addTwoNumbers2(&l1, &l2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
