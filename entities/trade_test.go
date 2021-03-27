package entities

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/miraclesu/uniswap-sdk-go/constants"
)

// nolint funlen
func TestTrade(t *testing.T) {
	token0, _ := NewToken(constants.Mainnet, common.HexToAddress("0x0000000000000000000000000000000000000001"), 18, "t0", "")
	token1, _ := NewToken(constants.Mainnet, common.HexToAddress("0x0000000000000000000000000000000000000002"), 18, "t1", "")
	token2, _ := NewToken(constants.Mainnet, common.HexToAddress("0x0000000000000000000000000000000000000003"), 18, "t2", "")
	token3, _ := NewToken(constants.Mainnet, common.HexToAddress("0x0000000000000000000000000000000000000004"), 18, "t3", "")

	tokenAmount_0_100, _ := NewTokenAmount(token0, big.NewInt(100))
	tokenAmount_0_1000, _ := NewTokenAmount(token0, big.NewInt(1000))
	tokenAmount_1_1000, _ := NewTokenAmount(token1, big.NewInt(1000))
	tokenAmount_1_1200, _ := NewTokenAmount(token1, big.NewInt(1200))
	tokenAmount_2_1000, _ := NewTokenAmount(token2, big.NewInt(1000))
	tokenAmount_2_1100, _ := NewTokenAmount(token2, big.NewInt(1100))
	tokenAmount_3_900, _ := NewTokenAmount(token3, big.NewInt(900))
	tokenAmount_3_1300, _ := NewTokenAmount(token3, big.NewInt(1300))

	pair_0_1, _ := NewPair(tokenAmount_0_1000, tokenAmount_1_1000)
	pair_0_2, _ := NewPair(tokenAmount_0_1000, tokenAmount_2_1100)
	pair_0_3, _ := NewPair(tokenAmount_0_1000, tokenAmount_3_900)
	pair_1_2, _ := NewPair(tokenAmount_1_1200, tokenAmount_2_1000)
	pair_1_3, _ := NewPair(tokenAmount_1_1200, tokenAmount_3_1300)

	// use WETH as ETHR
	tokenETHER := WETH[constants.Mainnet]
	tokenAmountETHER, _ := NewTokenAmount(tokenETHER, big.NewInt(100))
	tokenAmount_0_weth, _ := NewTokenAmount(tokenETHER, big.NewInt(1000))
	pair_weth_0, _ := NewPair(tokenAmount_0_weth, tokenAmount_0_1000)

	tokenAmount_0_0, _ := NewTokenAmount(token0, big.NewInt(0))
	tokenAmount_1_0, _ := NewTokenAmount(token1, big.NewInt(0))
	empty_pair_0_1, _ := NewPair(tokenAmount_0_0, tokenAmount_1_0)
	_ = empty_pair_0_1

	{
		route, _ := NewRoute([]*Pair{pair_weth_0}, tokenETHER, nil)
		trade, _ := NewTrade(route, tokenAmountETHER, constants.ExactInput)

		// can be constructed with ETHER as input
		{
			expect := tokenETHER.Currency
			output := trade.inputAmount.Currency
			if !expect.Equals(output) {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}
		{
			expect := token0.Currency
			output := trade.outputAmount.Currency
			if !expect.Equals(output) {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}

		// can be constructed with ETHER as input for exact output
		route, _ = NewRoute([]*Pair{pair_weth_0}, tokenETHER, token0)
		trade, _ = NewTrade(route, tokenAmount_0_100, constants.ExactOutput)
		{
			expect := tokenETHER.Currency
			output := trade.inputAmount.Currency
			if !expect.Equals(output) {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}
		{
			expect := token0.Currency
			output := trade.outputAmount.Currency
			if !expect.Equals(output) {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}

		route, _ = NewRoute([]*Pair{pair_weth_0}, token0, tokenETHER)
		// can be constructed with ETHER as output
		trade, _ = NewTrade(route, tokenAmountETHER, constants.ExactOutput)
		{
			expect := token0.Currency
			output := trade.inputAmount.Currency
			if !expect.Equals(output) {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}
		{
			expect := WETH[constants.Mainnet].Currency
			output := trade.outputAmount.Currency
			if !expect.Equals(output) {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}

		// can be constructed with ETHER as output for exact input
		trade, _ = NewTrade(route, tokenAmount_0_100, constants.ExactInput)
		{
			expect := token0.Currency
			output := trade.inputAmount.Currency
			if !expect.Equals(output) {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}
		{
			expect := tokenETHER.Currency
			output := trade.outputAmount.Currency
			if !expect.Equals(output) {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}
	}

	// bestTradeExactIn
	{
		pairs := []*Pair{}
		_, output := BestTradeExactIn(pairs, tokenAmount_0_100, token2,
			NewDefaultBestTradeOptions(), nil, tokenAmount_0_100, nil)
		//throws with empty pairs
		{
			expect := ErrInvalidPairs
			if expect != output {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}

		pairs = []*Pair{pair_0_2}
		_, output = BestTradeExactIn(pairs, tokenAmount_0_100, token2, &BestTradeOptions{},
			nil, tokenAmount_0_100, nil)
		// throws with max hops of 0
		{
			expect := ErrInvalidOption
			if expect != output {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}

		pairs = []*Pair{pair_0_1, pair_0_2, pair_1_2}
		result, _ := BestTradeExactIn(pairs, tokenAmount_0_100, token2,
			NewDefaultBestTradeOptions(), nil, tokenAmount_0_100, nil)
		// provides best route
		{
			{
				var tests = []struct {
					expect int
					output int
				}{
					{2, len(result)},
					{1, len(result[0].route.Pairs)},
					{2, len(result[1].route.Pairs)},
				}
				for i, test := range tests {
					if test.expect != test.output {
						t.Fatalf("test #%d: expect[%+v], but got[%+v]", i, test.expect, test.output)
					}
				}
			}
			{
				var tests = []struct {
					expect []*Token
					output []*Token
				}{
					{[]*Token{token0, token2}, result[0].route.Path},
					{[]*Token{token0, token1, token2}, result[1].route.Path},
				}
				for i, test := range tests {
					if len(test.expect) != len(test.output) {
						t.Fatalf("test #%d: expect[%+v], but got[%+v]", i, len(test.expect), (test.output))
					}
					for j := range test.expect {
						if !test.expect[j].Equals(test.output[j]) {
							t.Errorf("test #%d#%d: expect[%+v], but got[%+v]", i, j, test.expect[j], test.output[j])
						}
					}
				}
			}

			{
				tokenAmount_2_99, _ := NewTokenAmount(token2, big.NewInt(99))
				tokenAmount_2_69, _ := NewTokenAmount(token2, big.NewInt(69))
				var tests = []struct {
					expect *TokenAmount
					output *TokenAmount
				}{
					{result[0].inputAmount, tokenAmount_0_100},
					{result[0].outputAmount, tokenAmount_2_99},
					{result[1].inputAmount, tokenAmount_0_100},
					{result[1].outputAmount, tokenAmount_2_69},
				}
				for i, test := range tests {
					if !test.expect.Equals(test.output) {
						t.Errorf("test #%d: expect[%+v], but got[%+v]", i, test.expect, test.output)
					}
				}
			}
		}

		// doesnt throw for zero liquidity pairs
		// throws with max hops of 0
		{
			pairs := []*Pair{empty_pair_0_1}
			results, err := BestTradeExactIn(pairs, tokenAmount_0_100, token1,
				NewDefaultBestTradeOptions(), nil, tokenAmount_0_100, nil)
			if err != nil {
				t.Fatalf("err should be nil, got[%+v]", err)
			}
			expect := 0
			output := len(results)
			if expect != output {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}

		tokenAmount, _ := NewTokenAmount(token0, big.NewInt(10))
		result, _ = BestTradeExactIn(pairs, tokenAmount, token2,
			&BestTradeOptions{MaxNumResults: 3, MaxHops: 1}, nil, tokenAmount, nil)
		// respects maxHops
		{
			{
				var tests = []struct {
					expect int
					output int
				}{
					{1, len(result)},
					{1, len(result[0].route.Pairs)},
				}
				for i, test := range tests {
					if test.expect != test.output {
						t.Fatalf("test #%d: expect[%+v], but got[%+v]", i, test.expect, test.output)
					}
				}
			}
			{
				var tests = []struct {
					expect []*Token
					output []*Token
				}{
					{[]*Token{token0, token2}, result[0].route.Path},
				}
				for i, test := range tests {
					if len(test.expect) != len(test.output) {
						t.Fatalf("test #%d: expect[%+v], but got[%+v]", i, len(test.expect), (test.output))
					}
					for j := range test.expect {
						if !test.expect[j].Equals(test.output[j]) {
							t.Errorf("test #%d#%d: expect[%+v], but got[%+v]", i, j, test.expect[j], test.output[j])
						}
					}
				}
			}
		}

		tokenAmount, _ = NewTokenAmount(token0, big.NewInt(1))
		result, _ = BestTradeExactIn(pairs, tokenAmount, token2,
			nil, nil, nil, nil)
		// insufficient input for one pair
		{
			{
				var tests = []struct {
					expect int
					output int
				}{
					{1, len(result)},
					{1, len(result[0].route.Pairs)},
				}
				for i, test := range tests {
					if test.expect != test.output {
						t.Fatalf("test #%d: expect[%+v], but got[%+v]", i, test.expect, test.output)
					}
				}
			}
			{
				var tests = []struct {
					expect []*Token
					output []*Token
				}{
					{[]*Token{token0, token2}, result[0].route.Path},
				}
				for i, test := range tests {
					if len(test.expect) != len(test.output) {
						t.Fatalf("test #%d: expect[%+v], but got[%+v]", i, len(test.expect), (test.output))
					}
					for j := range test.expect {
						if !test.expect[j].Equals(test.output[j]) {
							t.Errorf("test #%d#%d: expect[%+v], but got[%+v]", i, j, test.expect[j], test.output[j])
						}
					}
				}
			}
			{
				expect, _ := NewTokenAmount(token2, big.NewInt(1))
				output := result[0].outputAmount
				if !expect.Equals(output) {
					t.Errorf("expect[%+v], but got[%+v]", expect, output)
				}
			}
		}

		tokenAmount, _ = NewTokenAmount(token0, big.NewInt(10))
		result, _ = BestTradeExactIn(pairs, tokenAmount, token2,
			&BestTradeOptions{MaxNumResults: 1, MaxHops: 3}, nil, nil, nil)
		// respects n
		{
			expect := 1
			output := len(result)
			if expect != output {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}

		pairs = []*Pair{pair_0_1, pair_0_3, pair_1_3}
		result, _ = BestTradeExactIn(pairs, tokenAmount, token2,
			&BestTradeOptions{MaxNumResults: 1, MaxHops: 3}, nil, nil, nil)
		// no path
		{
			expect := 0
			output := len(result)
			if expect != output {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}

		pairs = []*Pair{pair_weth_0, pair_0_1, pair_0_3, pair_1_3}
		result, _ = BestTradeExactIn(pairs, tokenAmountETHER, token3,
			nil, nil, nil, nil)
		// works for ETHER currency input
		{
			{
				expect := 2
				output := len(result)
				if expect != output {
					t.Fatalf("expect[%+v], but got[%+v]", expect, output)
				}
			}
			{
				var tests = []struct {
					expect *Currency
					output *Currency
				}{
					{_WETHCurrency, result[0].inputAmount.Currency},
					{token3.Currency, result[0].outputAmount.Currency},
					{_WETHCurrency, result[1].inputAmount.Currency},
					{token3.Currency, result[1].outputAmount.Currency},
				}
				for i, test := range tests {
					if !test.expect.Equals(test.output) {
						t.Fatalf("test #%d: expect[%+v], but got[%+v]", i, test.expect, test.output)
					}
				}
			}
			{
				var tests = []struct {
					expect []*Token
					output []*Token
				}{
					{[]*Token{tokenETHER, token0, token1, token3}, result[0].route.Path},
					{[]*Token{tokenETHER, token0, token3}, result[1].route.Path},
				}
				for i, test := range tests {
					if len(test.expect) != len(test.output) {
						t.Fatalf("test #%d: expect[%+v], but got[%+v]", i, len(test.expect), (test.output))
					}
					for j := range test.expect {
						if !test.expect[j].Equals(test.output[j]) {
							t.Errorf("test #%d#%d: expect[%+v], but got[%+v]", i, j, test.expect[j], test.output[j])
						}
					}
				}
			}
		}

		tokenAmount, _ = NewTokenAmount(token3, big.NewInt(100))
		result, _ = BestTradeExactIn(pairs, tokenAmount, tokenETHER,
			nil, nil, nil, nil)
		// works for ETHER currency output
		{
			{
				expect := 2
				output := len(result)
				if expect != output {
					t.Fatalf("expect[%+v], but got[%+v]", expect, output)
				}
			}
			{
				var tests = []struct {
					expect *Currency
					output *Currency
				}{
					{token3.Currency, result[0].inputAmount.Currency},
					{_WETHCurrency, result[0].outputAmount.Currency},
					{token3.Currency, result[1].inputAmount.Currency},
					{_WETHCurrency, result[1].outputAmount.Currency},
				}
				for i, test := range tests {
					if !test.expect.Equals(test.output) {
						t.Fatalf("test #%d: expect[%+v], but got[%+v]", i, test.expect, test.output)
					}
				}
			}
			{
				var tests = []struct {
					expect []*Token
					output []*Token
				}{
					{[]*Token{token3, token0, tokenETHER}, result[0].route.Path},
					{[]*Token{token3, token1, token0, tokenETHER}, result[1].route.Path},
				}
				for i, test := range tests {
					if len(test.expect) != len(test.output) {
						t.Fatalf("test #%d: expect[%+v], but got[%+v]", i, len(test.expect), (test.output))
					}
					for j := range test.expect {
						if !test.expect[j].Equals(test.output[j]) {
							t.Errorf("test #%d#%d: expect[%+v], but got[%+v]", i, j, test.expect[j], test.output[j])
						}
					}
				}
			}
		}

	}
}
