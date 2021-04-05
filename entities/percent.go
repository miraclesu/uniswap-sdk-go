package entities

import (
	"math/big"

	"github.com/miraclesu/uniswap-sdk-go/constants"
	"github.com/miraclesu/uniswap-sdk-go/number"
)

var (
	// Percent100 percent 100
	Percent100 = NewFraction(constants.B100, constants.One)
)

// Percent warps Fraction
type Percent struct {
	*Fraction
}

// NewPercent creates Percent
func NewPercent(num, deno *big.Int) *Percent {
	return &Percent{
		Fraction: NewFraction(num, deno),
	}
}

// ToSignificant format output
func (p *Percent) ToSignificant(significantDigits uint, opt ...number.Option) string {
	return p.Multiply(Percent100).ToSignificant(significantDigits, opt...)
}

// ToFixed format output
func (p *Percent) ToFixed(decimalPlaces uint, opt ...number.Option) string {
	return p.Multiply(Percent100).ToFixed(decimalPlaces, opt...)
}
