package leetcode

import "strings"

// LengthOfLongestSubstring is 3rd problem
func LengthOfLongestSubstring(s string) int {
	chars := strings.Split(s, "")
	length := len(chars)
	result := 0
	i := 0
	j := 0

	container := make(map[string]int)

	for j < length {
		if length-i <= result {
			break
		}

		val, exist := container[chars[j]]

		if exist && i <= val {
			temp := j - i

			if temp > result {
				result = temp
			}

			i = val + 1
		} else {
			container[chars[j]] = j
			j++
		}
	}

	temp := j - i
	if temp > result {
		result = temp
	}

	return result
}
