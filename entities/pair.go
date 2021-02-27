package entities

import (
	"math/big"
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
 * Returns the current mid price of the pair in terms of token0, i.e. the ratio of reserve1 to reserve0
 */
func (p *Pair) Token0Price() *Price {
	return NewPrice(p.Token0().Currency, p.Token1().Currency, p.TokenAmounts[0].Raw(), p.TokenAmounts[1].Raw())
}

/**
 * Returns the current mid price of the pair in terms of token1, i.e. the ratio of reserve0 to reserve1
 */

func (p *Pair) Token1Price() *Price {
	return NewPrice(p.Token1().Currency, p.Token0().Currency, p.TokenAmounts[1].Raw(), p.TokenAmounts[0].Raw())
}

/**
 * Return the price of the given token in terms of the other token in the pair.
 * @param token token to return price of
 */
func (p *Pair) PriceOf(token *Token) (*Price, error) {
	if !p.InvolvesToken(token) {
		return nil, ErrDiffToken
	}

	if token.Equals(p.Token0()) {
		return p.Token0Price(), nil
	}
	return p.Token1Price(), nil
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

func (p *Pair) ReserveOf(token *Token) (*TokenAmount, error) {
	if !p.InvolvesToken(token) {
		return nil, ErrDiffToken
	}

	if token.Equals(p.Token0()) {
		return p.Reserve0(), nil
	}
	return p.Reserve1(), nil
}

func (p *Pair) GetOutputAmount(inputAmount *TokenAmount) (*TokenAmount, *Pair, error) {
	if !p.InvolvesToken(inputAmount.Token) {
		return nil, nil, ErrDiffToken
	}

	inputReserve, outputReserve, token := p.Reserve0(), p.Reserve1(), p.Token1()
	if inputAmount.Token.Equals(p.Token1()) {
		inputReserve, outputReserve, token = p.Reserve1(), p.Reserve0(), p.Token0()
	}
	if inputReserve.Raw().Cmp(constants.Zero) == 0 || outputReserve.Raw().Cmp(constants.Zero) == 0 {
		return nil, nil, ErrInsufficientReserves
	}

	inputAmountWithFee := big.NewInt(1).Mul(inputAmount.Raw(), constants.B997)
	numerator := big.NewInt(1).Mul(inputAmountWithFee, outputReserve.Raw())
	denominator := big.NewInt(1).Add(big.NewInt(1).Mul(inputAmount.Raw(), constants.B1000), inputAmountWithFee)
	outputAmount, err := NewTokenAmount(token, big.NewInt(1).Div(numerator, denominator))
	if err != nil {
		return nil, nil, err
	}
	if outputAmount.Raw().Cmp(constants.Zero) == 0 {
		return nil, nil, ErrInsufficientInputAmount
	}

	tokenAmountA, err := inputAmount.Add(inputReserve)
	if err != nil {
		return nil, nil, err
	}
	tokenAmountB, err := outputReserve.Subtract(outputAmount)
	if err != nil {
		return nil, nil, err
	}
	pair, err := NewPair(tokenAmountA, tokenAmountB)
	if err != nil {
		return nil, nil, err
	}
	return outputAmount, pair, nil
}

func (p *Pair) GetInputAmount(outputAmount *TokenAmount) (*TokenAmount, *Pair, error) {
	if !p.InvolvesToken(outputAmount.Token) {
		return nil, nil, ErrDiffToken
	}

	outputReserve, inputReserve, token := p.Reserve0(), p.Reserve1(), p.Token0()
	if outputAmount.Token.Equals(p.Token0()) {
		outputReserve, inputReserve, token = p.Reserve1(), p.Reserve0(), p.Token1()
	}

	if inputReserve.Raw().Cmp(constants.Zero) == 0 || outputReserve.Raw().Cmp(constants.Zero) == 0 ||
		outputAmount.Raw().Cmp(inputReserve.Raw()) >= 0 {
		return nil, nil, ErrInsufficientReserves
	}

	numerator := big.NewInt(1).Mul(inputReserve.Raw(), outputAmount.Raw())
	numerator.Mul(numerator, constants.B1000)
	denominator := big.NewInt(1).Sub(outputReserve.Raw(), outputAmount.Raw())
	denominator.Mul(denominator, constants.B997)
	amount := big.NewInt(1).Div(numerator, denominator)
	amount.Add(amount, constants.One)
	inputAmount, err := NewTokenAmount(token, amount)
	if err != nil {
		return nil, nil, err
	}

	tokenAmountA, err := inputAmount.Add(inputReserve)
	if err != nil {
		return nil, nil, err
	}
	tokenAmountB, err := outputReserve.Subtract(outputAmount)
	if err != nil {
		return nil, nil, err
	}
	pair, err := NewPair(tokenAmountA, tokenAmountB)
	if err != nil {
		return nil, nil, err
	}
	return outputAmount, pair, nil
}
