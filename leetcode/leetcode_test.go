package leetcode_test

import (
	"testing"

	"github.com/whitesoil/leetcode/leetcode"
)

func TestAlgorithm(t *testing.T) {
	node5 := &leetcode.ListNode{5, nil}
	node4 := &leetcode.ListNode{4, node5}
	node3 := &leetcode.ListNode{3, node4}
	node2 := &leetcode.ListNode{2, node3}
	node1 := &leetcode.ListNode{1, node2}
	leetcode.ReverseKGroup(node1, 3)

}

func BenchmarkAlgorithm(b *testing.B) {
	for n := 0; n < b.N; n++ {

	}
}
