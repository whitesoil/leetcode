package leetcode

// Reverse is 7th problem
func Reverse(x int) int {
	num := x
	result := 0
	for num != 0 {
		result = result*10 + num%10

		if result > 2147483647 || result < -2147483648 {
			return 0
		}

		num = num / 10
	}

	return result
}
