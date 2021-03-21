package entities

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"

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
