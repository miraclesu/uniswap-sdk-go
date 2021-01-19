package entities

import (
	"math/big"

	"github.com/miraclesu/uniswap-sdk-go/constants"
	"github.com/miraclesu/uniswap-sdk-go/utils"
)

/**
 * The only instance of the base class `Currency`.
 */
var (
	ETHER, _ = newCurrency(18, "ETH", "Ether")
)

/**
 * A currency is any fungible financial instrument on Ethereum, including Ether and all ERC20 tokens.
 */
type Currency struct {
	Decimals int
	Symbol   string
	Name     string
}

/**
 * newCurrency an instance of the base class `Currency`. The only instance of the base class `Currency` is `Currency.ETHER`.
 * @param decimals decimals of the currency
 * @param symbol symbol of the currency
 * @param name of the currency
 */
func newCurrency(decimals int, symbol, name string) (*Currency, error) {
	if err := utils.ValidateSolidityTypeInstance(big.NewInt(int64(decimals)), constants.Uint8); err != nil {
		return nil, err
	}

	return &Currency{
		Decimals: decimals,
		Symbol:   symbol,
		Name:     name,
	}, nil
}
