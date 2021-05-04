package leetcode

import "strconv"

// IsPalindrome is 9th problem
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)

	i := 0
	j := len(str) - 1

	for i <= j {
		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}

	return true
}
