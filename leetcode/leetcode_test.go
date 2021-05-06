package leetcode_test

import (
	"testing"

	"github.com/whitesoil/leetcode/leetcode"
)

func TestAlgorithm(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := leetcode.MaxArea(height)

	if area != 49 {
		t.Error("Wrong: ", area)
	}

	height = []int{1, 1}
	area = leetcode.MaxArea(height)

	if area != 1 {
		t.Error("Wrong2: ", area)
	}

	height = []int{4, 3, 2, 1, 4}
	area = leetcode.MaxArea(height)

	if area != 16 {
		t.Error("Wrong3: ", area)
	}

	height = []int{1, 2, 1}
	area = leetcode.MaxArea(height)

	if area != 2 {
		t.Error("Wrong4: ", area)
	}

	height = []int{1, 2, 4, 3}
	area = leetcode.MaxArea(height)

	if area != 4 {
		t.Error("Wrong4: ", area)
	}

	height = []int{2, 3, 4, 5, 18, 17, 6}
	area = leetcode.MaxArea(height)

	if area != 17 {
		t.Error("Wrong4: ", area)
	}
}
func BenchmarkAlgorithm(b *testing.B) {
	for n := 0; n < b.N; n++ {
		height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		leetcode.MaxArea(height)
	}
}
