// 58: https://leetcode.cn/problems/length-of-last-word/
package leetcode

// 先去除末尾空格，然后计算末尾非空格字符数
// 时间复杂度：O(2n)
// 空间复杂度：O(1)
func lengthOfLastWord(s string) int {
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			s = s[:i+1]
			break
		}
	}

	n = len(s)
	for i := n - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			return n - i - 1
		}
	}

	return n
}
