package leetcode

import "bytes"

// Convert is 6th problem
func Convert(s string, numRows int) string {
	twoDim := make([]string, numRows)
	reverse := false
	index := 0

	for i := 0; i < len(s); i++ {
		var buffer bytes.Buffer

		buffer.WriteString(twoDim[index])
		buffer.WriteString(string(s[i]))

		twoDim[index] = buffer.String()

		if numRows == 1 {
			continue
		}

		if index == numRows-1 {
			reverse = true
		} else if index == 0 {
			reverse = false
		}

		if reverse {
			index--
		} else {
			index++
		}
	}

	var buffer bytes.Buffer

	for i := 0; i < len(twoDim); i++ {
		buffer.WriteString(twoDim[i])
	}

	return buffer.String()
}
