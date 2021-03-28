package constants

import (
	"testing"
)

func TestChainIDString(t *testing.T) {
	var tests = []struct {
		expect string
		output ChainID
	}{
		{"Mainnet", Mainnet},
		{"Ropsten", Ropsten},
		{"Rinkeby", Rinkeby},
		{"Goerli", Goerli},
		{"Kovan", Kovan},
		{"ChainID(2)", ChainID(2)},
	}
	for i, test := range tests {
		output := test.output.String()
		if test.expect != output {
			t.Fatalf("test #%d: expect[%+v], but got[%+v]", i, test.expect, output)
		}
	}
}
