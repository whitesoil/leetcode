package main

import "bytes"

// l: 문자열 길이, r: number of rows, i: index of characters
// 1번 라인 갯수: Ceil(l / (r+1))
func convert(s string, numRows int) string {
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

		if reverse == true {
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

func main() {
	print(convert("AB", 1))
}
