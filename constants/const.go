package constants

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

type TradeType int

const (
	ExactInput TradeType = iota
	ExactOutput
)

type Rounding int

const (
	RoundDown Rounding = iota
	RoundHalfUp
	RoundUp
)

// Valid check this rounding mode is valid
func (r Rounding) Valid() bool {
	return r == RoundDown ||
		r == RoundHalfUp ||
		r == RoundUp
}

const (
	Decimals18  = 18
	Univ2Symbol = "UNI-V2"
	Univ2Name   = "Uniswap V2"
)

var (
	MinimumLiquidity = big.NewInt(1000)

	Zero  = big.NewInt(0)
	One   = big.NewInt(1)
	Two   = big.NewInt(2)
	Three = big.NewInt(3)
	Five  = big.NewInt(5)
	Ten   = big.NewInt(10)

	B100  = big.NewInt(100)
	B997  = big.NewInt(997)
	B1000 = big.NewInt(1000)
)

type SolidityType string

const (
	Uint8   SolidityType = "uint8"
	Uint256 SolidityType = "uint256"
)

var (
	SolidityTypeMaxima = map[SolidityType]*big.Int{
		Uint8:   big.NewInt(0xff),
		Uint256: math.MaxBig256,
	}
)

var (
	FactoryAddress = common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	InitCodeHash   = common.FromHex("0x96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f")
)
