package main

// func findLogestPalindrome(chars []string) []string {
// 	length := len(chars)

// 	if length < 2 {
// 		return chars
// 	}

// 	result := true

// 	for i := 0; i < (length / 2); i++ {
// 		compare := strings.Compare(chars[i], chars[length-i-1])

// 		if compare != 0 {
// 			result = false
// 			break
// 		}
// 	}

// 	if result {
// 		return chars
// 	}

// 	opt1 := findLogestPalindrome(chars[0 : length-1])

// 	if len(opt1) == length-1 {
// 		return opt1
// 	}

// 	opt2 := findLogestPalindrome(chars[1:])

// 	if len(opt2) == length-1 {
// 		return opt2
// 	}

// 	if len(opt1) >= len(opt2) {
// 		return opt1
// 	} else {
// 		return opt2
// 	}
// }

// func longestPalindrome(s string) string {
// 	chars := strings.Split(s, "")
// 	result := findLogestPalindrome(chars)

// 	return strings.Join(result, "")
// }

// P를 Palindrome 이라 칭한다.
// 중심을 기준으로, 2n-1 (n > 1) 가 P 이기 위해서는 2(n-1)-1 또한 P여야 한다.
// 중심을 기준으로, 2n (n > 1) 가 P 이기 위해서는 2(n-1)
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

func longestPalindrome(s string) string {
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

func main() {
	print(longestPalindrome("babaddtattarrattatddetartrateedredividerb"))
}
