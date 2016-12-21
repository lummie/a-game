package scene

import "testing"

func TestRoundInt(t *testing.T) {
	type intTest struct {
		V  float64
		Ex int
	}

	tests := []intTest{
		{0.1, 0},
		{0.2, 0},
		{0.4, 0},
		{0.5, 1},
		{0.7, 1},
		{1.7, 2},
		{1.1, 1},
	}

	for _, test := range tests {
		if RoundInt(test.V) != test.Ex {
			t.Errorf("Expected %v got %v", test.Ex, RoundInt(test.V))
		}
	}
}

func TestRoundUp(t *testing.T) {
	type floatTest struct {
		V  float64
		Dp int
		Ex float64
	}
	tests := []floatTest{
		{0.1, 1, 0.1},
		{0.04, 1, 0.0},
		{0.06, 1, 0.1},
		{0.5, 0, 1.0},
		{0.7, 1, 0.7},
		{1.1, 1, 1.1},
		{1.04, 1, 1.0},
		{1.05, 1, 1.1},
		{1.5, 0, 2.0},
		{2.7, 0, 3.0},
	}

	for testIndex, test := range tests {
		if RoundUp(test.V, test.Dp) != test.Ex {
			t.Errorf("%v: Expected %v got %v", testIndex, test.Ex, RoundUp(test.V, test.Dp))
		}
	}
}