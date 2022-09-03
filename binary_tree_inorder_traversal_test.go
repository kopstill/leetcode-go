package leetcode

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tree := &TreeNode{1,
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
	fmt.Println(inorderTraversal(tree))
}

func Test1InorderTraversal(t *testing.T) {
	tree := &TreeNode{1,
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
	fmt.Println(inorderTraversal1(tree))
}
