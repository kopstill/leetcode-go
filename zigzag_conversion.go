// 6: https://leetcode.cn/problems/zigzag-conversion/
package leetcode

import "bytes"

// 此解法不会浪费二维数组的空间
// 时间复杂度：O(r * n)，二维数组的创建和遍历
// 空间复杂度：O(r * n)，二维数组占用的空间
func convert(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}

	t := make([][]string, r)

	p, q := 0, 0
	for i := 0; i < n; i++ {
		t[p] = append(t[p], string(s[i]))

		m := i % (r - 1)
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

	var res string
	for _, v := range t {
		for _, s := range v {
			res += s
		}
	}

	return res
}

// 基于以上解法创建二维字节矩阵优化时间和空间复杂度（无需手动遍历矩阵）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func convert1(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}

	t := make([][]byte, r)

	p, q := 0, 0
	for i := 0; i < n; i++ {
		t[p] = append(t[p], s[i])

		m := i % (r - 1)
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

	return string(bytes.Join(t, nil))
}

// 官方解法 1：利用二维矩阵模拟
// 时间复杂度：O(r⋅n)，其中 r = numRows，n 为字符串 s 的长度。时间主要消耗在矩阵的创建和遍历上，矩阵的行数为 r，列数可以视为 O(n)。
// 空间复杂度：O(r⋅n)。矩阵需要 O(r⋅n) 的空间。
func convert2(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}
	t := r*2 - 2
	c := (n + t - 1) / t * (r - 1)
	mat := make([][]byte, r)
	for i := range mat {
		mat[i] = make([]byte, c)
	}
	x, y := 0, 0
	for i, ch := range s {
		mat[x][y] = byte(ch)
		if i%t < r-1 {
			x++ // 向下移动
		} else {
			x--
			y++ // 向右上移动
		}
	}
	ans := make([]byte, 0, n)
	for _, row := range mat {
		for _, ch := range row {
			if ch > 0 {
				ans = append(ans, ch)
			}
		}
	}
	return string(ans)
}

// 官方解法 2：压缩矩阵空间
// 时间复杂度：O(n)
// 空间复杂度：O(n)。压缩后的矩阵需要 O(n) 的空间。
func convert3(s string, numRows int) string {
	r := numRows
	if r == 1 || r >= len(s) {
		return s
	}
	mat := make([][]byte, r)
	t, x := r*2-2, 0
	for i, ch := range s {
		mat[x] = append(mat[x], byte(ch))
		if i%t < r-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}
