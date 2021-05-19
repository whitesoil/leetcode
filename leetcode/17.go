package leetcode

// ascii code: a:97, x: 120, y: 121, z: 122, 2: 50 9: 57
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	byteDigits := []byte(digits)
	combinations := [][]byte{{}}

	for _, c := range byteDigits {
		container := [][]byte{}

		chars := keyMap(c)

		for _, cb := range combinations {
			for _, char := range chars {
				newCombination := make([]byte, len(cb)+1)
				copy(newCombination, append(cb, char))
				container = append(container, newCombination)
			}
		}

		combinations = container
	}

	output := make([]string, len(combinations))

	for i, cb := range combinations {
		output[i] = string(cb)
	}

	return output
}

func keyMap(key byte) []byte {
	switch key {
	case 50:
		return []byte("abc")
	case 51:
		return []byte("def")
	case 52:
		return []byte("ghi")
	case 53:
		return []byte("jkl")
	case 54:
		return []byte("mno")
	case 55:
		return []byte("pqrs")
	case 56:
		return []byte("tuv")
	case 57:
		return []byte("wxyz")
	default:
		return []byte{}
	}
}
