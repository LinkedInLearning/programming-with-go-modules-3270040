package circtests

import (
	"chapter2/ohm/current"
	"chapter2/ohm/power"
	"chapter2/ohm/resist"
	"chapter2/ohm/volt"
	"math"
	"testing"
)

// TestResistanceEq calculates resistance equival for resistors arranged in parallel
//
// +-----/\/\/-----+
// |       R1      |
// |               |
// +-----/\/\/-----+
// |       R2      |
// |               |
// +-----/\/\/-----+
// |       R3      |
// |               |
// +----- | |------+
//
//	Req = 1 / (1/R1 + 1/R2 + 1/R3)
func TestResistanceEq(t *testing.T) {
	tests := map[string]struct {
		r  []float64
		eq float64
	}{
		"one resistor":    {r: []float64{10}, eq: 10},
		"two resistors":   {r: []float64{10, 20}, eq: 6},
		"three resistors": {r: []float64{10, 20, 30}, eq: 5},
	}

	for _, test := range tests {
		req := resist.Rpara(test.r...)
		if math.Abs(req-test.eq) > 1.0 {
			t.Errorf("expecting resistance equivalance: %f, got: %f", test.eq, req)
		}
	}
}

// TestVolt finds voltage for simple circuit with resistance
// arranged in series:
//
// +--/\/\/-----/\/\/---/\/\/---+
// |   R1        R2       R3    |
// |                            |
// +--------- | |---------------+
//
//	V = IR
func TestVolt(t *testing.T) {
	tests := map[string]struct {
		i float64
		r []float64
		v float64
	}{
		"one resistor":    {i: 0.1, r: []float64{10}, v: 1},
		"two resistors":   {i: 0.1, r: []float64{10, 20}, v: 3},
		"three resistors": {i: 0.1, r: []float64{10, 20, 30}, v: 6},
	}

	for _, test := range tests {
		volt := volt.Vir(test.i, resist.Rser(test.r...))
		if math.Abs(volt-test.v) > 1.0 {
			t.Errorf("expecting voltage: %f, got: %f", test.v, volt)
		}
	}
}

// TestCurrent finds current for simple circuit with resistance
// arranged in parallel:
//
// +-----/\/\/-----+
// |       R1      |
// |               |
// +-----/\/\/-----+
// |       R2      |
// |               |
// +-----/\/\/-----+
// |       R3      |
// |               |
// +----- | |------+
//
//	I = V/R
func TestCurrent(t *testing.T) {
	tests := map[string]struct {
		i float64
		r []float64
		v float64
	}{
		"one resistor":    {i: 1.2, r: []float64{10}, v: 12},
		"two resistors":   {i: 1.8, r: []float64{10, 20}, v: 12},
		"three resistors": {i: 3.3, r: []float64{6, 12, 36}, v: 12},
	}

	for _, test := range tests {
		current := current.Ivr(test.v, resist.Rpara(test.r...))
		if math.Abs(current-test.i) > 1.0 {
			t.Errorf("expecting current: %f, got: %f", test.i, current)
		}
	}
}

// TestPower finds power value for simple circuit with voltage and resistance
// arranged in series:
//
// +--/\/\/-----/\/\/---/\/\/---+
// |   R1        R2       R3    |
// |                            |
// +--------- | |---------------+
//
//	Power = IV
func TestPowerRpara(t *testing.T) {
	tests := map[string]struct {
		i float64
		r []float64
		p float64
	}{
		"one resistor":    {i: 1.0, r: []float64{10}, p: 10},
		"two resistors":   {i: 1.5, r: []float64{10, 20}, p: 67},
		"three resistors": {i: 2.0, r: []float64{10, 20, 30}, p: 240},
	}

	for _, test := range tests {
		power := power.Piv(test.i, volt.Vir(test.i, resist.Rser(test.r...)))
		if math.Abs(power-test.p) > 1.0 {
			t.Errorf("expecting power: %f, got: %f", test.p, power)
		}
	}
}
