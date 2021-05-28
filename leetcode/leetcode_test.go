package leetcode_test

import (
	"testing"

	"github.com/whitesoil/leetcode/leetcode"
)

func TestAlgorithm(t *testing.T) {
	input := "()"
	output := leetcode.IsValid(input)

	if !output {
		t.Errorf("1. wrong: %t", output)
	}

	input = "()[]{}"
	output = leetcode.IsValid(input)

	if !output {
		t.Errorf("2. wrong: %t", output)
	}

}

func BenchmarkAlgorithm(b *testing.B) {
	for n := 0; n < b.N; n++ {

	}
}
