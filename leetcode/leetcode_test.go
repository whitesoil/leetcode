package leetcode_test

import (
	"testing"

	"github.com/whitesoil/leetcode/leetcode"
)

func TestAlgorithm(t *testing.T) {
	input := []string{"flower", "flow", "flight"}
	output := leetcode.LongestCommonPrefix(input)

	if output != "fl" {
		t.Error("wrong: " + output)
	}

	input = []string{"dog", "racecar", "car"}
	output = leetcode.LongestCommonPrefix(input)

	if output != "" {
		t.Error("wrong: " + output)
	}
}

func BenchmarkAlgorithm(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input := 1994
		leetcode.IntToRoman(input)
	}
}
