package utils

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/miraclesu/uniswap-sdk-go/constants"
)

func ValidateSolidityTypeInstance(value *big.Int, t constants.SolidityType) error {
	if constants.Zero.Cmp(value) < 0 || value.Cmp(constants.SolidityTypeMaxima[t]) > 0 {
		return fmt.Errorf(`%v is not a %s`, value, t)
	}
	return nil
}

// warns if addresses are not checksummed
func ValidateAndParseAddress(address string) common.Address {
	// TODO print warns if adderss is not checksummed
	return common.HexToAddress(address)
}
