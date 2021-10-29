package leetcode_test

import (
	"testing"

	"github.com/whitesoil/leetcode/leetcode"
)

func TestAlgorithm(t *testing.T) {
	haystack := ""
	needle := ""
	prediction := 0
	output := leetcode.StrStr(haystack, needle)

	if output != prediction {
		t.Errorf("\nhaystack: %s\nneedle: %s\nprediction: %d\noutput: %d", haystack, needle, prediction, output)
	}

	haystack = "hello"
	needle = "ll"
	prediction = 2
	output = leetcode.StrStr(haystack, needle)

	if output != prediction {
		t.Errorf("\nhaystack: %s\nneedle: %s\nprediction: %d\noutput: %d", haystack, needle, prediction, output)
	}

	haystack = "aaa"
	needle = "bba"
	prediction = -1
	output = leetcode.StrStr(haystack, needle)

	if output != prediction {
		t.Errorf("\nhaystack: %s\nneedle: %s\nprediction: %d\noutput: %d", haystack, needle, prediction, output)
	}
}

func BenchmarkAlgorithm(b *testing.B) {
	for n := 0; n < b.N; n++ {

	}
}
