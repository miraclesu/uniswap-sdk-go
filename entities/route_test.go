package entities

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/miraclesu/uniswap-sdk-go/constants"
)

// nolint funlen
func TestRoute(t *testing.T) {
	token0, err := NewToken(constants.Mainnet, common.HexToAddress("0x0000000000000000000000000000000000000001"), 18, "t0", "t0")
	if err != nil {
		t.Fatal(err)
	}
	token1, err := NewToken(constants.Mainnet, common.HexToAddress("0x0000000000000000000000000000000000000002"), 18, "t1", "t1")
	if err != nil {
		t.Fatal(err)
	}
	weth := WETH[constants.Mainnet]

	tokenAmount00, err := NewTokenAmount(token0, constants.B100)
	if err != nil {
		t.Fatal(err)
	}
	tokenAmount01, err := NewTokenAmount(token1, big.NewInt(200))
	if err != nil {
		t.Fatal(err)
	}
	tokenAmount0Weth0, err := NewTokenAmount(token0, constants.B100)
	if err != nil {
		t.Fatal(err)
	}
	tokenAmount0Weth1, err := NewTokenAmount(weth, constants.B100)
	if err != nil {
		t.Fatal(err)
	}
	tokenAmount1Weth0, err := NewTokenAmount(token1, big.NewInt(175))
	if err != nil {
		t.Fatal(err)
	}
	tokenAmount1Weth1, err := NewTokenAmount(weth, constants.B100)
	if err != nil {
		t.Fatal(err)
	}

	pair01, err := NewPair(tokenAmount00, tokenAmount01)
	if err != nil {
		t.Fatal(err)
	}
	pair0Weth, err := NewPair(tokenAmount0Weth0, tokenAmount0Weth1)
	if err != nil {
		t.Fatal(err)
	}
	pair1Weth, err := NewPair(tokenAmount1Weth0, tokenAmount1Weth1)
	if err != nil {
		t.Fatal(err)
	}

	// constructs a path from the tokens
	{
		route, err := NewRoute([]*Pair{pair01}, token0, nil)
		if err != nil {
			t.Fatal(err)
		}
		if len(route.Pairs) != 1 || route.Pairs[0] != pair01 {
			t.Error("wrong pairs for route")
		}
		if len(route.Path) != 2 || route.Path[0] != token0 || route.Path[1] != token1 {
			t.Error("wrong path for route")
		}
		if route.Input != token0 {
			t.Error("wrong input for route")
		}
		if route.Output != token1 {
			t.Error("wrong output for route")
		}
		if route.ChainID() != constants.Mainnet {
			t.Error("wrong chain id for route")
		}
	}

	// can have a token as both input and output
	{
		pairs := []*Pair{pair0Weth, pair01, pair1Weth}
		route, err := NewRoute(pairs, weth, nil)
		if err != nil {
			t.Fatal(err)
		}
		if len(route.Pairs) != len(pairs) {
			t.Fatal("wrong pairs for route")
		}
		for i, pair := range route.Pairs {
			if pair != pairs[i] {
				t.Error("wrong pairs for route")
			}
		}
		if route.Input != weth {
			t.Error("wrong input for route")
		}
		if route.Output != weth {
			t.Error("wrong output for route")
		}
	}

	{
		// supports ether output
		pairs := []*Pair{pair0Weth}
		route, err := NewRoute(pairs, token0, weth)
		if err != nil {
			t.Fatal(err)
		}
		if len(route.Pairs) != len(pairs) {
			t.Fatal("wrong pairs for route")
		}
		for i, pair := range route.Pairs {
			if pair != pairs[i] {
				t.Error("wrong pairs for route")
			}
		}
		if route.Input != token0 {
			t.Error("wrong input for route")
		}
		if route.Output != weth {
			t.Error("wrong output for route")
		}
	}
}
