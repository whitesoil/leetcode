package leetcode_test

import (
	"strconv"
	"testing"

	"github.com/whitesoil/leetcode/leetcode"
)

func TestAlgorithm(t *testing.T) {
	input := "III"
	output := leetcode.RomanToInt(input)

	if output != 3 {
		t.Error("wrong: " + strconv.Itoa(output))
	}

	input = "IV"
	output = leetcode.RomanToInt(input)

	if output != 4 {
		t.Error("wrong: " + strconv.Itoa(output))
	}

	input = "IX"
	output = leetcode.RomanToInt(input)

	if output != 9 {
		t.Error("wrong: " + strconv.Itoa(output))
	}

	input = "LVIII"
	output = leetcode.RomanToInt(input)

	if output != 58 {
		t.Error("wrong: " + strconv.Itoa(output))
	}

	input = "MCMXCIV"
	output = leetcode.RomanToInt(input)

	if output != 1994 {
		t.Error("wrong: " + strconv.Itoa(output))
	}
}

func BenchmarkAlgorithm(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input := 1994
		leetcode.IntToRoman(input)
	}
}
