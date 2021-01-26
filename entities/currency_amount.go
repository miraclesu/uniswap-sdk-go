package entities

import (
	"math/big"

	"github.com/miraclesu/uniswap-sdk-go/constants"
	"github.com/miraclesu/uniswap-sdk-go/utils"
)

type CurrencyAmount struct {
	*Fraction
	*Currency
}

// amount _must_ be raw, i.e. in the native representation
func NewCurrencyAmount(currency *Currency, amount *big.Int) (*CurrencyAmount, error) {
	if err := utils.ValidateSolidityTypeInstance(amount, constants.Uint256); err != nil {
		return nil, err
	}

	fraction := NewFraction(amount, big.NewInt(0).Exp(constants.Ten, big.NewInt(int64(currency.Decimals)), nil))
	return &CurrencyAmount{
		Fraction: fraction,
		Currency: currency,
	}, nil
}

func (c *CurrencyAmount) Raw() *big.Int {
	return c.Numerator
}
