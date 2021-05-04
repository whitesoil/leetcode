package leetcode

func findLogestPalindrome(s string, left int, right int) (int, int) {
	i := left
	j := right

	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			break
		}

		i--
		j++
	}

	return i + 1, j - 1
}

// LongestPalindrome is 5th problem
func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	left := 0
	right := 0

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			i, j := findLogestPalindrome(s, i, i+1)
			if (right - left) < (j - i) {
				left = i
				right = j
			}
		}

		if i > 0 && (s[i-1] == s[i+1]) {
			i, j := findLogestPalindrome(s, i-1, i+1)
			if (right - left) < (j - i) {
				left = i
				right = j
			}
		}
	}

	return s[left:(right + 1)]
}
