package leetcode

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	// t1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	// t2 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	t1 := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}
	t2 := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{8, nil, nil}}}
	fmt.Println(isSameTree(t1, t2))
}

func Test1IsSameTree(t *testing.T) {
	// t1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	// t2 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	t1 := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4,
				&TreeNode{8, nil, nil},
				&TreeNode{9, nil, nil}},
			&TreeNode{5,
				&TreeNode{10, nil, nil},
				&TreeNode{11, nil, nil}}},
		&TreeNode{3,
			&TreeNode{6,
				&TreeNode{12, nil, nil},
				&TreeNode{13, nil, nil}},
			&TreeNode{7,
				&TreeNode{14, nil, nil},
				&TreeNode{15, nil, nil}}},
	}
	t2 := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4,
				&TreeNode{8, nil, nil},
				&TreeNode{9, nil, nil}},
			&TreeNode{5,
				&TreeNode{10, nil, nil},
				&TreeNode{11, nil, nil}}},
		&TreeNode{3,
			&TreeNode{6,
				&TreeNode{12, nil, nil},
				&TreeNode{13, nil, nil}},
			&TreeNode{7,
				&TreeNode{14, nil, nil},
				&TreeNode{15, nil, nil}}},
	}
	fmt.Println(isSameTree1(t1, t2))
}
