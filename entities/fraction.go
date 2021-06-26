package entities

import (
	"math/big"

	"github.com/shopspring/decimal"

	"github.com/miraclesu/uniswap-sdk-go/constants"
	"github.com/miraclesu/uniswap-sdk-go/number"
)

var (
	// ZeroFraction zero fraction instance
	ZeroFraction = NewFraction(constants.Zero, nil)
)

// Fraction warps math franction
type Fraction struct {
	Numerator   *big.Int
	Denominator *big.Int

	opts *number.Options
}

// NewFraction creates a fraction
func NewFraction(num, deno *big.Int) *Fraction {
	if deno == nil {
		deno = constants.One
	}
	return &Fraction{
		Numerator:   num,
		Denominator: deno,
	}
}

// Quotient performs floor division
func (f *Fraction) Quotient() *big.Int {
	z := new(big.Int)
	return z.Div(f.Numerator, f.Denominator)
}

// Remainder remainder after floor division
func (f *Fraction) Remainder() *Fraction {
	z := new(big.Int)
	return NewFraction(z.Rem(f.Numerator, f.Denominator), f.Denominator)
}

// Invert inverts a fraction
func (f *Fraction) Invert() *Fraction {
	return NewFraction(f.Denominator, f.Numerator)
}

// Add adds two fraction and returns a new fraction
// nolint dupl
func (f *Fraction) Add(other *Fraction) *Fraction {
	if f.Denominator.Cmp(other.Denominator) == 0 {
		return NewFraction(big.NewInt(0).Add(f.Numerator, other.Numerator), f.Denominator)
	}

	return NewFraction(
		big.NewInt(0).Add(
			big.NewInt(0).Mul(f.Numerator, other.Denominator),
			big.NewInt(0).Mul(other.Numerator, f.Denominator),
		),
		big.NewInt(0).Mul(f.Denominator, other.Denominator),
	)
}

// Add subtracts two fraction and returns a new fraction
// nolint dupl
func (f *Fraction) Subtract(other *Fraction) *Fraction {
	if f.Denominator.Cmp(other.Denominator) == 0 {
		return NewFraction(big.NewInt(0).Sub(f.Numerator, other.Numerator), f.Denominator)
	}

	return NewFraction(
		big.NewInt(0).Sub(
			big.NewInt(0).Mul(f.Numerator, other.Denominator),
			big.NewInt(0).Mul(other.Numerator, f.Denominator),
		),
		big.NewInt(0).Mul(f.Denominator, other.Denominator),
	)
}

// LessThan identifies whether the caller is less than the other
func (f *Fraction) LessThan(other *Fraction) bool {
	return big.NewInt(0).Mul(f.Numerator, other.Denominator).
		Cmp(big.NewInt(0).Mul(other.Numerator, f.Denominator)) < 0
}

// EqualTo identifies whether the caller is equal to the other
func (f *Fraction) EqualTo(other *Fraction) bool {
	return big.NewInt(0).Mul(f.Numerator, other.Denominator).
		Cmp(big.NewInt(0).Mul(other.Numerator, f.Denominator)) == 0
}

// GreaterThan identifies whether the caller is greater than the other
func (f *Fraction) GreaterThan(other *Fraction) bool {
	return big.NewInt(0).Mul(f.Numerator, other.Denominator).
		Cmp(big.NewInt(0).Mul(other.Numerator, f.Denominator)) > 0
}

// Multiply mul two fraction and returns a new fraction
func (f *Fraction) Multiply(other *Fraction) *Fraction {
	return NewFraction(
		big.NewInt(0).Mul(f.Numerator, other.Numerator),
		big.NewInt(0).Mul(f.Denominator, other.Denominator),
	)
}

// Divide mul div two fraction and returns a new fraction
func (f *Fraction) Divide(other *Fraction) *Fraction {
	return NewFraction(
		big.NewInt(0).Mul(f.Numerator, other.Denominator),
		big.NewInt(0).Mul(f.Denominator, other.Numerator),
	)
}

// ToSignificant format output
func (f *Fraction) ToSignificant(significantDigits uint, opt ...number.Option) string {
	f.opts = number.New(number.WithGroupSeparator('\xA0'), number.WithRoundingMode(constants.RoundHalfUp))
	f.opts.Apply(opt...)
	f.opts.Apply(number.WithRoundingPrecision(int(significantDigits)))

	d := decimal.NewFromBigInt(f.Numerator, 0).Div(decimal.NewFromBigInt(f.Denominator, 0))
	if v, err := number.DecimalRound(d, f.opts); err == nil {
		d = v
	}

	return number.DecimalFormat(d, f.opts)
}

// ToFixed format output
func (f *Fraction) ToFixed(decimalPlaces uint, opt ...number.Option) string {
	f.opts = number.New(number.WithGroupSeparator('\xA0'), number.WithRoundingMode(constants.RoundHalfUp))
	f.opts.Apply(opt...)
	f.opts.Apply(number.WithDecimalPlaces(decimalPlaces))

	d := decimal.NewFromBigInt(big.NewInt(0).Div(f.Numerator, f.Denominator), 0)

	return number.DecimalFormat(d, f.opts)
}
