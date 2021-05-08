package leetcode

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	var output string

	for index, runeChar := range strs[0] {
		char := string(runeChar)

		for i := 1; i < len(strs); i++ {
			if len(strs[i]) <= index || char != string(strs[i][index]) {
				return output
			}
		}

		output += char
	}

	return output
}
