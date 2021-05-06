package leetcode

func IntToRoman(num int) string {
	arabian := num
	roman := ""

	for arabian > 0 {
		switch {
		case arabian >= 1000:
			quota := arabian / 1000
			for quota > 0 {
				roman = roman + "M"
				quota--
			}
			arabian = arabian % 1000
		case arabian >= 900:
			roman = roman + "CM"
			arabian = arabian - 900
		case arabian >= 500:
			roman = roman + "D"
			arabian = arabian - 500
		case arabian >= 400:
			roman = roman + "CD"
			arabian = arabian - 400
		case arabian >= 100:
			quota := arabian / 100
			for quota > 0 {
				roman = roman + "C"
				quota--
			}
			arabian = arabian % 100
		case arabian >= 90:
			roman = roman + "XC"
			arabian = arabian - 90
		case arabian >= 50:
			roman = roman + "L"
			arabian = arabian - 50
		case arabian >= 40:
			roman = roman + "XL"
			arabian = arabian - 40
		case arabian >= 10:
			quota := arabian / 10
			for quota > 0 {
				roman = roman + "X"
				quota--
			}
			arabian = arabian % 10
		case arabian == 9:
			roman = roman + "IX"
			arabian = 0
		case arabian >= 5:
			roman = roman + "V"
			arabian = arabian - 5
		case arabian == 4:
			roman = roman + "IV"
			arabian = 0
		case arabian >= 1:
			quota := arabian
			for quota > 0 {
				roman = roman + "I"
				quota--
			}
			arabian = 0
		}

	}

	return roman
}
