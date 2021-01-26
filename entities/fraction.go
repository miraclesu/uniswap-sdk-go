package entities

import (
	"math/big"

	"github.com/miraclesu/uniswap-sdk-go/constants"
)

type Fraction struct {
	Numerator   *big.Int
	Denominator *big.Int
}

func NewFraction(num, deno *big.Int) *Fraction {
	if deno == nil {
		deno = constants.One
	}
	return &Fraction{
		Numerator:   num,
		Denominator: deno,
	}
}

// performs floor division
func (f *Fraction) Quotient() *big.Int {
	z := new(big.Int)
	return z.Div(f.Numerator, f.Denominator)
}

// remainder after floor division
func (f *Fraction) Remainder() *Fraction {
	z := new(big.Int)
	return NewFraction(z.Rem(f.Numerator, f.Denominator), f.Denominator)
}

func (f *Fraction) Invert() *Fraction {
	return NewFraction(f.Denominator, f.Numerator)
}
