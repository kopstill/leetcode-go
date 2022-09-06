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
