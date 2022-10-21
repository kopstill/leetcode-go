// 389: https://leetcode.cn/problems/find-the-difference/
package leetcode

func findTheDifference(s, t string) (result byte) {
	tLength := len(t)
	tarr := make([]byte, tLength)
	for i := 0; i < tLength; i++ {
		tarr[i] = t[i]
	}

	for i := 0; i < len(s); i++ {
		if ok, idx := isContain(s[i], tarr); ok {
			tarr = append(tarr[:idx], tarr[idx+1:]...)
		}
	}

	return tarr[0]
}

func isContain(target byte, arr []byte) (bool, int) {
	for i, v := range arr {
		if v == target {
			return true, i
		}
	}
	return false, -1
}
