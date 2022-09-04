package leetcode

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	head := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}}
	res := deleteDuplicates(head)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
