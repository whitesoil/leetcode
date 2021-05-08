package leetcode

func RomanToInt(s string) int {
	output := 0
	var prev string

	for _, char := range s {
		switch string(char) {
		case "I":
			output += 1
			prev = "I"
		case "V":
			if prev == "I" {
				output += 3
				prev = ""
			} else {
				output += 5
			}
		case "X":
			if prev == "I" {
				output += 8
				prev = ""
			} else {
				output += 10
				prev = "X"
			}
		case "L":
			if prev == "X" {
				output += 30
				prev = ""
			} else {
				output += 50
			}
		case "C":
			if prev == "X" {
				output += 80
				prev = ""
			} else {
				output += 100
				prev = "C"
			}
		case "D":
			if prev == "C" {
				output += 300
				prev = ""
			} else {
				output += 500
			}
		case "M":
			if prev == "C" {
				output += 800
				prev = ""
			} else {
				output += 1000
				prev = "M"
			}
		}

	}

	return output
}
