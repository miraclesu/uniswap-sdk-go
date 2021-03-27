package entities

import "math/big"

type TokenAmount struct {
	*CurrencyAmount
	Token *Token
}

// amount _must_ be raw, i.e. in the native representation
func NewTokenAmount(token *Token, amount *big.Int) (*TokenAmount, error) {
	currencyAmount, err := NewCurrencyAmount(token.Currency, amount)
	if err != nil {
		return nil, err
	}

	return &TokenAmount{
		Token:          token,
		CurrencyAmount: currencyAmount,
	}, nil
}

func (t *TokenAmount) Add(other *TokenAmount) (*TokenAmount, error) {
	if !t.Token.Equals(other.Token) {
		return nil, ErrDiffToken
	}

	return NewTokenAmount(t.Token, big.NewInt(0).Add(t.Raw(), other.Raw()))
}

func (t *TokenAmount) Subtract(other *TokenAmount) (*TokenAmount, error) {
	if !t.Token.Equals(other.Token) {
		return nil, ErrDiffToken
	}

	return NewTokenAmount(t.Token, big.NewInt(0).Sub(t.Raw(), other.Raw()))
}

func (t *TokenAmount) Equals(other *TokenAmount) bool {
	return t.Token.Equals(other.Token) && t.Fraction.EqualTo(other.Fraction)
}
