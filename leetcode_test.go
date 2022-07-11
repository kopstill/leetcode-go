package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 5, 1, 6, 9, 7}
	target := 15
	result := twoSum(nums, target)
	fmt.Println(result)
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	// l1 := ListNode{4, &ListNode{9, &ListNode{5, nil}}}
	// l2 := ListNode{3, &ListNode{2, &ListNode{5, nil}}}

	// l1 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	// l2 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}

	result := addTwoNumbers(&l1, &l2)
	for {
		fmt.Println(result.Val)
		if result.Next != nil {
			result = result.Next
		} else {
			break
		}
	}
}

func TestAddTwoNumbers1(t *testing.T) {
	// l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	// l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	l1 := ListNode{4, &ListNode{9, &ListNode{5, nil}}}
	l2 := ListNode{3, &ListNode{2, &ListNode{5, nil}}}

	// l1 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	// l2 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}

	result := addTwoNumbers1(&l1, &l2)
	for {
		fmt.Println(result.Val)
		if result.Next != nil {
			result = result.Next
		} else {
			break
		}
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	// l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	// l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	// l1 := ListNode{4, &ListNode{9, &ListNode{5, nil}}}
	// l2 := ListNode{3, &ListNode{2, &ListNode{5, nil}}}

	l1 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}

	result := addTwoNumbers2(&l1, &l2)
	for {
		fmt.Println(result.Val)
		if result.Next != nil {
			result = result.Next
		} else {
			break
		}
	}
}
