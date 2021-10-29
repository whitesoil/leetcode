package leetcode

func StrStr(haystack string, needle string) int {
	lenOfNeedle := len(needle)

	if lenOfNeedle == 0 {
		return 0
	}

	lenOfHaystack := len(haystack)

	if lenOfHaystack == 0 {
		return -1
	}

	runesOfHaystack := []rune(haystack)
	runesOfNeedle := []rune(needle)

	for i := 0; i <= lenOfHaystack-lenOfNeedle; i++ {
		isEqual := equal(runesOfHaystack[i:i+lenOfNeedle], runesOfNeedle)

		if isEqual {
			return i
		}
	}

	return -1
}

func equal(a []rune, b []rune) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
