// 70: https://leetcode.cn/problems/climbing-stairs/solution/
package leetcode

// 动态规划
// 时间复杂度：循环执行 n 次，每次花费常数的时间代价，故渐进时间复杂度为 O(n)。
// 空间复杂度：这里只用了常数个变量作为辅助空间，故渐进空间复杂度为 O(1)。
func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
