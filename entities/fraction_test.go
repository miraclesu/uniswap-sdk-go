package entities

import (
	"math/big"
	"testing"
)

func TestQuotient(t *testing.T) {
	var tests = []struct {
		Input  [2]int64
		Output int64
	}{
		{[2]int64{8, 3}, 2},  // one below
		{[2]int64{12, 4}, 3}, // exact
		{[2]int64{16, 5}, 3}, // one above
	}
	for i, test := range tests {
		output := NewFraction(big.NewInt(test.Input[0]), big.NewInt(test.Input[1])).Quotient()
		if output.Int64() != test.Output {
			t.Errorf("test #%d: failed to match when it should (%+v != %+v)", i, output, test.Output)
		}
	}
}

func TestRemainder(t *testing.T) {
	var tests = []struct {
		Input  [2]int64
		Output [2]int64
	}{
		{[2]int64{8, 3}, [2]int64{2, 3}},
		{[2]int64{12, 4}, [2]int64{0, 4}},
		{[2]int64{16, 5}, [2]int64{1, 5}},
	}
	for i, test := range tests {
		output := NewFraction(big.NewInt(test.Input[0]), big.NewInt(test.Input[1])).Remainder()
		expect := NewFraction(big.NewInt(test.Output[0]), big.NewInt(test.Output[1]))
		if !output.EqualTo(expect) {
			t.Errorf("test #%d: failed to match when it should (%+v != %+v)", i, output, expect)
		}
	}
}

