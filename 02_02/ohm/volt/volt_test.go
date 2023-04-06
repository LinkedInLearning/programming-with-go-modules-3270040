package volt

import (
	"testing"
)

func TestVir(t *testing.T) {
	tests := []struct {
		i      float64
		r      float64
		result float64
	}{
		{i: 0, r: 0, result: 0},
		{i: 2, r: 2, result: 4},
		{i: 0.5, r: 8, result: 4},
	}

	for _, test := range tests {
		if test.result != Vir(test.i, test.r) {
			t.Errorf("unexpected test result")
		}
	}
}
