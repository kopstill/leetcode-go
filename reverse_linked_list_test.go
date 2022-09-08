package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	result := reverseList(head)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
