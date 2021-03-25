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

	_ = pair_0_1
	_ = pair_0_2
	_ = pair_0_3
	_ = pair_1_2
	_ = pair_1_3

	// use WETH as ETHR
	tokenETHR := WETH[constants.Mainnet]
	tokenAmount_0_weth, _ := NewTokenAmount(tokenETHR, big.NewInt(1000))
	pair_weth_0, _ := NewPair(tokenAmount_0_weth, tokenAmount_0_1000)

	tokenAmount_0_0, _ := NewTokenAmount(token0, big.NewInt(0))
	tokenAmount_1_0, _ := NewTokenAmount(token1, big.NewInt(0))
	empty_pair_0_1, _ := NewPair(tokenAmount_0_0, tokenAmount_1_0)
	_ = empty_pair_0_1

	{
		route, _ := NewRoute([]*Pair{pair_weth_0}, tokenETHR, nil)
		tokenAmountETHER, _ := NewTokenAmount(tokenETHR, big.NewInt(100))
		trade, _ := NewTrade(route, tokenAmountETHER, constants.ExactInput)

		// can be constructed with ETHER as input
		{
			expect := tokenETHR.Currency
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
		route, _ = NewRoute([]*Pair{pair_weth_0}, tokenETHR, token0)
		tokenAmount_0_100, _ := NewTokenAmount(token0, big.NewInt(100))
		trade, _ = NewTrade(route, tokenAmount_0_100, constants.ExactOutput)
		{
			expect := tokenETHR.Currency
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

		route, _ = NewRoute([]*Pair{pair_weth_0}, token0, tokenETHR)
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
			expect := tokenETHR.Currency
			output := trade.outputAmount.Currency
			if !expect.Equals(output) {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}
	}
}
