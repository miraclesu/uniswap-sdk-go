package entities

import "github.com/miraclesu/uniswap-sdk-go/constants"

type BestTradeOptions struct {
	// how many results to return
	MaxNumResults int
	// the maximum number of hops a trade should contain
	MaxHops int
}

func (o *BestTradeOptions) ReduceHops() {
	o.MaxHops--
}

// minimal interface so the input output comparator may be shared across types
type InputOutput interface {
	InputAmount() *TokenAmount
	OutputAmount() *TokenAmount
}

// comparator function that allows sorting trades by their output amounts, in decreasing order, and then input amounts
// in increasing order. i.e. the best trades have the most outputs for the least inputs and are sorted first
func InputOutputComparator(a, b InputOutput) int {
	// must have same input and output token for comparison
	if !a.InputAmount().Currency.Equals(b.InputAmount().Currency) ||
		!a.OutputAmount().Currency.Equals(b.OutputAmount().Currency) {
		// TODO: better error handling
		panic(ErrInvalidCurrency)
	}

	if a.OutputAmount().EqualTo(b.OutputAmount().Fraction) {
		if a.InputAmount().EqualTo(b.InputAmount().Fraction) {
			return 0
		}
		// trade A requires less input than trade B, so A should come first
		if a.InputAmount().LessThan(b.InputAmount().Fraction) {
			return -1
		}
		return 1
	}

	// tradeA has less output than trade B, so should come second
	if a.OutputAmount().LessThan(b.OutputAmount().Fraction) {
		return 1
	}
	return -1
}

// extension of the input output comparator that also considers other dimensions of the trade in ranking them
func TradeComparator(a, b *Trade) int {
	ioComp := InputOutputComparator(a, b)
	if ioComp != 0 {
		return ioComp
	}

	// consider lowest slippage next, since these are less likely to fail
	if a.priceImpact.LessThan(b.priceImpact.Fraction) {
		return -1
	}
	if a.priceImpact.GreaterThan(b.priceImpact.Fraction) {
		return 1
	}
	// finally consider the number of hops since each hop costs gas
	return len(a.route.Path) - len(b.route.Path)
}

// given an array of items sorted by `comparator`, insert an item into its sort index and constrain the size to
// `maxSize` by removing the last item
func SortedInsert(items []*Trade, add *Trade, maxSize int, comparator func(a, b *Trade) int) (sortedItems []*Trade, pop *Trade, err error) {
	if maxSize <= 0 {
		panic("MAX_SIZE_ZERO")
	}
	itemsLen := len(items)
	// this is an invariant because the interface cannot return multiple removed items if items.length exceeds maxSize
	if itemsLen > maxSize {
		panic("ITEMS_SIZE")
	}

	// short circuit first item add
	if itemsLen == 0 {
		items = append(items, add)
		return items, nil, nil
	}

	isFull := (itemsLen == maxSize)
	// short circuit if full and the additional item does not come before the last item
	if isFull && comparator(items[itemsLen-1], add) <= 0 {
		return items, add, nil
	}

	lo, hi := 0, itemsLen
	for lo < hi {
		mid := (hi-lo)/2 + lo
		if comparator(items[mid], add) <= 0 {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	items = append(items[:lo], append([]*Trade{add}, items[lo:]...)...)
	if isFull {
		pop = items[itemsLen]
		items = items[:itemsLen]
	}
	return items, pop, nil
}

/**
 * Given a list of pairs, and a fixed amount in, returns the top `maxNumResults` trades that go from an input token
 * amount to an output token, making at most `maxHops` hops.
 * Note this does not consider aggregation, as routes are linear. It's possible a better route exists by splitting
 * the amount in among multiple routes.
 * @param pairs the pairs to consider in finding the best trade
 * @param currencyAmountIn exact amount of input currency to spend
 * @param currencyOut the desired currency out
 * @param maxNumResults maximum number of results to return
 * @param maxHops maximum number of hops a returned trade can make, e.g. 1 hop goes through a single pair
 * @param currentPairs used in recursion; the current list of pairs
 * @param originalAmountIn used in recursion; the original value of the currencyAmountIn parameter
 * @param bestTrades used in recursion; the current list of best trades
 */
func BestTradeExactIn(
	pairs []*Pair,
	currencyAmountIn *TokenAmount,
	currencyOut *Token,
	options *BestTradeOptions,
	// used in recursion.
	currentPairs []*Pair,
	originalAmountIn *TokenAmount,
	bestTrades []*Trade,
) (sortedItems []*Trade, pop *Trade, err error) {
	if len(pairs) == 0 {
		panic("PAIRS")
	}
	if options == nil || options.MaxHops <= 0 {
		panic("MAX_HOPS")
	}
	if !(originalAmountIn == currencyAmountIn || len(currentPairs) > 0) {
		panic("INVALID_RECURSION")
	}

	amountIn, tokenOut := currencyAmountIn, currencyOut
	for i := 0; i < len(pairs); i++ {
		pair := pairs[i]
		// pair irrelevant
		if !pair.Token0().Equals(amountIn.Token) && !pair.Token1().Equals(amountIn.Token) {
			continue
		}
		if pair.Reserve0().EqualTo(ZeroFraction) || pair.Reserve1().EqualTo(ZeroFraction) {
			continue
		}

		amountOut, _, err := pair.GetOutputAmount(amountIn)
		if err != nil {
			// input too low
			if err == ErrInsufficientInputAmount {
				continue
			}
			return nil, nil, err
		}

		// we have arrived at the output token, so this is the final trade of one of the paths
		if amountOut.Token.Equals(tokenOut) {
			var route *Route
			route, err = NewRoute(append(currentPairs, pair), originalAmountIn.Token, currencyOut)
			if err != nil {
				return nil, nil, err
			}
			var trade *Trade
			trade, err = NewTrade(route, originalAmountIn, constants.ExactInput)
			if err != nil {
				return nil, nil, err
			}
			bestTrades, pop, err = SortedInsert(bestTrades, trade, options.MaxNumResults, TradeComparator)
			if err != nil {
				return nil, nil, err
			}
			continue
		}

		if options.MaxHops > 1 && len(pairs) > 1 {
			pairsExcludingThisPair := append(pairs[:i], pairs[i+1:]...)
			options.ReduceHops()
			// otherwise, consider all the other paths that lead from this token as long as we have not exceeded maxHops
			bestTrades, pop, err = BestTradeExactIn(
				pairsExcludingThisPair,
				amountOut,
				currencyOut,
				options,
				append(currentPairs, pair),
				originalAmountIn,
				bestTrades,
			)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	return bestTrades, pop, nil
}