func TestInvert(t *testing.T) {
	var tests = []struct {
		Input  [2]int64
		Output [2]int64
	}{
		{[2]int64{5, 10}, [2]int64{10, 5}},
	}
	for i, test := range tests {
		output := NewFraction(big.NewInt(test.Input[0]), big.NewInt(test.Input[1])).Invert()
		if output.Numerator.Int64() != test.Output[0] || output.Denominator.Int64() != test.Output[1] {
			t.Errorf("test #%d: failed to match when it should (%+v != %+v)", i, output, test.Output)
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		Input  [4]int64
		Output [2]int64
	}{
		{[4]int64{1, 10, 4, 12}, [2]int64{52, 120}},
		{[4]int64{1, 5, 2, 5}, [2]int64{3, 5}},
	}
	for i, test := range tests {
		output := NewFraction(big.NewInt(test.Input[0]), big.NewInt(test.Input[1])).Add(
			NewFraction(big.NewInt(test.Input[2]), big.NewInt(test.Input[3])),
		)
		expect := NewFraction(big.NewInt(test.Output[0]), big.NewInt(test.Output[1]))
		if !output.EqualTo(expect) {
			t.Errorf("test #%d: failed to match when it should (%+v != %+v)", i, output, expect)
		}
	}
}

func TestSubtract(t *testing.T) {
	var tests = []struct {
		Input  [4]int64
		Output [2]int64
	}{
		{[4]int64{1, 10, 4, 12}, [2]int64{-28, 120}},
		{[4]int64{3, 5, 2, 5}, [2]int64{1, 5}},
	}
	for i, test := range tests {
		output := NewFraction(big.NewInt(test.Input[0]), big.NewInt(test.Input[1])).Subtract(
			NewFraction(big.NewInt(test.Input[2]), big.NewInt(test.Input[3])),
		)
		expect := NewFraction(big.NewInt(test.Output[0]), big.NewInt(test.Output[1]))
		if !output.EqualTo(expect) {
			t.Errorf("test #%d: failed to match when it should (%+v != %+v)", i, output, expect)
		}
	}
}

func TestLessThan(t *testing.T) {
	var tests = []struct {
		Input  [4]int64
		Output bool
	}{
		{[4]int64{1, 10, 4, 12}, true},
		{[4]int64{1, 3, 4, 12}, false},
	}
	for i, test := range tests {
		output := NewFraction(big.NewInt(test.Input[0]), big.NewInt(test.Input[1])).LessThan(
			NewFraction(big.NewInt(test.Input[2]), big.NewInt(test.Input[3])),
		)
		if output != test.Output {
			t.Errorf("test #%d: failed to match when it should (%+v != %+v)", i, output, test.Output)
		}
	}
}

func TestEqualTo(t *testing.T) {
	var tests = []struct {
		Input  [4]int64
		Output bool
	}{
		{[4]int64{1, 10, 4, 12}, false},
		{[4]int64{1, 3, 4, 12}, true},
		{[4]int64{5, 12, 4, 12}, false},
	}
	for i, test := range tests {
		output := NewFraction(big.NewInt(test.Input[0]), big.NewInt(test.Input[1])).EqualTo(
			NewFraction(big.NewInt(test.Input[2]), big.NewInt(test.Input[3])),
		)
		if output != test.Output {
			t.Errorf("test #%d: failed to match when it should (%+v != %+v)", i, output, test.Output)
		}
	}
}

func TestGreaterThan(t *testing.T) {
	var tests = []struct {
		Input  [4]int64
		Output bool
	}{
		{[4]int64{1, 1, 2, 2}, false},
		{[4]int64{1, 10, 4, 12}, false},
		{[4]int64{1, 3, 4, 12}, false},
		{[4]int64{5, 12, 4, 12}, true},
	}
	for i, test := range tests {
		output := NewFraction(big.NewInt(test.Input[0]), big.NewInt(test.Input[1])).GreaterThan(
			NewFraction(big.NewInt(test.Input[2]), big.NewInt(test.Input[3])),
		)
		if output != test.Output {
			t.Errorf("test #%d: failed to match when it should (%+v != %+v)", i, output, test.Output)
		}
	}
}

func TestMultiply(t *testing.T) {
	var tests = []struct {
		Input  [4]int64
		Output [2]int64
	}{
		{[4]int64{1, 10, 4, 12}, [2]int64{4, 120}},
		{[4]int64{1, 3, 4, 12}, [2]int64{4, 36}},
		{[4]int64{5, 12, 4, 12}, [2]int64{20, 144}},
	}
	for i, test := range tests {
		output := NewFraction(big.NewInt(test.Input[0]), big.NewInt(test.Input[1])).Multiply(
			NewFraction(big.NewInt(test.Input[2]), big.NewInt(test.Input[3])),
		)
		expect := NewFraction(big.NewInt(test.Output[0]), big.NewInt(test.Output[1]))
		if !output.EqualTo(expect) {
			t.Errorf("test #%d: failed to match when it should (%+v != %+v)", i, output, expect)
		}
	}
}

func TestDivide(t *testing.T) {
	var tests = []struct {
		Input  [4]int64
		Output [2]int64
	}{
		{[4]int64{1, 10, 1, 12}, [2]int64{12, 10}},
		{[4]int64{1, 10, 4, 12}, [2]int64{12, 40}},
		{[4]int64{1, 3, 4, 12}, [2]int64{12, 12}},
		{[4]int64{5, 12, 4, 12}, [2]int64{60, 48}},
	}
	for i, test := range tests {
		output := NewFraction(big.NewInt(test.Input[0]), big.NewInt(test.Input[1])).Divide(
			NewFraction(big.NewInt(test.Input[2]), big.NewInt(test.Input[3])),
		)
		expect := NewFraction(big.NewInt(test.Output[0]), big.NewInt(test.Output[1]))
		if !output.EqualTo(expect) {
			t.Errorf("test #%d: failed to match when it should (%+v != %+v)", i, output, expect)
		}
	}
}

func TestToSignificant(t *testing.T) {
	var tests = []struct {
		Input  [2]int64
		Output string
		Format uint
	}{
		{[2]int64{30, 10}, "3", 0},
		{[2]int64{4, 10}, "0.4", 1},
		{[2]int64{126, 100}, "1.3", 1},
		{[2]int64{126, 100}, "1.26", 2},
		{[2]int64{124, 100}, "1.2", 1},
		{[2]int64{124, 100}, "1.24", 2},
	}
	for i, test := range tests {
		output := NewFraction(big.NewInt(test.Input[0]), big.NewInt(test.Input[1])).ToSignificant(test.Format)
		expect := test.Output
		if !(output == expect) {
			t.Errorf("test #%d: failed to match when it should (%+v != %+v)", i, output, expect)
		}
	}
}
