package leetcode

// MyAtoi is 8th problem
// "0" == 48, "9" == 57, "+" == 43, "-" == 45, " " == 32
func MyAtoi(str string) int {
	sign := 0
	start := false
	result := 0

	for i := 0; i < len(str); i++ {
		char := str[i]

		if char != 32 && char != 43 && char != 45 && (char < 48 || 57 < char) {
			// 공백이 아닌 문자는 항상 종료한다.
			break
		}

		if (char == 43 || char == 45) && start == false {
			// 부호 뒤에는 반드시 숫자가 와야한다
			// +-, -+ 허용 안한다
			start = true
			sign = int(char)
			continue
		}

		if 48 <= char && char <= 57 {
			result = result*10 + int(char-48)

			// overflow의 경우 2147483647 또는 -2147483648 리턴
			if (sign == 0 || sign == 43) && result > 2147483647 {
				return 2147483647
			}

			// overflow의 경우 2147483647 또는 -2147483648 리턴
			if sign == 45 && result > 2147483648 {
				return -2147483648
			}

			start = true
			continue
		}

		if (char < 48 || 57 < char) && start == true {
			// 공백은 숫자 앞에만 허용 된다
			// 부호 뒤에는 반드시 숫자가 와야한다
			// +-, -+ 허용 안한다
			// 숫자 뒤에는 숫자만 와야한다.
			break
		}

		if char == 32 {
			// 숫자 앞의 공백은 허용한다.
			continue
		}
	}

	if sign == 45 {
		result = -1 * result
	}

	return result
}
