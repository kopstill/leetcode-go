// 94: https://leetcode.cn/problems/binary-tree-inorder-traversal/
package leetcode

// 递归：隐式维护栈
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

// 迭代：显式维护栈
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func inorderTraversal1(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}
