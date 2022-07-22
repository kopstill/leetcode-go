// 6: https://leetcode.cn/problems/zigzag-conversion/
package leetcode

func convert(s string, numRows int) string {
	if len(s) < 2 || numRows < 2 {
		return s
	}

	t := make([][]string, numRows)

	p := 0
	q := 0
	for i := 0; i < len(s); i++ {
		x := append(t[p], string(s[i]))
		t[p] = x

		m := i % (numRows - 1)
		q = q % 2

		if i == 0 {
			p++
		} else {
			if m == 0 && q == 0 {
				p--
				q++
			} else if m == 0 && q != 0 {
				p++
				q--
			} else if m != 0 && q == 0 {
				p++
			} else {
				p--
			}
		}
	}

	var r string
	for _, v := range t {
		for _, vv := range v {
			r += vv
		}
	}

	return r
}
