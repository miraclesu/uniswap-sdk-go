package entities

import (
	"fmt"

	"github.com/miraclesu/uniswap-sdk-go/constants"
)

var (
	ErrInvalidSlippageTolerance = fmt.Errorf("invalid slippage tolerance")
)

/**
 * Represents a trade executed against a list of pairs.
 * Does not account for slippage, i.e. trades that front run this trade and move the price.
 */
type Trade struct {
	/**
	 * The route of the trade, i.e. which pairs the trade goes through.
	 */
	route *Route
	/**
	 * The type of the trade, either exact in or exact out.
	 */
	tradeType constants.TradeType
	/**
	 * The input amount for the trade assuming no slippage.
	 */
	inputAmount *TokenAmount
	/**
	 * The output amount for the trade assuming no slippage.
	 */
	outputAmount *TokenAmount
	/**
	 * The price expressed in terms of output amount/input amount.
	 */
	executionPrice *Price
	/**
	 * The mid price after the trade executes assuming no slippage.
	 */
	nextMidPrice *Price
	/**
	 * The percent difference between the mid price before the trade and the trade execution price.
	 */
	priceImpact *Percent
}

func (t *Trade) InputAmount() *TokenAmount {
	return t.inputAmount
}

func (t *Trade) OutputAmount() *TokenAmount {
	return t.outputAmount
}

/**
 * Constructs an exact in trade with the given amount in and route
 * @param route route of the exact in trade
 * @param amountIn the amount being passed in
 */
func ExactIn(route *Route, amountIn *TokenAmount) (*Trade, error) {
	return NewTrade(route, amountIn, constants.ExactInput)
}

/**
 * Constructs an exact out trade with the given amount out and route
 * @param route route of the exact out trade
 * @param amountOut the amount returned by the trade
 */
func ExactOut(route *Route, amountOut *TokenAmount) (*Trade, error) {
	return NewTrade(route, amountOut, constants.ExactOutput)
}

// NewTrade creates a new trade
// nolint gocyclo
func NewTrade(route *Route, amount *TokenAmount, tradeType constants.TradeType) (*Trade, error) {
	amounts := make([]*TokenAmount, len(route.Path))
	nextPairs := make([]*Pair, len(route.Pairs))

	if tradeType == constants.ExactInput {
		if !route.Input.Currency.Equals(amount.Token.Currency) {
			return nil, ErrInvalidCurrency
		}
		if !route.Input.Equals(amount.Token) {
			return nil, ErrDiffToken
		}

		amounts[0] = amount
		for i := 0; i < len(route.Path)-1; i++ {
			outputAmount, nextPair, err := route.Pairs[i].GetOutputAmount(amounts[i])
			if err != nil {
				return nil, err
			}
			amounts[i+1] = outputAmount
			nextPairs[i] = nextPair
		}
	} else {
		if !route.Output.Currency.Equals(amount.Token.Currency) {
			return nil, ErrInvalidCurrency
		}
		if !route.Output.Equals(amount.Token) {
			return nil, ErrDiffToken
		}

		amounts[len(amounts)-1] = amount
		for i := len(route.Path) - 1; i > 0; i-- {
			inputAmount, nextPair, err := route.Pairs[i-1].GetInputAmount(amounts[i])
			if err != nil {
				return nil, err
			}
			amounts[i-1] = inputAmount
			nextPairs[i-1] = nextPair
		}
	}

	route, err := NewRoute(nextPairs, route.Input, nil)
	if err != nil {
		return nil, err
	}
	nextMidPrice, err := NewPriceFromRoute(route)
	if err != nil {
		return nil, err
	}
	inputAmount := amount
	if tradeType == constants.ExactOutput {
		inputAmount = amounts[0]
		if route.Input.Currency.Equals(ETHER) {
			inputAmount = wrappedAmount(inputAmount.CurrencyAmount, route.ChainID())
		}
	}
	outputAmount := amount
	if tradeType == constants.ExactInput {
		outputAmount = amounts[len(amounts)-1]
		if route.Output.Currency.Equals(ETHER) {
			outputAmount = wrappedAmount(outputAmount.CurrencyAmount, route.ChainID())
		}
	}
	price := NewPrice(inputAmount.Currency, outputAmount.Currency, inputAmount.Raw(), outputAmount.Raw())
	return &Trade{
		route:          route,
		tradeType:      tradeType,
		inputAmount:    inputAmount,
		outputAmount:   outputAmount,
		executionPrice: price,
		nextMidPrice:   nextMidPrice,
		priceImpact:    computePriceImpact(route.MidPrice, inputAmount, outputAmount),
	}, nil
}

/**
 * Given a currency amount and a chain ID, returns the equivalent representation as the token amount.
 * In other words, if the currency is ETHER, returns the WETH token amount for the given chain. Otherwise, returns
 * the input currency amount.
 */
func wrappedAmount(currencyAmount *CurrencyAmount, chainID constants.ChainID) *TokenAmount {
	// TODO: returns an error maybe better
	if !currencyAmount.Currency.Equals(ETHER) {
		panic(ErrInvalidCurrency)
	}

	token := WETH[chainID]
	ca, _ := NewCurrencyAmount(token.Currency, currencyAmount.Raw())
	return &TokenAmount{
		CurrencyAmount: ca,
		Token:          token,
	}
}

/**
 * Returns the percent difference between the mid price and the execution price, i.e. price impact.
 * @param midPrice mid price before the trade
 * @param inputAmount the input amount of the trade
 * @param outputAmount the output amount of the trade
 */
func computePriceImpact(midPrice *Price, inputAmount, outputAmount *TokenAmount) *Percent {
	exactQuote := midPrice.Raw().Multiply(NewFraction(inputAmount.Raw(), nil))
	slippage := exactQuote.Subtract(NewFraction(outputAmount.Raw(), nil)).Divide(exactQuote)
	return &Percent{
		Fraction: slippage,
	}
}

/**
 * Get the minimum amount that must be received from this trade for the given slippage tolerance
 * @param slippageTolerance tolerance of unfavorable slippage from the execution price of this trade
 */
func (t *Trade) MinimumAmountOut(slippageTolerance *Percent) (*TokenAmount, error) {
	if slippageTolerance.LessThan(ZeroFraction) {
		return nil, ErrInvalidSlippageTolerance
	}

	if t.tradeType == constants.ExactOutput {
		return t.outputAmount, nil
	}

	slippageAdjustedAmountOut := NewFraction(constants.One, nil).
		Add(slippageTolerance.Fraction).
		Invert().
		Multiply(NewFraction(t.outputAmount.Raw(), nil)).Quotient()
	return NewTokenAmount(t.outputAmount.Token, slippageAdjustedAmountOut)
}

/**
 * Get the maximum amount in that can be spent via this trade for the given slippage tolerance
 * @param slippageTolerance tolerance of unfavorable slippage from the execution price of this trade
 */
func (t *Trade) MaximumAmountIn(slippageTolerance *Percent) (*TokenAmount, error) {
	if slippageTolerance.LessThan(ZeroFraction) {
		return nil, ErrInvalidSlippageTolerance
	}

	if t.tradeType == constants.ExactInput {
		return t.inputAmount, nil
	}

	slippageAdjustedAmountIn := NewFraction(constants.One, nil).
		Add(slippageTolerance.Fraction).
		Multiply(NewFraction(t.inputAmount.Raw(), nil)).Quotient()
	return NewTokenAmount(t.inputAmount.Token, slippageAdjustedAmountIn)
}
