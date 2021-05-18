package leetcode_test

import (
	"testing"

	"github.com/whitesoil/leetcode/leetcode"
)

func TestAlgorithm(t *testing.T) {
	input := []int{-1, 2, 1, -4}
	target := 1
	output := leetcode.ThreeSumClosest(input, target)

	if output != 2 {
		t.Errorf("1. wrong: %d", output)
	}

	input = []int{0, 2, 1, -3}
	target = 0
	output = leetcode.ThreeSumClosest(input, target)

	if output != 0 {
		t.Errorf("2. wrong: %d", output)
	}

	input = []int{1, 2, 4, 8, 16, 32, 64, 128}
	target = 82
	output = leetcode.ThreeSumClosest(input, target)

	if output != 82 {
		t.Errorf("3. wrong: %d", output)
	}
}

func BenchmarkAlgorithm(b *testing.B) {
	for n := 0; n < b.N; n++ {

	}
}
