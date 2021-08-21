package entities

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"

	"github.com/miraclesu/uniswap-sdk-go/constants"
	"github.com/miraclesu/uniswap-sdk-go/number"
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
	var err error
	for i := 1; i < length; i++ {
		price, err = price.Multiply(prices[i])
		if err != nil {
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
	return NewFraction(p.Fraction.Numerator, p.Fraction.Denominator)
}

func (p *Price) Adjusted() *Fraction {
	return p.Fraction.Multiply(p.Scalar)
}

func (p *Price) Invert() *Price {
	return NewPrice(p.QuoteCurrency, p.BaseCurrency, p.Numerator, p.Denominator)
}

func (p *Price) Multiply(other *Price) (*Price, error) {
	if !p.QuoteCurrency.Equals(other.BaseCurrency) {
		return nil, ErrInvalidCurrency
	}

	fraction := p.Fraction.Multiply(other.Fraction)
	return NewPrice(p.BaseCurrency, other.QuoteCurrency, fraction.Denominator, fraction.Numerator), nil
}

// performs floor division on overflow
func (p *Price) Quote(currencyAmount *CurrencyAmount) (*CurrencyAmount, error) {
	if !p.BaseCurrency.Equals(currencyAmount.Currency) {
		return nil, ErrInvalidCurrency
	}

	return NewEther(p.Fraction.Multiply(NewFraction(currencyAmount.Raw(), nil)).Quotient())
}

func (p *Price) ToSignificant(significantDigits uint, opt ...number.Option) string {
	return p.Adjusted().ToSignificant(significantDigits, opt...)
}

func (p *Price) ToFixed(decimalPlaces uint, opt ...number.Option) string {
	return p.Adjusted().ToFixed(decimalPlaces, opt...)
}
