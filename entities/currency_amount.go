package entities

import (
	"errors"
	"math/big"

	"github.com/miraclesu/uniswap-sdk-go/constants"
	"github.com/miraclesu/uniswap-sdk-go/utils"
)

var (
	ErrInsufficientReserves    = errors.New("doesn't have insufficient reserves")
	ErrInsufficientInputAmount = errors.New("the input amount insufficient reserves")
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

/**
 * Helper that calls the constructor with the ETHER currency
 * @param amount ether amount in wei
 */
func NewEther(amount *big.Int) (*CurrencyAmount, error) {
	return NewCurrencyAmount(ETHER, amount)
}
