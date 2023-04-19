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

func TestVip(t *testing.T) {
	tests := []struct {
		i      float64
		p      float64
		result float64
	}{
		{i: 1, p: 1, result: 1},
		{i: 6, p: 3, result: 0.5},
		{i: 0.5, p: 8, result: 16},
	}

	for _, test := range tests {
		if test.result != Vip(test.i, test.p) {
			t.Errorf("unexpected test result: Vip(%f, %f) = %f", test.i, test.p, Vip(test.i, test.p))
		}
	}
}

func TestVser(t *testing.T) {
	tests := []struct {
		volts  []float64
		result float64
	}{
		{volts: []float64{2, 4, 6, 8}, result: 20},
		{volts: []float64{0, 0, 0, 0}, result: 0},
		{volts: []float64{0.2, 0.3, 2.0}, result: 2.5},
	}

	for _, test := range tests {
		if test.result != Vser(test.volts...) {
			t.Errorf("unexpected test result: Vser(%#v) = %f, expecting %f", test.volts, Vser(test.volts...), test.result)
		}
	}
}
