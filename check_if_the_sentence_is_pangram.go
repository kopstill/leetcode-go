// 1832: https://leetcode.cn/problems/check-if-the-sentence-is-pangram/
package leetcode

// 暴力轮询
// 时间复杂度：O(m*n)
// 空间复杂度：O(1)
func checkIfPangram(sentence string) bool {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(sentence); i++ {
		if len(alphabet) == 0 {
			return true
		}
		for j := 0; j < len(alphabet); j++ {
			if sentence[i] == alphabet[j] {
				alphabet = alphabet[:j] + alphabet[j+1:]
				break
			}
		}
	}

	return len(alphabet) == 0
}
