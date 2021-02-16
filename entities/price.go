package entities

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"

	"github.com/miraclesu/uniswap-sdk-go/constants"
)

type Price struct {
	*Fraction
	BaseCurrency  *Currency // input i.e. denominator
	QuoteCurrency *Currency // output i.e. numerator
	Scalar        *Fraction // used to adjust the raw fraction w/r/t the decimals of the {base,quote}Token
}

func NewPriceFromRoute(route *Route) (*Price, error) {
	length := len(route.Pairs)
	// NOTE: check route Pairs len?
	prices := make([]*Price, length)
	for i := range route.Pairs {
		if route.Path[i].Equals(route.Pairs[i].Token0()) {
			prices[i] = NewPrice(route.Pairs[i].Reserve0().Currency, route.Pairs[i].Reserve1().Currency,
				route.Pairs[i].Reserve0().Raw(), route.Pairs[i].Reserve1().Raw())
		} else {
			prices[i] = NewPrice(route.Pairs[i].Reserve1().Currency, route.Pairs[i].Reserve0().Currency,
				route.Pairs[i].Reserve1().Raw(), route.Pairs[i].Reserve0().Raw())
		}
	}

	price := prices[0]
	for i := 1; i < length; i++ {
		if err := price.Multiply(prices[i]); err != nil {
			return nil, err
		}
	}
	return price, nil
}

// denominator and numerator _must_ be raw, i.e. in the native representation
func NewPrice(baseCurrency, quoteCurrency *Currency, denominator, numerator *big.Int) *Price {
	return &Price{
		Fraction:      NewFraction(numerator, denominator),
		BaseCurrency:  baseCurrency,
		QuoteCurrency: quoteCurrency,
		Scalar: NewFraction(math.BigPow(constants.Ten.Int64(), int64(baseCurrency.Decimals)),
			math.BigPow(constants.Ten.Int64(), int64(quoteCurrency.Decimals))),
	}
}

func (p *Price) Raw() *Fraction {
	return p.Fraction
}

func (p *Price) Adjusted() *Fraction {
	p.Fraction.Multiply(p.Scalar)
	return p.Fraction
}

func (p *Price) Invert() {
	p.BaseCurrency, p.QuoteCurrency = p.QuoteCurrency, p.BaseCurrency
}

func (p *Price) Multiply(other *Price) error {
	if !p.QuoteCurrency.Equals(other.BaseCurrency) {
		return ErrInvalidCurrency
	}

	p.Fraction.Multiply(other.Fraction)
	p.QuoteCurrency = other.QuoteCurrency
	return nil
}

// performs floor division on overflow
func (p *Price) Quote(currencyAmount *CurrencyAmount) {
	// TODO
}

func (p *Price) ToSignificant(significantDigits uint) string {
	return p.Adjusted().ToSignificant(significantDigits)
}

func (p *Price) ToFixed(decimalPlaces uint) string {
	return p.Adjusted().ToFixed(decimalPlaces)
}
