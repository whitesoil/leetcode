package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/whitesoil/leetcode/leetcode"
)

func TestAlgorithm(t *testing.T) {
	input := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	output := leetcode.RemoveElement(input, val)

	fmt.Println(input)
	fmt.Println(output)

}

func BenchmarkAlgorithm(b *testing.B) {
	for n := 0; n < b.N; n++ {

	}
}
