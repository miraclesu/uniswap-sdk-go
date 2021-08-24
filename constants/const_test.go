package constants

import (
	"crypto/rand"
	"math/big"
	"testing"
)

// NOTE: Make sure that the RoundUp here is the largest constant
func randWholeNumber() int {
	max := big.NewInt(10)
	min := int(RoundUp + 1)
	i, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}

	return int(i.Int64()) + min
}

func randNegativeNumber() int {
	max := big.NewInt(10)
	i, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}

	return -1 - int(i.Int64())
}

func TestRounding_Valid(t *testing.T) {
	tests := []struct {
		name string
		r    Rounding
		want bool
	}{
		{"should return true if Rounding is RoundDown", RoundDown, true},
		{"should return true if Rounding is RoundHalfUp", RoundHalfUp, true},
		{"should return true if Rounding is RoundUp", RoundUp, true},
		{"should return true if Rounding is other whole numbers", Rounding(randWholeNumber()), false},
		{"should return true if Rounding is other negative numbers", Rounding(randNegativeNumber()), false},
	}
	for _, tt := range tests {
		r, want := tt.r, tt.want
		t.Run(tt.name, func(t *testing.T) {
			if got := r.Valid(); got != want {
				t.Errorf("Valid() = %v, want %v", got, want)
			}
		})
	}
}
