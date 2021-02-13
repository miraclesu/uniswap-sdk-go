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

	pair.LiquidityToken, err = NewToken(tokenAmountA.Token.ChainID, address, constants.Decimals18, constants.Univ2Symbol, constants.Univ2Name)
	return pair, err
}

func (p *Pair) GetAddress() (string, error) {
	return _PairAddressCache.GetAddress(p.TokenAmounts[0].Token.Address.String(), p.TokenAmounts[1].Token.Address.String())
}

/**
 * Returns true if the token is either token0 or token1
 * @param token to check
 */
func (p *Pair) InvolvesToken(token *Token) bool {
	return token.Equals(p.TokenAmounts[0].Token) || token.Equals(p.TokenAmounts[1].Token)
}

/**
 * Returns the chain ID of the tokens in the pair.
 */
func (p *Pair) ChainID() constants.ChainID {
	return p.Token0().ChainID
}

func (p *Pair) Token0() *Token {
	return p.TokenAmounts[0].Token
}

func (p *Pair) Token1() *Token {
	return p.TokenAmounts[1].Token
}

func (p *Pair) Reserve0() *TokenAmount {
	return p.TokenAmounts[0]
}

func (p *Pair) Reserve1() *TokenAmount {
	return p.TokenAmounts[1]
}
