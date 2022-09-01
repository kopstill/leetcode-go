// https://www.hackerrank.com/x/tests/545682/candidates/44199341/report
package hackerrank

// 时间复杂度：O(n*m)
// 空间复杂度：O(1)
func gemstones(rocks []string) int32 {
	// Write your code here
	n := len(rocks)
	if n == 0 {
		return 0
	}

	counter := make(map[string]int)
	for i := 0; i < n; i++ {
		rock := rocks[i]
		for j := 0; j < len(rock); j++ {
			if i == 0 {
				counter[string(rock[j])] = 1
			} else {
				if counter[string(rock[j])] == i {
					counter[string(rock[j])] = i + 1
				}
			}
		}
	}

	var result int32 = 0
	for _, v := range counter {
		if v == n {
			result++
		}
	}
	return result
}
