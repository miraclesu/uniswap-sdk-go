package entities

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/miraclesu/uniswap-sdk-go/constants"
	"github.com/miraclesu/uniswap-sdk-go/utils"
)

func TestGetAddress(t *testing.T) {
	var tests = []struct {
		Input  [2]string
		Output string
	}{
		{
			// CRO,USDC
			[2]string{"0xa0b73e1ff0b80914ab6fe0444e65848c4c34450b", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"},
			"0x35e8777ec43cb99e935588dcd0305268f06c1274",
		},
		{
			// CRO,USDC
			// cover cache
			[2]string{"0xa0b73e1ff0b80914ab6fe0444e65848c4c34450b", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"},
			"0x35e8777ec43cb99e935588dcd0305268f06c1274",
		},
		{
			// WBTC,DAI
			[2]string{"0x2260fac5e5542a773aa44fbcfedf7c193bc2c599", "0x6b175474e89094c44da98b954eedeac495271d0f"},
			"0x7a53d28d5855a0addcd0b6cf4129dcef2d9c28e3",
		},
		{
			// WBTC,AAVE
			// cover cache
			[2]string{"0x2260fac5e5542a773aa44fbcfedf7c193bc2c599", "0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9"},
			"0xa3603484ebfa1675778ce2de02c0ce96678a4f34",
		},
	}
	for i, test := range tests {
		output := _PairAddressCache.GetAddress(common.HexToAddress(test.Input[0]), common.HexToAddress(test.Input[1]))
		if output.String() != utils.ValidateAndParseAddress(test.Output).String() {
			t.Errorf("test #%d: failed to match when it should (%s != %s)", i, output, test.Output)
		}
	}
}

// nolint funlen
func TestPair(t *testing.T) {
	USDC, _ := NewToken(constants.Mainnet, common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"), 18, "USDC", "USD Coin")
	DAI, _ := NewToken(constants.Mainnet, common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"), 18, "DAI", "DAI Stablecoin")
	tokenAmountUSDC, _ := NewTokenAmount(USDC, constants.B100)
	tokenAmountDAI, _ := NewTokenAmount(DAI, constants.B100)
	tokenAmountUSDC101, _ := NewTokenAmount(USDC, big.NewInt(101))
	tokenAmountDAI101, _ := NewTokenAmount(DAI, big.NewInt(101))

	// cannot be used for tokens on different chains
	{
		tokenAmountB, _ := NewTokenAmount(WETH[constants.Rinkeby], constants.B100)
		_, output := NewPair(tokenAmountUSDC, tokenAmountB)
		expect := ErrDiffChainID
		if expect != output {
			t.Errorf("expect[%+v], but got[%+v]", expect, output)
		}
	}

	// returns the correct address
	{
		output := _PairAddressCache.GetAddress(DAI.Address, USDC.Address)
		expect := "0xB5214eDee5741324a13539bcc207Bc549e2491FF"
		if output.String() != expect {
			t.Errorf("expect[%+v], but got[%+v]", expect, output)
		}
	}

	{
		pairA, _ := NewPair(tokenAmountUSDC, tokenAmountDAI)
		pairB, _ := NewPair(tokenAmountDAI, tokenAmountUSDC)
		expect := DAI
		// always is the token that sorts before
		output := pairA.Token0()
		if !expect.Equals(output) {
			t.Errorf("expect[%+v], but got[%+v]", expect, output)
		}
		output = pairB.Token0()
		if !expect.Equals(output) {
			t.Errorf("expect[%+v], but got[%+v]", expect, output)
		}

		expect = USDC
		// always is the token that sorts after
		output = pairA.Token1()
		if !expect.Equals(output) {
			t.Errorf("expect[%+v], but got[%+v]", expect, output)
		}
		output = pairB.Token1()
		if !expect.Equals(output) {
			t.Errorf("expect[%+v], but got[%+v]", expect, output)
		}
	}

	{
		pairA, _ := NewPair(tokenAmountUSDC, tokenAmountDAI101)
		pairB, _ := NewPair(tokenAmountDAI101, tokenAmountUSDC)
		expect := tokenAmountDAI101
		// always comes from the token that sorts before
		output := pairA.Reserve0()
		if !expect.Fraction.EqualTo(output.Fraction) {
			t.Errorf("expect[%+v], but got[%+v]", expect.Fraction, output.Fraction)
		}
		output = pairB.Reserve0()
		if !expect.Fraction.EqualTo(output.Fraction) {
			t.Errorf("expect[%+v], but got[%+v]", expect.Fraction, output.Fraction)
		}

		expect = tokenAmountUSDC
		// always comes from the token that sorts after
		output = pairA.Reserve1()
		if !expect.Fraction.EqualTo(output.Fraction) {
			t.Errorf("expect[%+v], but got[%+v]", expect.Fraction, output.Fraction)
		}
		output = pairB.Reserve1()
		if !expect.Fraction.EqualTo(output.Fraction) {
			t.Errorf("expect[%+v], but got[%+v]", expect.Fraction, output.Fraction)
		}
	}

	{
		pairA, _ := NewPair(tokenAmountUSDC101, tokenAmountDAI)
		pairB, _ := NewPair(tokenAmountDAI, tokenAmountUSDC101)
		expect := NewPrice(DAI.Currency, USDC.Currency, constants.B100, big.NewInt(101))
		// returns price of token0 in terms of token1
		output := pairA.Token0Price()
		if !expect.Fraction.EqualTo(output.Fraction) {
			t.Errorf("expect[%+v], but got[%+v]", expect.Fraction, output.Fraction)
		}
		output = pairB.Token0Price()
		if !expect.Fraction.EqualTo(output.Fraction) {
			t.Errorf("expect[%+v], but got[%+v]", expect.Fraction, output.Fraction)
		}

		expect = NewPrice(USDC.Currency, DAI.Currency, big.NewInt(101), constants.B100)
		// returns price of token1 in terms of token0
		output = pairA.Token1Price()
		if !expect.Fraction.EqualTo(output.Fraction) {
			t.Errorf("expect[%+v], but got[%+v]", expect.Fraction, output.Fraction)
		}
		output = pairB.Token1Price()
		if !expect.Fraction.EqualTo(output.Fraction) {
			t.Errorf("expect[%+v], but got[%+v]", expect.Fraction, output.Fraction)
		}
	}
}
