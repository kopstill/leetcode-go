// 94: https://leetcode.cn/problems/binary-tree-inorder-traversal/
package leetcode

// 递归
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
