// 67: https://leetcode.cn/problems/add-binary/
package leetcode

import (
	"fmt"
	"strconv"
)

// 模拟相加
// 时间复杂度：O(max(m,n))
// 空间复杂度：O(1)
func addBinary(a string, b string) string {
	m, n := len(a), len(b)

	res := ""
	x := 0
	for m > 0 || n > 0 {
		r := 0
		if m > 0 {
			p, _ := strconv.Atoi(string(a[m-1]))
			r += p
		}
		if n > 0 {
			q, _ := strconv.Atoi(string(b[n-1]))
			r += q
		}
		r += x

		if r > 1 {
			if r%2 == 1 {
				res = "1" + res
			} else {
				res = "0" + res
			}
			x = 1
		} else {
			res = fmt.Sprint(r) + res
			x = 0
		}
		m--
		n--
	}

	if x == 1 {
		res = "1" + res
	}

	return res
}

// 官方：模拟
// 复杂度分析
// 假设 n = max{|a|,|b|}。
// 时间复杂度：O(n)，这里的时间复杂度来源于顺序遍历 a 和 b。
// 空间复杂度：O(1)，除去答案所占用的空间，这里使用了常数个临时变量。
func addBinary1(a, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
