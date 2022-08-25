package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	list1 := ListNode{1, &ListNode{3, &ListNode{5, &ListNode{7, &ListNode{9, nil}}}}}
	list2 := ListNode{2, &ListNode{4, &ListNode{6, &ListNode{8, &ListNode{10, nil}}}}}
	// list2 := ListNode{2, &ListNode{4, &ListNode{6, &ListNode{8, &ListNode{10, &ListNode{100, nil}}}}}}
	// resultList := mergeTwoSortedLists(nil, nil)
	resultList := mergeTwoSortedLists(&list1, &list2)
	for resultList != nil {
		fmt.Println(resultList.Val)
		resultList = resultList.Next
	}
}
