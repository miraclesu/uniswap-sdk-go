package entities

import (
	"sync"

	"github.com/miraclesu/uniswap-sdk-go/constants"
)

var (
	_PairAddressCache = &PairAddressCache{
		lk:      new(sync.RWMutex),
		address: make(map[string]map[string]string, 16),
	}
)

type TokenAmounts [2]*TokenAmount

type Tokens [2]*Token

func NewTokenAmounts(tokenAmountA, tokenAmountB *TokenAmount) (TokenAmounts, error) {
	ok, err := tokenAmountA.Token.SortsBefore(tokenAmountB.Token)
	if err != nil {
		return TokenAmounts{}, err
	}
	if ok {
		return TokenAmounts{tokenAmountA, tokenAmountB}, nil
	}
	return TokenAmounts{tokenAmountB, tokenAmountA}, nil
}

type PairAddressCache struct {
	lk *sync.RWMutex
	// token0 address : token1 address : pair address
	address map[string]map[string]string
}

// TODO
func (p *PairAddressCache) GetAddress(addressA, addressB string) (string, error) {
	return "", nil
}

type Pair struct {
	LiquidityToken *Token
	// sorted tokens
	TokenAmounts
}

func NewPair(tokenAmountA, tokenAmountB *TokenAmount) (*Pair, error) {
	tokenAmounts, err := NewTokenAmounts(tokenAmountA, tokenAmountB)
	if err != nil {
		return nil, err
	}

	pair := &Pair{
		TokenAmounts: tokenAmounts,
	}
	address, err := pair.GetAddress()
	if err != nil {
		return nil, err
	}

	pair.LiquidityToken, err = NewToken(tokenAmountA.ChainID, address, constants.Decimals18, constants.Univ2Symbol, constants.Univ2Name)
	return pair, err
}

func (p *Pair) GetAddress() (string, error) {
	return _PairAddressCache.GetAddress(p.TokenAmounts[0].Address.String(), p.TokenAmounts[1].Address.String())
}
