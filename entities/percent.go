package entities

import "github.com/miraclesu/uniswap-sdk-go/constants"

var (
	Percent100 = NewFraction(constants.B100, constants.One)
)

type Percent struct {
	Fraction
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
