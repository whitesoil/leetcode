package leetcode

// '*' == 42, '.' == 46
func isStringMatch(s string, p string, sIndex int, pIndex int) bool {
	// P index is end
	if pIndex == len(p) {
		return sIndex == len(s)
	}

	// S index is end
	if sIndex == len(s) {
		if pIndex == len(p)-1 {
			return false
		}

		if pIndex < len(p)-1 && p[pIndex+1] != 42 {
			return false
		}
	}

	// Next character is '*'
	if pIndex < len(p)-1 && p[pIndex+1] == 42 {
		// p is '.'
		if p[pIndex] == 46 {
			for i := 0; i <= len(s)-sIndex; i++ {
				res := isStringMatch(s, p, sIndex+i, pIndex+2)

				if res {
					return true
				}
			}

			return false
		}

		for i := 0; i <= len(s)-sIndex; i++ {
			res := isStringMatch(s, p, sIndex+i, pIndex+2)

			if res {
				return true
			}

			if i < len(s)-sIndex && s[sIndex+i] != p[pIndex] {
				break
			}
		}

		return false
	}

	// p is '.' OR p and s is same
	if p[pIndex] == 46 || s[sIndex] == p[pIndex] {
		return isStringMatch(s, p, sIndex+1, pIndex+1)
	}

	return false
}

// IsMatch is 10th problem
func IsMatch(s string, p string) bool {
	return isStringMatch(s, p, 0, 0)
}
