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
			"0xeafac2E662Ec23860836Da89d3711BFF0260CD8D",
		},
		{
			// CRO,USDC
			// cover cache
			[2]string{"0xa0b73e1ff0b80914ab6fe0444e65848c4c34450b", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"},
			"0xeafac2E662Ec23860836Da89d3711BFF0260CD8D",
		},
		{
			// WBTC,DAI
			[2]string{"0x2260fac5e5542a773aa44fbcfedf7c193bc2c599", "0x6b175474e89094c44da98b954eedeac495271d0f"},
			"0x231B7589426Ffe1b75405526fC32aC09D44364c4",
		},
		{
			// WBTC,AAVE
			// cover cache
			[2]string{"0x2260fac5e5542a773aa44fbcfedf7c193bc2c599", "0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9"},
			"0x48978eF5BeB2d69e27DeF9C046cEbE18Ab5708Ad",
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
		expect := "0xAE461cA67B15dc8dc81CE7615e0320dA1A9aB8D5"
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

	{
		pair, _ := NewPair(tokenAmountUSDC101, tokenAmountDAI)
		// returns price of token in terms of other token
		expect := pair.Token0Price()
		output, _ := pair.PriceOf(tokenAmountDAI.Token)
		if !expect.Fraction.EqualTo(output.Fraction) {
			t.Errorf("expect[%+v], but got[%+v]", expect.Fraction, output.Fraction)
		}

		expect = pair.Token1Price()
		output, _ = pair.PriceOf(tokenAmountUSDC101.Token)
		if !expect.Fraction.EqualTo(output.Fraction) {
			t.Errorf("expect[%+v], but got[%+v]", expect.Fraction, output.Fraction)
		}

		{
			// throws if invalid token
			expect := ErrDiffToken
			_, output := pair.PriceOf(WETH[constants.Mainnet])
			if expect != output {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}
	}

	{
		pairA, _ := NewPair(tokenAmountUSDC, tokenAmountDAI101)
		pairB, _ := NewPair(tokenAmountDAI101, tokenAmountUSDC)
		expect := tokenAmountUSDC
		// returns reserves of the given token
		output, _ := pairA.ReserveOf(USDC)
		if !expect.Fraction.EqualTo(output.Fraction) {
			t.Errorf("expect[%+v], but got[%+v]", expect.Fraction, output.Fraction)
		}
		output, _ = pairB.ReserveOf(USDC)
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

		{
			// throws if not in the pair
			expect := ErrDiffToken
			_, output := pairB.ReserveOf(WETH[constants.Mainnet])
			if expect != output {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}
	}

	{
		pairA, _ := NewPair(tokenAmountUSDC, tokenAmountDAI)
		pairB, _ := NewPair(tokenAmountDAI, tokenAmountUSDC)
		expect := constants.Mainnet
		// returns the token0 chainId
		output := pairA.ChainID()
		if expect != output {
			t.Errorf("expect[%+v], but got[%+v]", expect, output)
		}
		output = pairB.ChainID()
		if expect != output {
			t.Errorf("expect[%+v], but got[%+v]", expect, output)
		}

		{
			expect := true
			// involvesToken
			output := pairA.InvolvesToken(USDC)
			if expect != output {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
			output = pairA.InvolvesToken(DAI)
			if expect != output {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
			expect = false
			output = pairA.InvolvesToken(WETH[constants.Mainnet])
			if expect != output {
				t.Errorf("expect[%+v], but got[%+v]", expect, output)
			}
		}

		{
			tokenA, _ := NewToken(constants.Rinkeby, common.HexToAddress("0x0000000000000000000000000000000000000001"), 18, "", "")
			tokenB, _ := NewToken(constants.Rinkeby, common.HexToAddress("0x0000000000000000000000000000000000000002"), 18, "", "")
			tokenAmountA, _ := NewTokenAmount(tokenA, big.NewInt(0))
			tokenAmountB, _ := NewTokenAmount(tokenB, big.NewInt(0))
			pair, _ := NewPair(tokenAmountA, tokenAmountB)
			{
				tokenAmount, _ := NewTokenAmount(pair.LiquidityToken, big.NewInt(0))
				tokenAmountA, _ := NewTokenAmount(tokenA, big.NewInt(1000))
				tokenAmountB, _ := NewTokenAmount(tokenB, big.NewInt(1000))
				// getLiquidityMinted:0
				expect := ErrInsufficientInputAmount
				_, output := pair.GetLiquidityMinted(tokenAmount, tokenAmountA, tokenAmountB)
				if expect != output {
					t.Errorf("expect[%+v], but got[%+v]", expect, output)
				}

				tokenAmountA, _ = NewTokenAmount(tokenA, big.NewInt(1000000))
				tokenAmountB, _ = NewTokenAmount(tokenB, big.NewInt(1))
				_, output = pair.GetLiquidityMinted(tokenAmount, tokenAmountA, tokenAmountB)
				if expect != output {
					t.Errorf("expect[%+v], but got[%+v]", expect, output)
				}

				tokenAmountA, _ = NewTokenAmount(tokenA, big.NewInt(1001))
				tokenAmountB, _ = NewTokenAmount(tokenB, big.NewInt(1001))
				{
					expect := "1"
					liquidity, _ := pair.GetLiquidityMinted(tokenAmount, tokenAmountA, tokenAmountB)
					output := liquidity.Raw().String()
					if expect != output {
						t.Errorf("expect[%+v], but got[%+v]", expect, output)
					}
				}
			}

			// getLiquidityMinted:!0
			tokenAmountA, _ = NewTokenAmount(tokenA, big.NewInt(10000))
			tokenAmountB, _ = NewTokenAmount(tokenB, big.NewInt(10000))
			pair, _ = NewPair(tokenAmountA, tokenAmountB)
			{
				tokenAmount, _ := NewTokenAmount(pair.LiquidityToken, big.NewInt(10000))
				tokenAmountA, _ = NewTokenAmount(tokenA, big.NewInt(2000))
				tokenAmountB, _ = NewTokenAmount(tokenB, big.NewInt(2000))
				expect := "2000"
				liquidity, _ := pair.GetLiquidityMinted(tokenAmount, tokenAmountA, tokenAmountB)
				output := liquidity.Raw().String()
				if expect != output {
					t.Errorf("expect[%+v], but got[%+v]", expect, output)
				}
			}

			// getLiquidityValue:!feeOn
			tokenAmountA, _ = NewTokenAmount(tokenA, big.NewInt(1000))
			tokenAmountB, _ = NewTokenAmount(tokenB, big.NewInt(1000))
			pair, _ = NewPair(tokenAmountA, tokenAmountB)
			tokenAmount, _ := NewTokenAmount(pair.LiquidityToken, big.NewInt(1000))
			tokenAmount500, _ := NewTokenAmount(pair.LiquidityToken, big.NewInt(500))
			{
				liquidityValue, _ := pair.GetLiquidityValue(tokenA, tokenAmount, tokenAmount, false, nil)
				{
					expect := true
					output := liquidityValue.Token.Equals(tokenA)
					if expect != output {
						t.Errorf("expect[%+v], but got[%+v]", expect, output)
					}
				}
				{
					expect := "1000"
					output := liquidityValue.Raw().String()
					if expect != output {
						t.Errorf("expect[%+v], but got[%+v]", expect, output)
					}
				}

				liquidityValue, _ = pair.GetLiquidityValue(tokenA, tokenAmount, tokenAmount500, false, nil)
				// 500
				{
					expect := true
					output := liquidityValue.Token.Equals(tokenA)
					if expect != output {
						t.Errorf("expect[%+v], but got[%+v]", expect, output)
					}
				}
				{
					expect := "500"
					output := liquidityValue.Raw().String()
					if expect != output {
						t.Errorf("expect[%+v], but got[%+v]", expect, output)
					}
				}

				liquidityValue, _ = pair.GetLiquidityValue(tokenB, tokenAmount, tokenAmount, false, nil)
				// tokenB
				{
					expect := true
					output := liquidityValue.Token.Equals(tokenB)
					if expect != output {
						t.Errorf("expect[%+v], but got[%+v]", expect, output)
					}
				}
				{
					expect := "1000"
					output := liquidityValue.Raw().String()
					if expect != output {
						t.Errorf("expect[%+v], but got[%+v]", expect, output)
					}
				}
			}

			// getLiquidityValue:feeOn
			{
				liquidityValue, _ := pair.GetLiquidityValue(tokenA, tokenAmount500, tokenAmount500, true, big.NewInt(500*500))
				{
					expect := true
					output := liquidityValue.Token.Equals(tokenA)
					if expect != output {
						t.Errorf("expect[%+v], but got[%+v]", expect, output)
					}
				}
				{
					expect := "917" // ceiling(1000 - (500 * (1 / 6)))
					output := liquidityValue.Raw().String()
					if expect != output {
						t.Errorf("expect[%+v], but got[%+v]", expect, output)
					}
				}
			}
		}
	}
}
