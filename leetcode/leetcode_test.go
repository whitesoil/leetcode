package leetcode_test

import (
	"testing"

	"github.com/whitesoil/leetcode/leetcode"
)

func TestAlgorithm(t *testing.T) {
	input := "23"
	output := leetcode.LetterCombinations(input)

	if len(output) != 9 {
		t.Errorf("1. wrong: %d", len(output))
	}

	input = ""
	output = leetcode.LetterCombinations(input)

	if len(output) != 0 {
		t.Errorf("2. wrong: %d", len(output))
	}

	input = "2"
	output = leetcode.LetterCombinations(input)

	if len(output) != 3 {
		t.Errorf("3. wrong: %d", len(output))
	}
}

func BenchmarkAlgorithm(b *testing.B) {
	for n := 0; n < b.N; n++ {

	}
}
