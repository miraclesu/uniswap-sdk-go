package entities

import (
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"

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

func (f *Fraction) Multiply(other *Fraction) {
	f.Numerator.Mul(f.Numerator, other.Numerator)
	f.Denominator.Mul(f.Denominator, other.Denominator)
}

// NOTE: format, rounding
// TODO
func (f *Fraction) ToSignificant(significantDigits uint) string {
	d := decimal.NewFromBigInt(big.NewInt(0).Div(f.Numerator, f.Denominator), 0)
	return fmt.Sprintf("%v", d)
}

// TODO
func (f *Fraction) ToFixed(decimalPlaces uint) string {
	d := big.NewInt(0).Div(f.Numerator, f.Denominator)
	return fmt.Sprintf("%v", d)
}
