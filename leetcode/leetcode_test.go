package leetcode_test

import (
	"testing"

	"github.com/whitesoil/leetcode/leetcode"
)

func TestAlgorithm(t *testing.T) {
	node4 := &leetcode.ListNode{4, nil}
	node3 := &leetcode.ListNode{3, node4}
	node2 := &leetcode.ListNode{2, node3}
	node1 := &leetcode.ListNode{1, node2}
	leetcode.SwapPairs(node1)

}

func BenchmarkAlgorithm(b *testing.B) {
	for n := 0; n < b.N; n++ {

	}
}
