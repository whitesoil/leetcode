package leetcode_test

import (
	"testing"

	"github.com/whitesoil/leetcode/leetcode"
)

func TestAlgorithm(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	output := leetcode.ThreeSum(input)

	if len(output) != 2 {
		t.Errorf("Wrong: %d", len(output))
	}

	input = []int{-2, 0, 0, 2, 2}
	output = leetcode.ThreeSum(input)

	if len(output) != 1 {
		t.Errorf("Wrong: %d", len(output))
	}
}

func BenchmarkAlgorithm(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input := 1994
		leetcode.IntToRoman(input)
	}
}
