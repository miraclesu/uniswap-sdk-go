package entities

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/miraclesu/uniswap-sdk-go/constants"
)

// nolint funlen
func TestToken(t *testing.T) {
	addressOne := common.HexToAddress("0x0000000000000000000000000000000000000001")
	addressTwo := common.HexToAddress("0x0000000000000000000000000000000000000002")

	// fails if address differs
	{
		token1, err := NewToken(constants.Mainnet, addressOne, 18, "", "")
		if err != nil {
			t.Fatal(err)
		}
		token2, err := NewToken(constants.Mainnet, addressTwo, 18, "", "")
		if err != nil {
			t.Fatal(err)
		}
		if token1.Equals(token2) {
			t.Error("should false if address differs")
		}
	}

	// false if chain id differs
	{
		token1, err := NewToken(constants.Ropsten, addressOne, 18, "", "")
		if err != nil {
			t.Fatal(err)
		}
		token2, err := NewToken(constants.Mainnet, addressOne, 18, "", "")
		if err != nil {
			t.Fatal(err)
		}
		if token1.Equals(token2) {
			t.Error("should false if chain id differs")
		}
	}

	// true if only decimals differs
	{
		token1, err := NewToken(constants.Mainnet, addressOne, 9, "", "")
		if err != nil {
			t.Fatal(err)
		}
		token2, err := NewToken(constants.Mainnet, addressOne, 18, "", "")
		if err != nil {
			t.Fatal(err)
		}
		if !token1.Equals(token2) {
			t.Error("should true if only decimals differs")
		}
	}

	// true if address is the same
	{
		token1, err := NewToken(constants.Mainnet, addressOne, 18, "", "")
		if err != nil {
			t.Fatal(err)
		}
		token2, err := NewToken(constants.Mainnet, addressOne, 18, "", "")
		if err != nil {
			t.Fatal(err)
		}
		if !token1.Equals(token2) {
			t.Error("should true if address is the same")
		}
	}

	// true on reference equality
	{
		token, err := NewToken(constants.Mainnet, addressOne, 18, "", "")
		if err != nil {
			t.Fatal(err)
		}
		if !token.Equals(token) {
			t.Error("should true on reference equality")
		}
	}

	// true even if name/symbol/decimals differ
	{
		token1, err := NewToken(constants.Mainnet, addressOne, 9, "abc", "def")
		if err != nil {
			t.Fatal(err)
		}
		token2, err := NewToken(constants.Mainnet, addressOne, 18, "ghi", "jkl")
		if err != nil {
			t.Fatal(err)
		}
		if !token1.Equals(token2) {
			t.Error("true even if name/symbol/decimals differ")
		}
	}
}
