package leetcode

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	new := 0
	for {
		last := x % 10
		x = x / 10
		new = new*10 + last
		if new == x {
			return true
		}
		if new > x {
			if new/10 == x {
				return true
			}
			break
		}
	}

	return false
}
