// 414: https://leetcode.cn/problems/fizz-buzz/
package leetcode

import (
	"strconv"
	"strings"
)

// 时间复杂度：O(n)
// 空间复杂度：O(n)。需要长度为 n 的数组存储返回值。
func fizzBuzz(n int) []string {
	result := []string{}

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}

// 时间复杂度：O(n)
// 空间复杂度：O(1)。注意返回值不计入空间复杂度。
func fizzBuzz1(n int) (ans []string) {
	for i := 1; i <= n; i++ {
		sb := &strings.Builder{}
		if i%3 == 0 {
			sb.WriteString("Fizz")
		}
		if i%5 == 0 {
			sb.WriteString("Buzz")
		}
		if sb.Len() == 0 {
			sb.WriteString(strconv.Itoa(i))
		}
		ans = append(ans, sb.String())
	}
	return
}
