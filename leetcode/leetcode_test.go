package leetcode_test

import (
	"testing"

	"github.com/whitesoil/leetcode/leetcode"
)

func TestAlgorithm(t *testing.T) {
	input := []int{1, 0, -1, 0, -2, 2}
	target := 0
	output := leetcode.FourSum(input, target)

	if len(output) != 3 {
		t.Errorf("1. wrong: %d", len(output))
	}

	input = []int{2, 2, 2, 2, 2}
	target = 8
	output = leetcode.FourSum(input, target)

	if len(output) != 1 {
		t.Errorf("2. wrong: %d", len(output))
	}
}

func BenchmarkAlgorithm(b *testing.B) {
	for n := 0; n < b.N; n++ {

	}
}
