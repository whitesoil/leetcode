package leetcode

func IsValid(s string) bool {
	chars := []byte(s)
	stack := make([]byte, 0, len(s))

	for _, c := range chars {
		switch c {
		case 40: // (
			stack = append(stack, 40)
		case 41: // )
			if len(stack) == 0 || stack[len(stack)-1] != 40 {
				return false
			}
			stack = stack[:len(stack)-1]
		case 91: // [
			stack = append(stack, 91)
		case 93: // ]
			if len(stack) == 0 || stack[len(stack)-1] != 91 {
				return false
			}
			stack = stack[:len(stack)-1]
		case 123: // {
			stack = append(stack, 123)
		case 125: // }
			if len(stack) == 0 || stack[len(stack)-1] != 123 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
