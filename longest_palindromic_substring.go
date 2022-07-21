package leetcode

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	r := s[0:1]
	for i := 0; i < len(s)-1; i++ {
		p1, p2 := i, len(s)-1
		c := 0
		t := ""
		for {
			if s[p1] == s[p2] {
				c++
				p1++
				p2--
			} else {
				p1 = i
				p2 = p2 + c - 1
				c = 0
			}
			if p1 >= p2 {
				break
			}
		}
		if c != 0 {
			if p1 > p2 {
				t = s[i : i+c*2]
			} else {
				t = s[i : i+c*2+1]
			}
			if len(t) > len(r) {
				r = t
			}
		}
	}

	return r
}
