package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/whitesoil/leetcode/leetcode"
)

func TestAlgorithm(t *testing.T) {
	input := []int{1, 1}
	output := leetcode.RemoveDuplicates(input)

	fmt.Println(input)
	fmt.Println(output)

}

func BenchmarkAlgorithm(b *testing.B) {
	for n := 0; n < b.N; n++ {

	}
}
