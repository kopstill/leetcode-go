// 344: https://leetcode.cn/problems/reverse-string/
package leetcode

// 未优化的双指针
// 时间复杂度：O(N)。
// 空间复杂度：O(1)。
func reverseString(s []byte) {
	q := len(s) - 1
	tmp := byte(0)
	for p := 0; p < len(s); p++ {
		if p < q {
			tmp = s[p]
			s[p] = s[q]
			s[q] = tmp
			q--
		}
	}
}

// 双指针
// 时间复杂度：O(N)，其中 N 为字符数组的长度。一共执行了 N/2 次的交换。
// 空间复杂度：O(1)。只使用了常数空间来存放若干变量。
func reverseString1(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}
