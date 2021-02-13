package entities

import (
	"fmt"

	"github.com/miraclesu/uniswap-sdk-go/constants"
)

var (
	ErrInvalidPairs         = fmt.Errorf("invalid pairs")
	ErrInvalidPairsChainIDs = fmt.Errorf("invalid pairs chainIDs")
	ErrInvalidInput         = fmt.Errorf("invalid token input")
	ErrInvalidOutput        = fmt.Errorf("invalid token output")
	ErrInvalidPath          = fmt.Errorf("invalid pairs for path")
)

type Route struct {
	Pairs    []*Pair
	Path     []*Token
	Input    *Token
	Output   *Token
	MidPrice *Price
}

func NewRoute(pairs []*Pair, input, output *Token) (*Route, error) {
	if len(pairs) == 0 {
		return nil, ErrInvalidPairs
	}

	for i := range pairs {
		if pairs[i].ChainID() != pairs[0].ChainID() {
			return nil, ErrInvalidPairsChainIDs
		}
	}

	if !pairs[0].InvolvesToken(input) {
		return nil, ErrInvalidInput
	}
	if !(output == nil || pairs[len(pairs)-1].InvolvesToken(output)) {
		return nil, ErrInvalidOutput
	}

	path := make([]*Token, len(pairs)+1)
	path[0] = input
	for i := range pairs {
		currentInput := path[i]
		if !(currentInput.Equals(pairs[i].Token0()) || currentInput.Equals(pairs[i].Token1())) {
			return nil, ErrInvalidPath
		}
		currentOutput := pairs[i].Token0()
		if currentInput.Equals(pairs[i].Token0()) {
			currentOutput = pairs[i].Token1()
		}
		path[i+1] = currentOutput
	}

	if output == nil {
		output = path[len(pairs)]
	}

	route := &Route{
		Pairs:  pairs,
		Path:   path,
		Input:  input,
		Output: output,
	}
	var err error
	route.MidPrice, err = NewPriceFromRoute(route)
	return route, err
}

func (r *Route) ChainID() constants.ChainID {
	return r.Pairs[0].ChainID()
}
