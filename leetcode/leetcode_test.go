package leetcode_test

import (
	"testing"

	"github.com/whitesoil/leetcode/leetcode"
)

func TestAlgorithm(t *testing.T) {
	input := 3
	output := leetcode.IntToRoman(input)

	if output != "III" {
		t.Error("Wrong: ", output)
	}

	input = 4
	output = leetcode.IntToRoman(input)

	if output != "IV" {
		t.Error("Wrong: ", output)
	}

	input = 9
	output = leetcode.IntToRoman(input)

	if output != "IX" {
		t.Error("Wrong: ", output)
	}

	input = 58
	output = leetcode.IntToRoman(input)

	if output != "LVIII" {
		t.Error("Wrong: ", output)
	}

	input = 1994
	output = leetcode.IntToRoman(input)

	if output != "MCMXCIV" {
		t.Error("Wrong: ", output)
	}
}
func BenchmarkAlgorithm(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input := 1994
		leetcode.IntToRoman(input)
	}
}
