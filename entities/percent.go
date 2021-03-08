package entities

import (
	"math/big"

	"github.com/miraclesu/uniswap-sdk-go/constants"
)

var (
	Percent100 = NewFraction(constants.B100, constants.One)
)

type Percent struct {
	*Fraction
}

func NewPercent(num, deno *big.Int) *Percent {
	return &Percent{
		Fraction: NewFraction(num, deno),
	}
}

// NOTE: format, rounding
// TODO
func (p *Percent) ToSignificant(significantDigits uint) string {
	p.Multiply(Percent100)
	return p.Fraction.ToSignificant(significantDigits)
}

// TODO
func (p *Percent) ToFixed(decimalPlaces uint) string {
	p.Multiply(Percent100)
	return p.Fraction.ToFixed(decimalPlaces)
}
