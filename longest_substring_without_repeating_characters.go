package leetcode

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 || length == 1 {
		return length
	}

	max := 1

	r := []rune(s)

	t := []string{}

	for i := 0; i < len(r); i++ {
		t = append(t, string(r[i]))
		for j := i + 1; j < len(r); j++ {
			if isRepeat(t, string(r[j])) {
				t = nil
				break
			} else {
				t = append(t, string(r[j]))
			}
			if len(t) > max {
				max = len(t)
			}
		}
	}

	return max
}

func isRepeat(slc []string, r string) (flag bool) {
	for _, sv := range slc {
		if sv == r {
			return true
		}
	}

	return
}
