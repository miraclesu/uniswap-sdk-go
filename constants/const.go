package constants

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
)

// TradeType
const (
	ExactInput = iota + 1
	ExactOutput
)

// Rounding
const (
	RoundDown = iota + 1
	RoundHalfUp
	RoundUp
)

const (
	Decimals18 = 18
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
