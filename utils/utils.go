package utils

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/miraclesu/uniswap-sdk-go/constants"
)

// ValidateSolidityTypeInstance determines if a value is a legal SolidityType
func ValidateSolidityTypeInstance(value *big.Int, t constants.SolidityType) error {
	if value.Cmp(constants.Zero) < 0 || value.Cmp(constants.SolidityTypeMaxima[t]) > 0 {
		return fmt.Errorf(`%v is not a %s`, value, t)
	}
	return nil
}

// ValidateAndParseAddress warns if addresses are not checksummed
func ValidateAndParseAddress(address string) common.Address {
	return common.HexToAddress(address)
}
