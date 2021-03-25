package entities

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/miraclesu/uniswap-sdk-go/constants"
	"github.com/miraclesu/uniswap-sdk-go/utils"
)

var (
	ErrDiffChainID = fmt.Errorf("diff chain id")
	ErrDiffToken   = fmt.Errorf("diff token")
	ErrSameAddrss  = fmt.Errorf("same address")

	_WETHCurrency, _ = newCurrency(constants.Decimals18, "WETH", "Wrapped Ether")

	WETH = map[constants.ChainID]*Token{
		constants.Mainnet: {
			Currency: _WETHCurrency,
			ChainID:  constants.Mainnet,
			Address:  utils.ValidateAndParseAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		},
		constants.Ropsten: {
			Currency: _WETHCurrency,
			ChainID:  constants.Ropsten,
			Address:  utils.ValidateAndParseAddress("0xc778417E063141139Fce010982780140Aa0cD5Ab"),
		},
		constants.Rinkeby: {
			Currency: _WETHCurrency,
			ChainID:  constants.Rinkeby,
			Address:  utils.ValidateAndParseAddress("0xc778417E063141139Fce010982780140Aa0cD5Ab"),
		},
		constants.Goerli: {
			Currency: _WETHCurrency,
			ChainID:  constants.Goerli,
			Address:  utils.ValidateAndParseAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"),
		},
		constants.Kovan: {
			Currency: _WETHCurrency,
			ChainID:  constants.Kovan,
			Address:  utils.ValidateAndParseAddress("0xd0A1E359811322d97991E03f863a0C30C2cF029C"),
		},
	}
)

/**
 * Represents an ERC20 token with a unique address and some metadata.
 */
type Token struct {
	*Currency

	constants.ChainID
	common.Address
}

func NewToken(chainID constants.ChainID, address common.Address, decimals int, symbol, name string) (*Token, error) {
	currency, err := newCurrency(decimals, symbol, name)
	if err != nil {
		return nil, err
	}

	return &Token{
		Currency: currency,
		ChainID:  chainID,
		Address:  address,
	}, nil
}

/**
 * Returns true if the two tokens are equivalent, i.e. have the same chainId and address.
 * @param other other token to compare
 */
func (t *Token) Equals(other *Token) bool {
	if t == other {
		return true
	}

	return t.ChainID == other.ChainID && t.Address == other.Address
}

/**
 * Returns true if the address of this token sorts before the address of the other token
 * @param other other token to compare
 * @throws if the tokens have the same address
 * @throws if the tokens are on different chains
 */
func (t *Token) SortsBefore(other *Token) (bool, error) {
	if t.ChainID != other.ChainID {
		return false, ErrDiffChainID
	}
	if t.Address == other.Address {
		return false, ErrSameAddrss
	}

	return strings.ToLower(t.Address.String()) < strings.ToLower(other.Address.String()), nil
}

// NewETHRToken creates a token that currency is ETH
func NewETHRToken(chainID constants.ChainID, address common.Address) *Token {
	return &Token{
		Currency: ETHER,
		ChainID:  chainID,
		Address:  address,
	}
}
